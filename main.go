package main

import (
    "github.com/extism/go-pdk"
)

//export fetch_url_content
func fetch_url_content() int32 {
    // Read function argument from memory, which is the URL
    input := string(pdk.Input())

    // Write information to the logs indicating the URL being fetched
    pdk.Log(pdk.LogInfo, "Fetching URL: "+input)

    // Make an HTTP GET request to the URL
    req := pdk.NewHTTPRequest(pdk.MethodGet, input)
    res := req.Send()

    // Check if the request was successful by examining if the body is not empty
    if len(res.Body()) == 0 {
        errMsg := "Failed to fetch URL"
        pdk.Log(pdk.LogError, errMsg)
        mem := pdk.AllocateString(errMsg)
        pdk.OutputMemory(mem)
        return 1 // Indicate an error occurred
    }

    // Assuming the request was successful, return the body (HTML content)
    htmlContent := string(res.Body())
    mem := pdk.AllocateString(htmlContent)
    pdk.OutputMemory(mem)

    return 0 // Success
}

func main() {}

