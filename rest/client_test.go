package rest

import (
	"context"
	"io"
	"testing"

	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
)

// mockSearchClient implements GenericSearch for testing
type mockSearchClient struct {
	pages       []search.Result[model.Resource]
	calls       int
	lastOptions search.Options // Track the options from the last call
}

func (m *mockSearchClient) CapabilityStatement(ctx context.Context) (basic.CapabilityStatement, error) {
	return basic.CapabilityStatement{}, nil
}

func (m *mockSearchClient) Search(ctx context.Context, resourceType string, parameters search.Parameters, options search.Options) (search.Result[model.Resource], error) {
	// Store the options for verification in tests
	m.lastOptions = options

	// If cursor is provided, use it to determine which page to return
	if options.Cursor != "" {
		// Parse cursor as page number for simplicity
		pageNum := int(options.Cursor[0] - '0') // Simple cursor parsing for test
		if pageNum >= 0 && pageNum < len(m.pages) {
			m.calls++
			return m.pages[pageNum], nil
		}
	}

	// Return first page by default
	if len(m.pages) > 0 {
		m.calls++
		return m.pages[0], nil
	}

	return search.Result[model.Resource]{}, nil
}

func TestIterator(t *testing.T) {
	tests := []struct {
		name              string
		setupTest         func() (interface{}, *mockSearchClient)
		expectedCalls     []iteratorCall
		expectedMockCalls int
	}{
		{
			name: "single_page",
			setupTest: func() (interface{}, *mockSearchClient) {
				patient1 := &r4.Patient{Id: &r4.Id{Value: ptr.To("patient1")}}
				patient2 := &r4.Patient{Id: &r4.Id{Value: ptr.To("patient2")}}

				initialResult := search.Result[r4.Patient]{
					Resources: []r4.Patient{*patient1, *patient2},
					Next:      "", // No next page
				}

				mockClient := &mockSearchClient{
					pages: []search.Result[model.Resource]{},
				}

				return initialResult, mockClient
			},
			expectedCalls: []iteratorCall{
				{expectError: false, expectedResourceCount: 2},
				{expectError: true, expectedErr: io.EOF},
			},
			expectedMockCalls: 0,
		},
		{
			name: "multiple_pages",
			setupTest: func() (interface{}, *mockSearchClient) {
				initialResult := search.Result[model.Resource]{
					Resources: []model.Resource{&r4.Patient{Id: &r4.Id{Value: ptr.To("patient1")}}},
					Next:      "1", // Points to second page
				}

				page2 := search.Result[model.Resource]{
					Resources: []model.Resource{&r4.Patient{Id: &r4.Id{Value: ptr.To("patient2")}}},
					Next:      "2", // Points to third page
				}
				page3 := search.Result[model.Resource]{
					Resources: []model.Resource{&r4.Patient{Id: &r4.Id{Value: ptr.To("patient3")}}},
					Next:      "", // No more pages
				}

				mockClient := &mockSearchClient{
					pages: []search.Result[model.Resource]{
						{},    // page 0 (not used)
						page2, // page 1
						page3, // page 2
					},
				}

				return initialResult, mockClient
			},
			expectedCalls: []iteratorCall{
				{expectError: false, expectedResourceCount: 1},
				{expectError: false, expectedResourceCount: 1},
				{expectError: false, expectedResourceCount: 1},
				{expectError: true, expectedErr: io.EOF},
			},
			expectedMockCalls: 2,
		},
		{
			name: "empty_result",
			setupTest: func() (interface{}, *mockSearchClient) {
				initialResult := search.Result[r4.Patient]{
					Resources: []r4.Patient{},
					Next:      "", // No next page
				}

				mockClient := &mockSearchClient{
					pages: []search.Result[model.Resource]{},
				}

				return initialResult, mockClient
			},
			expectedCalls: []iteratorCall{
				{expectError: false, expectedResourceCount: 0},
				{expectError: true, expectedErr: io.EOF},
			},
			expectedMockCalls: 0,
		},
		{
			name: "multiple_eof_calls",
			setupTest: func() (interface{}, *mockSearchClient) {
				patient1 := &r4.Patient{Id: &r4.Id{Value: ptr.To("patient1")}}

				initialResult := search.Result[r4.Patient]{
					Resources: []r4.Patient{*patient1},
					Next:      "", // No next page
				}

				mockClient := &mockSearchClient{
					pages: []search.Result[model.Resource]{},
				}

				return initialResult, mockClient
			},
			expectedCalls: []iteratorCall{
				{expectError: false, expectedResourceCount: 1},
				{expectError: true, expectedErr: io.EOF},
				{expectError: true, expectedErr: io.EOF},
				{expectError: true, expectedErr: io.EOF},
			},
			expectedMockCalls: 0,
		},
		{
			name: "pointer_patients",
			setupTest: func() (interface{}, *mockSearchClient) {
				patient1 := &r4.Patient{Id: &r4.Id{Value: ptr.To("patient1")}}
				patient2 := &r4.Patient{Id: &r4.Id{Value: ptr.To("patient2")}}

				initialResult := search.Result[*r4.Patient]{
					Resources: []*r4.Patient{patient1},
					Next:      "1", // Points to second page
				}

				mockClient := &mockSearchClient{
					pages: []search.Result[model.Resource]{
						{}, // page 0 (not used)
						{ // page 1
							Resources: []model.Resource{patient2},
							Next:      "", // No more pages
						},
					},
				}

				return initialResult, mockClient
			},
			expectedCalls: []iteratorCall{
				{expectError: false, expectedResourceCount: 1},
				{expectError: false, expectedResourceCount: 1},
				{expectError: true, expectedErr: io.EOF},
			},
			expectedMockCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialResult, mockClient := tt.setupTest()

			var iter interface{}

			// Create iterator based on the type of initial result
			switch result := initialResult.(type) {
			case search.Result[r4.Patient]:
				iter = Iterator(context.Background(), mockClient, result)
			case search.Result[model.Resource]:
				iter = Iterator(context.Background(), mockClient, result)
			case search.Result[*r4.Patient]:
				iter = Iterator(context.Background(), mockClient, result)
			default:
				t.Fatalf("Unsupported result type: %T", initialResult)
			}

			// Execute all expected calls
			for i, expectedCall := range tt.expectedCalls {
				var err error
				var resourceCount int

				// Call Next() based on iterator type
				switch it := iter.(type) {
				case *iterator[r4.Patient]:
					page, nextErr := it.Next()
					err = nextErr
					if err == nil {
						resourceCount = len(page.Resources)
					}
				case *iterator[model.Resource]:
					page, nextErr := it.Next()
					err = nextErr
					if err == nil {
						resourceCount = len(page.Resources)
					}
				case *iterator[*r4.Patient]:
					page, nextErr := it.Next()
					err = nextErr
					if err == nil {
						resourceCount = len(page.Resources)
					}
				default:
					t.Fatalf("Unsupported iterator type: %T", iter)
				}

				// Verify results
				if expectedCall.expectError {
					if err != expectedCall.expectedErr {
						t.Errorf("Call %d: expected error %v, got %v", i+1, expectedCall.expectedErr, err)
					}
				} else {
					if err != nil {
						t.Errorf("Call %d: unexpected error: %v", i+1, err)
					}
					if resourceCount != expectedCall.expectedResourceCount {
						t.Errorf("Call %d: expected %d resources, got %d", i+1, expectedCall.expectedResourceCount, resourceCount)
					}
				}
			}

			// Verify mock calls
			if mockClient.calls != tt.expectedMockCalls {
				t.Errorf("Expected %d mock calls, got %d", tt.expectedMockCalls, mockClient.calls)
			}
		})
	}
}

