package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
)

func main() {
  var url string = "http://localhost:8080/"
  fmt.Println(url)

  resp, err := http.Get(url)
  defer resp.Body.Close()

  responseData, err := ioutil.ReadAll(resp.Body)

  fmt.Println(resp)
  fmt.Println("OKKKK")
  fmt.Println(err)
  fmt.Println(string(responseData))
}

