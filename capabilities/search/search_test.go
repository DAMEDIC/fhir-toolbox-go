package search

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestQueryString(t *testing.T) {
	tests := []struct {
		name    string
		options Options
		want    string
	}{
		{
			name: "plain params",
			options: Options{
				Params:   Params{"a": AndList{{{Value: "value"}}}},
				Includes: nil,
				Count:    0,
				Cursor:   "",
			},
			want: "a=value",
		},
		{
			name: "params with modifier",
			options: Options{
				Params:   Params{"a": AndList{{{Modifier: "modifier", Value: "value"}}}},
				Includes: nil,
				Count:    0,
				Cursor:   "",
			},
			want: "a%3Amodifier=value",
		},
		{
			name: "params with prefix",
			options: Options{
				Params:   Params{"a": AndList{{{Prefix: "prefix", Value: "value"}}}},
				Includes: nil,
				Count:    0,
				Cursor:   "",
			},
			want: "a=prefixvalue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.options.QueryString())
		})
	}
}

func TestParseOptions(t *testing.T) {
	tests := []struct {
		name         string
		capabilities Capabilities
		params       url.Values
		tz           *time.Location
		maxCount     int
		defaultCount int
		want         Options
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "simple parameter",
			capabilities: Capabilities{
				Params: map[string]ParamDesc{"param": {Type: Token}},
			},
			params:       url.Values{"param": []string{"value"}},
			maxCount:     1,
			defaultCount: 1,
			want: Options{
				Params: Params{"param": AndList{{{Value: "value"}}}},
				Count:  1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "parameter with modifier",
			capabilities: Capabilities{
				Params: map[string]ParamDesc{"param": {Type: Token, Modifiers: []Modifier{ModifierAbove}}},
			},
			params:       url.Values{"param:above": []string{"value"}},
			maxCount:     1,
			defaultCount: 1,
			want: Options{
				Params: Params{"param": AndList{{{Value: "value", Modifier: ModifierAbove}}}},
				Count:  1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "parameter with unsupported modifier",
			capabilities: Capabilities{
				Params: map[string]ParamDesc{"param": {Type: Token}},
			},
			params:       url.Values{"param:above": []string{"value"}},
			maxCount:     1,
			defaultCount: 1,
			wantErr:      assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOptions(tt.capabilities, tt.params, tt.tz, tt.maxCount, tt.defaultCount)

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
			tt.wantErr(t, err)
		})
	}
}
