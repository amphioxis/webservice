package main

import (
  "net/http"
  "net/http/httputil"
  "fmt"
//  "io"
  "time"
  "os"
  "flag"
//  "encoding/json"
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

type Git struct {
  Hash, Projectname string
}

type Path struct {
  path_1, path_2 string
}

type Url struct {
  path, key, value string
}

type Log struct {
  date, status, request string
  numberOfRequests int
}

func main() {

  u := Url{"", "", ""}
  l := Log{"", "", "", 0}
  g := Git{"", ""}
  p := Path{"", ""}

  path_1 := flag.String("path_1", "helloworld", "path to take output from")
  path_2 := flag.String("path_2", "versionz", "path to take JSON file from")
  port := flag.String("port", "8080", "port on which the service is listening")
  maxReq := flag.Int("maxReq", 10, "maximum of allowed requests, if reached service shuts down")
  hash := flag.String("hash", "", "hash of the project")
  projectURL := flag.String("projectURL", "", "name of the project")
  flag.Parse()

  g.Hash = *hash
  g.Projectname = *projectURL
  p.path_1 = *path_1
  p.path_2 = *path_2

  var toEdit string
  var responseValue string
  var i int = 0


  http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

    i++
    l.numberOfRequests = i
    l.date = time.Now().String()
    toEdit = queryOutput(toEdit, response, request)
    l.request = toEdit
    u = editRequest(toEdit, u)

		if u.value != "" {
      responseValue = camelCaseToSpace(u)
		}

    l.status = sendResponse(u, p, g, responseValue, response) //, request)
    writeLog(l)

    if i == *maxReq {
      fmt.Println("Reached", i, "requests, shutdown")
      os.Exit(0)
    }
  })

  s := &http.Server{
    Addr:           ":" + *port,
    ReadTimeout:    30 * time.Second,
    WriteTimeout:   30 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
//    log.Fatal(s.ListenAndServe())
  s.ListenAndServe()
//    http.ListenAndServe(":8080", nil)
}
