package model_test

import (
	"fhir-toolbox/model"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestMemSizeR4(t *testing.T) {
	tests := []struct {
		name    string
		element model.Element
		want    int
	}{
		{
			name:    "empty account",
			element: r4.Account{},
			want:    int(unsafe.Sizeof(r4.Account{})),
		},
		{
			name:    "account with id",
			element: r4.Account{Id: &r4.Id{Value: utils.Ptr("1")}},
			want:    int(unsafe.Sizeof(r4.Account{})+unsafe.Sizeof(r4.Id{})+unsafe.Sizeof("")) + len("1"),
		},
		{
			name: "account with extensions",
			element: &r4.Account{
				Extension: []r4.Extension{
					{
						Url: "http://example.com",
					},
				},
			},
			want: int(unsafe.Sizeof(r4.Account{})+unsafe.Sizeof(r4.Extension{})) +
				// because Extension.url is not a pointer, the size of the string header is already included
				len("http://example.com"),
		},
		{
			name: "account with extensions sliced",
			element: &r4.Account{
				Extension: []r4.Extension{
					{
						Url: "http://example.com",
					},
					{},
				}[:1],
			},
			want: int(unsafe.Sizeof(r4.Account{})+
				// unused capacity is counted as well
				2*unsafe.Sizeof(r4.Extension{})) + len("http://example.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.element.MemSize(); got != tt.want {
				assert.Equal(t, tt.want, got, "MemSize() should return the size of the element")
			}
		})
	}
}
