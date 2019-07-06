// curl -T main.go -H "Content-Type : text/plain" http://localhost:8080

// 임의의 바디를 POST로 전송

package main

import (
  "log"
  "net/http"
  "os"
)

func main() {
  file, err := os.Open("main.go")
  if err != nil {
    panic(err)
  }
  
  resp,err := http.Post("http://localhost:8080", "text/plain", file)
  if err != nil {
    panic(err)
  }
  
  log.Println("Status", resp.Status)
}


