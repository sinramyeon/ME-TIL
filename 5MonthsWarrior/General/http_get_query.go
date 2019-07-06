// curl -G --data-urldencode "query=hello world" http://localhost:8080
// -G == --get
// GET 메서드로 쿼리 전송

package main

import (
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
)

func main() {
  values := url.Values{
    "query" : {"hello world"},
  }
  
  resp, _ := http.Get("http://localhost:8080"+"?"+values.Encode())
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  log.Println(string(body))

}
