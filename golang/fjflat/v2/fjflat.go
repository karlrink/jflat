
package main

import (
    "fmt"
    "bufio"
    "io/ioutil"
    "os"
    "strconv"
    "github.com/valyala/fastjson"
    "encoding/json"
)


func FlattenJson(j *fastjson.Value) map[string]interface{} {

	out := make(map[string]interface{})

	var flatten func(*fastjson.Value, string)
	flatten = func(x *fastjson.Value, name string) {
		switch x.Type() {
		case fastjson.TypeObject:
			objectLen := x.GetObject().Len()
			x.GetObject().Visit(func(k []byte, v *fastjson.Value) {
				flatten(v, name+string(k)+".")
			})
			if objectLen == 0 {
				out[name[:len(name)-1]] = x.String()
			}
		case fastjson.TypeArray:
			arrayLen := len(x.GetArray())
			for i := 0; i < arrayLen; i++ {
				flatten(x.GetArray()[i], name+strconv.Itoa(i)+".")
			}
			if arrayLen == 0 {
				out[name[:len(name)-1]] = x.String()
			}
		default:
			out[name[:len(name)-1]] = string(x.GetStringBytes())
		}
	}
	flatten(j, "")
	return out
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	data, err := ioutil.ReadAll(stdin)
	if err != nil {
		fmt.Printf("Failed to read: %v \n", err)
		os.Exit(1)
	}

    var p fastjson.Parser

    docObj, err := p.ParseBytes(data)
    if err != nil {
        fmt.Printf("Failed to parse: %v \n", err)
        os.Exit(1)
    }



	flat := FlattenJson(docObj)

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.Encode(flat)
}


