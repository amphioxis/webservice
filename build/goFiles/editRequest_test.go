package main

import (
  "testing"
)

func TestEditRequest(t *testing.T) {

  u := Url{"", "", ""}
  u0 := editRequest("GET / HTTP/1.1", u)
  v0 := Url{"/", "", ""}
  if (u0 != v0) {
    t.Errorf("editRequest failed if no path is given")
  }

  u1 := editRequest("GET /path HTTP/1.1", u)
  v1 := Url{"/path", "", ""}
  if (u1 != v1) {
    t.Errorf("editRequest failed if only path is given")
  }

  u2 := editRequest("GET /path? HTTP/1.1", u)
  v2 := Url{"/path", "", ""}
  if (u2 != v2) {
    t.Errorf("editRequest failed if only path with '?' is given")
  }

  u3 := editRequest("GET /path?key HTTP/1.1", u)
  v3 := Url{"/path", "key", ""}
  if (u3 != v3) {
    t.Errorf("editRequest failed if path with key is given")
  }

  u4 := editRequest("GET /path?key= HTTP/1.1", u)
  v4 := Url{"/path", "key", ""}
  if (u4 != v4) {
    t.Errorf("editRequest failed if path with key and '=' is given")
  }

  u5 := editRequest("GET /path?key=value HTTP/1.1", u)
  v5 := Url{"/path", "key", "value"}
  if (u5 != v5) {
    t.Errorf("editRequest failed if path with key and value is given")
  }
}
