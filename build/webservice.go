package main

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "regexp"
    "strings"
    "io"
    "time"
)

func queryOutput(toEdit string, response http.ResponseWriter, request *http.Request) (string) {

//  Save a copy of this request for editing.
    requestDump, err := httputil.DumpRequest(request, true)

    if err != nil {
      fmt.Println(err)
    }
//    fmt.Println(string(requestDump))

    toEdit = string(requestDump)
//    fmt.Println(toEdit)
    return toEdit

//   fmt.Println("response: ")
//   fmt.Println(response)

}

func editRequest(toEdit string, u Url) (Url) {

//   fmt.Println("enter editRequest")
//   fmt.Println("toEdit: ")
//   fmt.Println(toEdit)

    re := regexp.MustCompile(`/.* `) // picks the path with query of GET /path?key=value HTTP/1.1
//    fmt.Printf("%q\n", re.FindString(toEdit))
    toEdit = re.FindString(toEdit)
//   fmt.Println("path with query:", toEdit)
    re = regexp.MustCompile(`/.[^\?]*`) // picks the path of /path?key=value
    u.path = re.FindString(toEdit)
//   fmt.Println("path:", u.path + "!")
    u.path = strings.Replace(u.path, " ", "", -1)
//    fmt.Println("path:", u.path + "!")
//   fmt.Println("path:", u.path + "!")

//    re = regexp.MustCompile(`\?(.*?)\=`) //picks ?key=
    re = regexp.MustCompile(`\?.*`) //picks ?key=value
    u.key = re.FindString(toEdit)
//   fmt.Println("key:", u.key) // ?name=value

    re = regexp.MustCompile(`\=.*`) //picks ?key=value
    u.value = re.FindString(u.key)
//   fmt.Println("value1:", u.value) // =value


//    var toCut string = u.path + u.key //
//   fmt.Println("toCut:", toCut) // /path?key=
//    u.value = strings.TrimPrefix(toEdit, toCut) // picks value of /path?key=value
    u.key = strings.TrimSuffix(u.key, u.value) // ?name
    u.key = strings.TrimPrefix(u.key, "?")
    u.key = strings.Replace(u.key, " ", "", -1)
    u.value = strings.TrimPrefix(u.value, "=")
    u.value = strings.Replace(u.value, " ", "", -1)
//   fmt.Println("key:", u.key)
//   fmt.Println("value:", u.value)

//   fmt.Println("exit editRequest")
		return u
}

func camelCaseToSpace(u Url) (string) {

  var a string
  var b string = u.value
  var responseValue string

  for b != "" {
    re := regexp.MustCompile(`^.[^A-Z]*`)
    a = re.FindString(b)
//    fmt.Println(a)

    if a != "" {
      b = strings.TrimPrefix(b, a)
//      fmt.Println("b: ", b)
      responseValue += a + " "
    } else {
      responseValue += b
    }
  }

  responseValue = strings.TrimSuffix(responseValue, " ")
  responseValue = "\"" + responseValue + "\""

//  fmt.Println(responseValue)
  return responseValue
}

func sendResponse(u Url, responseValue string, response http.ResponseWriter, request *http.Request) {
//  fmt.Println("enter sendResponse")
//  fmt.Println("u.path1:", u.path)
  if u.path == "/helloworld" {
//    fmt.Println("u.path2:", u.path)
    if u.key == "" {
//    fmt.Println("u.key:", u.key)
      io.WriteString(response, "\"Hello Stranger\"")
    } else if u.key == "name" {
//    fmt.Println("u.key:", u.key)
      if u.value != ""{
        io.WriteString(response, responseValue)
      } else {
        io.WriteString(response, "no value")
      }
    } else {
      io.WriteString(response, "unkown key")
    }
  } else {
		io.WriteString(response, "unkown path")
	}
//  fmt.Println("exit sendResponse")
}

type Url struct {
  path, key, value string
}

func main() {

    var toEdit string
    var responseValue string

    u := Url{"", "", ""}

    http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
        toEdit = queryOutput(toEdit, response, request)
        u = editRequest(toEdit, u)
        responseValue = camelCaseToSpace(u)
        sendResponse(u, responseValue, response, request)
    })

    s := &http.Server{
      Addr:           ":8080",
//      Handler:        myHandler,
      ReadTimeout:    30 * time.Second,
      WriteTimeout:   30 * time.Second,
      MaxHeaderBytes: 1 << 20,
    }
//    log.Fatal(s.ListenAndServe())
    s.ListenAndServe()
//    http.ListenAndServe(":8080", nil)
}
