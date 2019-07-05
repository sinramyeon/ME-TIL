http
---

> HTTP : 웹 브라우저와 웹 서버가 통신하는 절차와 형식을 규정

RFC(IETF가 만든 규약 문서) 에 따라 다양한 프로토콜이 정해졌고, 브라우저와 서버 간 기능이 여러 소프트웨어로 구현되고, 그 후 RFC 규칙에 따라졌다. 

테스트 서버
```
package main

import (
 "fmt" "log"
 "net/http"
 "net/http/httputil"
)


func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
      http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
      return
    }
    fmt.Println(string(dump))
    fmt.Fprintf(w, "<html><body>hello</body></html>"
}

func main() {

  var httpServer http.Server
  http.HandleFunc("/", handler)
  httpServer.Addr = ":8080"
  log.Println(httpServer.ListenAndServe())

}
```

위 코드를 보면 메인함수가 실행되면, "/" 경로로 접속 시, handler 함수가 실행되고, 포트 8080으로 접속되도록 설정되어 있다.

curl(다양한 통신 프로토콜을 이용하여 데이터를 전송하기 위한 라이브러리와 명령 줄 도구)을 통해 http:1.0에서 무슨 일을 할 수 있는지 알아본다.

> curl --http1.0 http://localhost:8080/greeting

greeting이라는 페이지를 요청해서, 페이지 응답을 받아올 수 있다. 수신 후에는 서버와 연결이 끊어진다.

> curl -v --http1.0 http://localhost:8080/greeting

자세한 옵션을 표시하기 위한 -v를 넣어보면, 아래와 같은 정보를 확인할 수 있다. HTTP로 요청하고, 응답받은 내용으로 GET 메서드(요청 메서드), 헤더, 200(스테이터스 코드) 등을 받아올 수 있는 것이다.

```
Trying ::1...
...
GET /gretting HTTP/1.0 200 OK
Host : localhost:8080
User-Agent : curl/7.48.0
Accept : */*
```

스테이터스 코드는 아래와 같다.

100번대 : 처리가 계속됨
200번대 : 성공
300번대 : 리디렉트
(301 Moved Permanently, 302 Found, 303 See Other 등...)
400번대 : 클라이언트 오류
(400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found 등...)
500번대 : 서버 내부 오류
(500 Internal Sever Error, 502 Bad Gateway 등...)
