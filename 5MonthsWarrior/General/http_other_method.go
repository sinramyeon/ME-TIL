// Get, Head, Post 말고 타 메서드
// curl -X DELETE http://localhost:8080

package main

import (
  "log"
  "net/http"
  "net/http/httputil"
)

func main() {

  client := &http.Client{}
  request, err := http.NewRequest("DELETE", "http://localhost:8080", nil)
   if err != nil {
    panic(err)
  }
  
  resp, err := client.Do(request)
   if err != nil {
    panic(err)
  }
  
  dump, err := httputil.DumpResponse(resp, true)
   if err != nil {
    panic(err)
  }
  
  log.Println(string(dump))

}
