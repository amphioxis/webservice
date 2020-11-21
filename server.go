package main

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "regexp"
)

func queryOutput(toEdit string, response http.ResponseWriter, request *http.Request) (string) {

//  Save a copy of this request for editing.
    requestDump, err := httputil.DumpRequest(request, true)

    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(string(requestDump))

    toEdit = string(requestDump)
    fmt.Println(toEdit)
    return toEdit

//   fmt.Println("response: ")
//   fmt.Println(response)


}

func editRequest(toEdit string, u Url) (string) {

    fmt.Println(toEdit)

    re := regexp.MustCompile(`/.* `)
    fmt.Printf("%q\n", re.FindString(toEdit))
    toEdit = re.FindString(toEdit)
    fmt.Print(toEdit) // /helloworld?name=AlfredENeumann 

    re = regexp.MustCompile(`/.[^\?]*`)
    u.path = re.FindString(toEdit)
    fmt.Print(u.path)
		return toEdit
}

type Url struct {
  path, key, value string
}

func main() {

    var toEdit string

    u := Url{"", "", ""}

    http.HandleFunc("/helloworld", func(response http.ResponseWriter, request *http.Request) {
        toEdit = queryOutput(toEdit, response, request)
        editRequest(toEdit, u)
    })

    http.ListenAndServe(":8080", nil)
}
