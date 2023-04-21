// fjflat_test.go

package main

import (
	"testing"

	"github.com/valyala/fastjson"
)

func TestFlattenJson(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test case 1",
			input: `{"key1":{"key2":true}}`,
			want:  `{"key1.key2":true}`,
		},
		{
			name:  "Test case 2",
			input: `{"key":["val1",2]}`,
			want:  `{"key.0":"val1","key.1":2}`,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j, err := fastjson.Parse(tt.input)
			if err != nil {
				t.Errorf("Error parsing JSON input: %v", err)
			}
			got := FlattenJson(j)
			if gotStr := string(got.MarshalTo(nil)); gotStr != tt.want {
				t.Errorf("FlattenJson() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

