//  This function saves a copy of this request as string for editing.
package main

import (
  "net/http"
  "net/http/httputil"
  "fmt"
)

func requestToString(toEdit string, request *http.Request) (string) {

    requestDump, err := httputil.DumpRequest(request, true)

    if err != nil {
      fmt.Println(err)
    }

    toEdit = string(requestDump)
    return toEdit
}
