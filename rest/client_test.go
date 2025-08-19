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
	pages []search.Result[model.Resource]
	calls int
}

func (m *mockSearchClient) CapabilityStatement(ctx context.Context) (basic.CapabilityStatement, error) {
	return basic.CapabilityStatement{}, nil
}

func (m *mockSearchClient) Search(ctx context.Context, resourceType string, parameters search.Parameters, options search.Options) (search.Result[model.Resource], error) {
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
