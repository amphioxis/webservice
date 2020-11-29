package main

import (
//  "fmt"
  "regexp"
  "strings"
)

func editRequest(toEdit string, u Url) (Url) {

  re := regexp.MustCompile(`/.* `) // picks the path with query of GET /path?key=value HTTP/1.1
  toEdit = re.FindString(toEdit)
  re = regexp.MustCompile(`/.[^\?]*`) // picks the path of /path?key=value
  u.path = re.FindString(toEdit)
  u.path = strings.Replace(u.path, " ", "", -1)

  re = regexp.MustCompile(`\?.*`) //picks ?key=value
  u.key = re.FindString(toEdit)

  re = regexp.MustCompile(`\=.*`) //picks ?key=value
  u.value = re.FindString(u.key)


  u.key = strings.TrimSuffix(u.key, u.value) // ?name
  u.key = strings.TrimPrefix(u.key, "?")
  u.key = strings.Replace(u.key, " ", "", -1)
  u.value = strings.TrimPrefix(u.value, "=")
  u.value = strings.Replace(u.value, " ", "", -1)

  return u
}
