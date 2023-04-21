// fjflat_test.go

package main

import (
	"testing"
	"github.com/valyala/fastjson"
)

func TestFlattenJson(t *testing.T) {

    input := `{"key1":{"key2":true}}`
    want  := `{"key1.key2":true}`

    j, err := fastjson.Parse(input)
    if err != nil {
        t.Errorf("Error parsing JSON input: %v", err)
    }

    got := FlattenJson(j)

    if gotStr := string(got.MarshalTo(nil)); gotStr != want {
        t.Errorf("FlattenJson() = %v, want %v", gotStr, want)
    }
}

