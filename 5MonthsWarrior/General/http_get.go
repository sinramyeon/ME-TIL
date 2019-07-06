// curl http://localhost:8080

package main

import (
 "io/ioutil"
 "log"
 "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:8080") //http.Response 안에 서버에서 오는 다양한 정보가 있다
  if err != nil {
    panic(err)
  }
  
  
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  
  log.Println(string(body))
  
  // 스테이터스 코드 가져오기
  
  log.Println("Status", resp.Status)
  log.Println("Code", resp.StatusCode)
  
  // 헤더 가져오기
  
  log.Println("Header", resp.Header)
  log.Println("Content-Length", resp.Header.Get("Content-Length"))

}
