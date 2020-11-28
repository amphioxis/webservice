package main

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "regexp"
    "strings"
    "io"
    "time"
    "os"
    "flag"
    "encoding/json"
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

func getProjectname(g Git) (Git) {

//  fmt.Println("enter getProjectname")
//  fmt.Println(g.projectname)
  re := regexp.MustCompile(`.*/`)
  var toCut string = re.FindString(g.Projectname)
  g.Projectname = strings.TrimPrefix(g.Projectname, toCut)
//  fmt.Println(g.projectname)
  g.Projectname = strings.TrimSuffix(g.Projectname, ".git")
//  fmt.Println(g.projectname)

  return g
}

func sendResponse(u Url, p Path, g Git, responseValue string, response http.ResponseWriter, request *http.Request) (string) {

//  fmt.Println("enter sendResponse")
//  fmt.Println("u.path1:", u.path)
  if u.path == "/" + p.path_1 {
//    fmt.Println("u.path2:", u.path)
    if u.key == "" {
//    fmt.Println("u.key:", u.key)
      response.WriteHeader(200)
      io.WriteString(response, "\"Hello Stranger\"")
      return "200"
    } else if u.key == "name" {
//    fmt.Println("u.key:", u.key)
      if u.value != ""{
        response.WriteHeader(200)
        io.WriteString(response, responseValue)
        return "200"
      } else {
        response.WriteHeader(400)
        io.WriteString(response, "no value")
        return "400"
      }
    } else {
      response.WriteHeader(400)
      io.WriteString(response, "unkown key")
      return "400"
    }
  } else if u.path == "/" + p.path_2 {
    g = getProjectname(g)
    j, _ := json.Marshal(g)
//    fmt.Println(string(j))
//    response.Header().Set("Webserver", "new Content")
    response.WriteHeader(200)
    io.WriteString(response, string(j))
//    fmt.Println("projectname:", g.projectname)
//    fmt.Println("hash:", g.hash)
    return "200"
  } else {
    response.WriteHeader(400)
		io.WriteString(response, "unkown path")
    return "400"
	}

  response.WriteHeader(500)
  return "500"
//  fmt.Println("exit sendResponse")
}

func writeLog(l Log) (int) {

  fmt.Println("Request number", l.numberOfRequests, ": \n")
  fmt.Println("Time of request:", l.date)
  fmt.Println("HTTP-Status:", l.status)
  fmt.Println("Request:")
  fmt.Println(l.request)
  return l.numberOfRequests
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
        responseValue = camelCaseToSpace(u)
        l.status = sendResponse(u, p, g, responseValue, response, request)
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
