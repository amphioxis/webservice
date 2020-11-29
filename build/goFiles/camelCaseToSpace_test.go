package main

import (
  "testing"
)

func TestCamelCaseToSpace(t *testing.T) {

  var responseValue string

  u := Url{"/path", "key", "AlfredENeumann"}
  responseValue = camelCaseToSpace(u)

  if responseValue != "\"Alfred E Neumann\"" {
    t.Errorf("camelToSpace failed with given value")
  }
}
