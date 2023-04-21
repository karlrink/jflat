// fjflat.go

package main

import (
    "fmt"
    "bufio"
    "io/ioutil"
    "os"
    "strconv"
    "github.com/valyala/fastjson"
)


func FlattenJson(j *fastjson.Value) *fastjson.Value {
    out := fastjson.MustParse(`{}`)
    var flatten func(*fastjson.Value, string)
    flatten = func(x *fastjson.Value, name string) {
        switch x.Type() {
        case fastjson.TypeObject:
            objectLen := x.GetObject().Len()
            x.GetObject().Visit(func(k []byte, v *fastjson.Value) {
                flatten(v, name+string(k)+".")
            })
            if objectLen == 0 {
                out.Set(name[:len(name)-1], x)
            }
        case fastjson.TypeArray:
            arrayLen := len(x.GetArray())
            for i := 0; i < arrayLen; i++ {
                flatten(x.GetArray()[i], name+strconv.Itoa(i)+".")
            }
            if arrayLen == 0 {
                out.Set(name[:len(name)-1], x)
            }
        default:
            out.Set(name[:len(name)-1], x)
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

    fmt.Println(flat)

}
