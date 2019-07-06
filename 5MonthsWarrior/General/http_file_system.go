// 로컬 파일에 액세스

package main

import (
  "log"
  "net/http"
  "net/http/httputil"
)

func main() {
  transport := &ttp.Transport{}
  transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
  client := http.Client {
    Transport : transport,
  }
  
  resp, err := client.Get("file://./main.go")
   if err != nil {
    panic(err)
  }
  
  dump, err := httputil.DumpResponse(resp, true)
   if err != nil {
    panic(err)
  }
  
  log.Println(string(dump))

}
