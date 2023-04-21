// jflat.go
//
// echo '{"key1": {"key2": ["val1","val2"]}}' | ./jflat

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "encoding/json"
    "strings"
    "strconv"
    //"reflect"
)

func FlattenJSON(j map[string]interface{}) map[string]interface{} {

    out := make(map[string]interface{})

    var flatten func(x interface{}, name string)

    flatten = func(x interface{}, name string) {

        switch value := x.(type) {

        case map[string]interface{}:
            for k, v := range value {
                flatten(v, name+k+".")
            }
        case []interface{}:
            for i, v := range value {
                flatten(v, name+strconv.Itoa(i)+".")
            }
        default:
            out[strings.TrimSuffix(name, ".")] = value
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

    // validate json bool https://pkg.go.dev/encoding/json#Valid
    isJson := json.Valid([]byte(data))
    if isJson == false {
        fmt.Println("Invalid JSON!")
        os.Exit(1)
    }

    var jdata map[string]interface{}

    unmarsh := json.Unmarshal([]byte(data), &jdata)
    if unmarsh != nil {
        // print out error if not nil
        fmt.Println(unmarsh)
        os.Exit(1)
	}

    /*
    fmt.Println("TypeOf: ",reflect.TypeOf(jdata)) //map[string]interface {}
    for key, value := range jdata {
		fmt.Println(key, ":", value)
    }
    jdata["New_Key"] = "with a new string value"
    */

    flat := FlattenJSON(jdata)

    //fmt.Println(flat)

    enc := json.NewEncoder(os.Stdout)
    enc.Encode(flat)

}

