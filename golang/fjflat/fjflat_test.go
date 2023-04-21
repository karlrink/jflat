// fjflat_test.go

package main

import (
	"reflect"
	"testing"

	"github.com/valyala/fastjson"
)

func TestFlattenJson(t *testing.T) {
	type args struct {
		j *fastjson.Value
	}
	tests := []struct {
		name string
		args args
		want *fastjson.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenJson(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
