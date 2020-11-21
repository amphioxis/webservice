package main

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "regexp"
)

func queryOutput(response http.ResponseWriter, request *http.Request) {

//  Save a copy of this request for editing.
    requestDump, err := httputil.DumpRequest(request, true)

    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(string(requestDump))

    var toEdit string = string(requestDump)
    fmt.Println(toEdit)

    re := regexp.MustCompile(`/.* `)
    fmt.Printf("%q\n", re.FindString(toEdit))
    toEdit = re.FindString(toEdit)
    fmt.Print(toEdit) // /helloworld?name=AlfredENeumann 



//    fmt.Println("response: ")
//   fmt.Println(response)


}

func main() {
    http.HandleFunc("/helloworld", func(response http.ResponseWriter, request *http.Request) {
        queryOutput(response, request)
    })
    http.ListenAndServe(":8080", nil)
}
