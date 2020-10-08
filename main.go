package main

import (
    "fmt"
    "net/http"
)
//letsgo runnnnn

func hello(w http.ResponseWriter, req *http.Request) {

    // Before I messed with this it was just a "string literal" now its a variable, and that variable is also a string. Functionally, it's the same, it's just an organization trick

    // string literal style
    fmt.Fprintf(w, "hello\n")

    // variable style
    fmt.Fprintf(w, something)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}

// see right here the <!DOCTYPE html>, so the browser knows its html
const something string = `ur mom`