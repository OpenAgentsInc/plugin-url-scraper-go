package main

import (
    "github.com/extism/go-pdk"
    "github.com/valyala/fastjson"
)

//export say_hello
func say_hello() int32 {

    // read function argument from memory
    input := pdk.Input()

    // 1️⃣ write information to the logs
    pdk.Log(pdk.LogInfo, "👋 hello this is wasm 💜") 

    // 2️⃣ get the value associated with the `route` key 
    // in the config object
    route, _ := pdk.GetConfig("route")
    // the value of `route` is
    // https://jsonplaceholder.typicode.com/todos/1

    // 3️⃣ write information to the logs
    pdk.Log(pdk.LogInfo, "🌍 calling "+route)

    // 4️⃣ make an HTTP request
    req := pdk.NewHTTPRequest(pdk.MethodGet, route)
    res := req.Send()

    // Read the result of the request
    parser := fastjson.Parser{}
    jsonValue, _ := parser.Parse(string(res.Body()))
    title := string(jsonValue.GetStringBytes("title"))

    // Prepare the return value
    output := "param: " + string(input) + " title: " + title

    mem := pdk.AllocateString(output)
    // copy output to host memory
    pdk.OutputMemory(mem)

    return 0
}

func main() {}

