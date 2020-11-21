package main

import (
    "net/http"
    "fmt"
)

func queryOutput(response http.ResponseWriter, request *http.Request) {
    fmt.Print("request: ")
    fmt.Println(request)
    fmt.Println("response: ")
    fmt.Println(response)
}

func main() {
    http.HandleFunc("/helloworld", func(response http.ResponseWriter, request *http.Request) {
        queryOutput(response, request)
    })
    http.ListenAndServe(":8080", nil)
}





// &{GET /helloworld HTTP/1.1 1 1 map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] {} <nil> 0 [] false localhost:8080 map[] map[] <nil> map[] [::1]:35960 /helloworld <nil> <nil> <nil> 0xc0001180c0}
