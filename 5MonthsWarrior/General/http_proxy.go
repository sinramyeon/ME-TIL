// curl -x http://localhost:8080 http://github.com

package main

import (
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
)

func main() {
  proxyUrl, err := url.Parse("http://localhost:8080")
   if err != nil {
    panic(err)
  }
  
  client := http.Client{
    Transport : &http.Transport {
      Proxy : http.ProxyURL(proxyUrl),
    },
  }
  
  resp, err := client.Get("http://github.com")
   if err != nil {
    panic(err)
  }
  
  dump, err := http.DumpResponse(resp, true)
   if err != nil {
    panic(err)
  }
  
  log.Println(string(dump))
}
