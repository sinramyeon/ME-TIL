
package main

import (
  "log"
  "net/http"
  "strings"
)

func main() {
  reader := strings.NewReader("test")
  resp, err := http.Post("http://localhost:8080", "text/plain", reader)
  if err != nil {
    panic(err)
  }
  
  log.Println("Status", resp.Status)

}