type iteratorCall struct {
	expectError           bool
	expectedErr           error
	expectedResourceCount int
}

func TestIteratorCountParameter(t *testing.T) {
	// Create initial result with 3 resources
	initialResult := search.Result[r4.Patient]{
		Resources: []r4.Patient{
			{Id: &r4.Id{Value: ptr.To("patient1")}},
			{Id: &r4.Id{Value: ptr.To("patient2")}},
			{Id: &r4.Id{Value: ptr.To("patient3")}},
		},
		Next: "1", // Points to next page
	}

	// Create second page with 2 resources
	page2 := search.Result[model.Resource]{
		Resources: []model.Resource{
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient4")}},
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient5")}},
		},
		Next: "", // No more pages
	}

	mockClient := &mockSearchClient{
		pages: []search.Result[model.Resource]{
			{},    // page 0 (not used)
			page2, // page 1
		},
	}

	iter := Iterator(context.Background(), mockClient, initialResult)

	// First call should return initial result without calling Search
	result1, err := iter.Next()
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}
	if len(result1.Resources) != 3 {
		t.Errorf("Expected 3 resources in first result, got %d", len(result1.Resources))
	}
	if mockClient.calls != 0 {
		t.Errorf("Expected 0 mock calls after first Next(), got %d", mockClient.calls)
	}

	// Second call should fetch next page and pass Count=3 (length of initial result)
	result2, err := iter.Next()
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}
	if len(result2.Resources) != 2 {
		t.Errorf("Expected 2 resources in second result, got %d", len(result2.Resources))
	}
	if mockClient.calls != 1 {
		t.Errorf("Expected 1 mock call after second Next(), got %d", mockClient.calls)
	}

	// Verify that Count was set to the length of the previous page (3)
	if mockClient.lastOptions.Count != 3 {
		t.Errorf("Expected Count=3 in search options, got Count=%d", mockClient.lastOptions.Count)
	}

	// Verify cursor was passed correctly
	if mockClient.lastOptions.Cursor != "1" {
		t.Errorf("Expected Cursor='1' in search options, got Cursor='%s'", mockClient.lastOptions.Cursor)
	}

	// Third call should return EOF
	_, err = iter.Next()
	if err != io.EOF {
		t.Errorf("Expected EOF on third call, got: %v", err)
	}
}

