// This function writes different response depending on its input 

package main

import (
  "net/http"
  "io"
  "encoding/json"
)

func sendResponse(u Url, p Path, g Git, responseValue string, response http.ResponseWriter) (string) {

  if u.path == "/" + p.path_1 {
    if u.key == "" {
      response.WriteHeader(200)
      io.WriteString(response, "\"Hello Stranger\"")
      return "200"
    } else if u.key == "name" {
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
    if g.Projectname == "" {
      response.WriteHeader(400)
      io.WriteString(response, "no git repository found")
      return "400"
    }
    g = getProjectname(g)
    j, _ := json.Marshal(g)
    response.WriteHeader(200)
    io.WriteString(response, string(j))
    return "200"
  } else {
    response.WriteHeader(400)
		io.WriteString(response, "unkown path")
    return "400"
	}

  response.WriteHeader(500)
  return "500"
}
