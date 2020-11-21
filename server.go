package main

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "regexp"
    "strings"
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

func editRequest(toEdit string, u Url) (Url) {

    fmt.Println(toEdit)

    re := regexp.MustCompile(`/.* `)
    fmt.Printf("%q\n", re.FindString(toEdit))
    toEdit = re.FindString(toEdit)
    fmt.Println(toEdit) // /helloworld?name=AlfredENeumann

    re = regexp.MustCompile(`/.[^\?]*`)
    u.path = re.FindString(toEdit)
    fmt.Println(u.path) // /helloworld

    re = regexp.MustCompile(`\?(.*?)\=`)
    u.key = re.FindString(toEdit)
    fmt.Println(u.key) // ?name=
    var toCut string = u.path + u.key
    fmt.Println(toCut)
    u.value = strings.TrimPrefix(toEdit, toCut)
    fmt.Println(u.value) // AlfredENeumann
    u.key = strings.TrimPrefix(u.key, "?")
    u.key = strings.TrimSuffix(u.key, "=")
    fmt.Println(u.key) // name

		return u
}

func camelCaseToSpace(u Url) (string) {

  var a string
  var b string = u.value
  var responseValue string

  for b != "" {
    re := regexp.MustCompile(`^.[^A-Z]*`)
    a = re.FindString(b)
    fmt.Println(a)

    if a != "" {
      b = strings.TrimPrefix(b, a)
      fmt.Println("b: ", b)
      responseValue += a + " "
    } else {
      responseValue += b
    }
  }

  return responseValue
}

type Url struct {
  path, key, value string
}

func main() {

    var toEdit string
    var responseValue string

    u := Url{"", "", ""}

    http.HandleFunc("/helloworld", func(response http.ResponseWriter, request *http.Request) {
        toEdit = queryOutput(toEdit, response, request)
        u = editRequest(toEdit, u)
        responseValue = camelCaseToSpace(u)
    })

    http.ListenAndServe(":8080", nil)
}
