package main

import (
  "regexp"
  "strings"
)

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

  return responseValue
}
