/* This function takes the Url struct and separates CamelCase strings e.g. CamelCase to Camel Case.
 * Numbers are handled like lower case letters.
 */

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

    if a != "" {
      b = strings.TrimPrefix(b, a)
      responseValue += a + " "
    } else {
      responseValue += b
    }
  }

  responseValue = strings.TrimSuffix(responseValue, " ")
  responseValue = "\"" + responseValue + "\""

  return responseValue
}