func TestIteratorCountParameterMultiplePages(t *testing.T) {
	// Test that Count is maintained across multiple pages
	initialResult := search.Result[r4.Patient]{
		Resources: []r4.Patient{
			{Id: &r4.Id{Value: ptr.To("patient1")}},
			{Id: &r4.Id{Value: ptr.To("patient2")}},
		},
		Next: "1", // Points to page 1
	}

	page1 := search.Result[model.Resource]{
		Resources: []model.Resource{
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient3")}},
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient4")}},
		},
		Next: "2", // Points to page 2
	}

	page2 := search.Result[model.Resource]{
		Resources: []model.Resource{
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient5")}},
		},
		Next: "", // No more pages
	}

	mockClient := &mockSearchClient{
		pages: []search.Result[model.Resource]{
			{},    // page 0 (not used)
			page1, // page 1
			page2, // page 2
		},
	}

	iter := Iterator(context.Background(), mockClient, initialResult)

	// First call returns initial result (2 resources)
	result1, err := iter.Next()
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}
	if len(result1.Resources) != 2 {
		t.Errorf("Expected 2 resources in first result, got %d", len(result1.Resources))
	}

	// Second call fetches page 1, Count should be 2
	result2, err := iter.Next()
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}
	if len(result2.Resources) != 2 {
		t.Errorf("Expected 2 resources in second result, got %d", len(result2.Resources))
	}
	if mockClient.lastOptions.Count != 2 {
		t.Errorf("Expected Count=2 for page 1, got Count=%d", mockClient.lastOptions.Count)
	}
	if mockClient.lastOptions.Cursor != "1" {
		t.Errorf("Expected Cursor='1' for page 1, got Cursor='%s'", mockClient.lastOptions.Cursor)
	}

	// Third call fetches page 2, Count should still be 2 (from page 1)
	result3, err := iter.Next()
	if err != nil {
		t.Fatalf("Third call failed: %v", err)
	}
	if len(result3.Resources) != 1 {
		t.Errorf("Expected 1 resource in third result, got %d", len(result3.Resources))
	}
	if mockClient.lastOptions.Count != 2 {
		t.Errorf("Expected Count=2 for page 2, got Count=%d", mockClient.lastOptions.Count)
	}
	if mockClient.lastOptions.Cursor != "2" {
		t.Errorf("Expected Cursor='2' for page 2, got Cursor='%s'", mockClient.lastOptions.Cursor)
	}

	// Fourth call should return EOF
	_, err = iter.Next()
	if err != io.EOF {
		t.Errorf("Expected EOF on fourth call, got: %v", err)
	}
}

func TestIteratorCountParameterEmptyInitialResult(t *testing.T) {
	// Test Count parameter when initial result is empty but has next page
	initialResult := search.Result[r4.Patient]{
		Resources: []r4.Patient{}, // Empty initial result
		Next:      "1",            // But has next page
	}

	page1 := search.Result[model.Resource]{
		Resources: []model.Resource{
			&r4.Patient{Id: &r4.Id{Value: ptr.To("patient1")}},
		},
		Next: "", // No more pages
	}

	mockClient := &mockSearchClient{
		pages: []search.Result[model.Resource]{
			{},    // page 0 (not used)
			page1, // page 1
		},
	}

	iter := Iterator(context.Background(), mockClient, initialResult)

	// First call returns empty initial result
	result1, err := iter.Next()
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}
	if len(result1.Resources) != 0 {
		t.Errorf("Expected 0 resources in first result, got %d", len(result1.Resources))
	}

	// Second call fetches page 1, Count should be 0 (from empty initial result)
	result2, err := iter.Next()
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}
	if len(result2.Resources) != 1 {
		t.Errorf("Expected 1 resource in second result, got %d", len(result2.Resources))
	}
	if mockClient.lastOptions.Count != 0 {
		t.Errorf("Expected Count=0 for page 1, got Count=%d", mockClient.lastOptions.Count)
	}
	if mockClient.lastOptions.Cursor != "1" {
		t.Errorf("Expected Cursor='1' for page 1, got Cursor='%s'", mockClient.lastOptions.Cursor)
	}

	// Third call should return EOF
	_, err = iter.Next()
	if err != io.EOF {
		t.Errorf("Expected EOF on third call, got: %v", err)
	}
}
