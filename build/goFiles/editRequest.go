/*	This function takes the 'toEdit' string which contains the full http request
 *	and parses the path with query out if given and returns an URL struct holding
 * 	the path, the value and the key of the request.
 */

package main

import (
  "regexp"
  "strings"
)

func editRequest(toEdit string, u Url) (Url) {

  re := regexp.MustCompile(`/.* `) // picks the path with query '\path?key=value' of the line 'GET /path?key=value HTTP/1.1' from full request
  toEdit = re.FindString(toEdit)
  re = regexp.MustCompile(`/.[^\?]*`) // picks the path '\path' of '/path?key=value'
  u.path = re.FindString(toEdit)
  u.path = strings.Replace(u.path, " ", "", -1)

  re = regexp.MustCompile(`\?.*`) //picks the query '?key=value' of '/path?key=value'
  u.key = re.FindString(toEdit)

  re = regexp.MustCompile(`\=.*`) //picks the value '=value' of '/key=value'
  u.value = re.FindString(u.key)

  u.key = strings.TrimSuffix(u.key, u.value)
  u.key = strings.TrimPrefix(u.key, "?")
  u.key = strings.Replace(u.key, " ", "", -1)
  u.value = strings.TrimPrefix(u.value, "=")
  u.value = strings.Replace(u.value, " ", "", -1)

  return u
}
