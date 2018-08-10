# CORS란

Cross Origin Resource Sharing 의 약자 CORS 에 대해 알아보자.
CORS란 Cross site Http Request를 가능하게 해주는 표준 규약이다.


다른 도메인에서 가져오는 리소스가 필요할 때, 원래는 보안상의 이유로 자신과 동일한 도메인으로만 HTTp request를 보내게 제한을 해놨다. XMLHttpRequest가 크로스도메인을 요청할 수 있게 하는 것이 CORS이다.

무슨 소리냐면 타 도메인에서 태그`<>`들을 가져올 때, <script></script> 태그를 포함한 Cross site http requeste는 프로토콜, 호스트명, 포트가 같아야 가져와 진다. 그렇지만 ajax를 쓰다보면 이걸 피할수가 없고 그때 CORS요청이라는 방법을 사용한다.

CORSsms 4가지 조합으로 이루어져 있고 이 중 목적에 맞는 방식을 선택하면 된다.


## Simple Request

- GET, HEAD, POST중 하나
- POST면 Content-Type이 아래 셋 중 하나
	v application/x-www-form-urlencoded
	v multipart/form-data
	v text/plain
- 커스텀 헤더를 전송하지 않음

서버에 1번 요청하고 서버도 1번 회신후 처리 종료됨

request예제

```
GET /resources/public-data/ HTTP/1.1

Host: bar.other

User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1. 9.1b3pre) Gecko/20081130 Minefield/3.1b3pre

Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8

Accept-Language: en-us,en;q=0.5

Accept-Encoding: gzip,deflate

Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7

Connection: keep-alive

Referer: http://foo.example/examples/access-control/simpleXSInvocation.html

Origin: http://foo.example
```

response예제

```

HTTP/1.1 200 OK

Date: Mon, 01 Dec 2008 00:23:53 GMT

Server: Apache/2.0.61 

Access-Control-Allow-Origin: *

Keep-Alive: timeout=2, max=100

Connection: Keep-Alive

Transfer-Encoding: chunked

Content-Type: application/xml

```

foo.example 사이트에서 bar.other로 요청과 응답이 되고있는데 
Access-Control-Allow-Origin: * 으로 표시되어 있다. bar.other는 cross-stie방식 모든 도메인으로부터 접근이 가능하다는 뜻이다.


## Preflight Request


![](http://wiki.gurubee.net/download/attachments/28607117/cors_flow.png)

Simple Request가 아닐 때 해당. 예비 요청과 본 요청으로 나뉘는데 처음엔 예비요청을 보내서 응답을 받으면 본요청을 보낸다.

그렇지만 개발자가 따로 구분하지 않고 `Access-Control-`헤더를 정해주면 options요청으로 오는 예비 요청과 GET,POST등으로 오는 본 요청은 서버가 알아서 처리한다.

즉 예비요청, 예비응답, 본요청, 본응답으로 이뤄짐.

request 예제

```
OPTIONS /resources/post-here/ HTTP/1.1

Host: bar.other

User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre

Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8

Accept-Language: en-us,en;q=0.5

Accept-Encoding: gzip,deflate

Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7

Connection: keep-alive

Origin: http://foo.example

Access-Control-Request-Method: POST

Access-Control-Request-Headers: X-PINGOTHER
```

response 예제

```

HTTP/1.1 200 OK

Date: Mon, 01 Dec 2008 01:15:39 GMT

Server: Apache/2.0.61 (Unix)

Access-Control-Allow-Origin: http://foo.example

Access-Control-Allow-Methods: POST, GET, OPTIONS

Access-Control-Allow-Headers: X-PINGOTHER

Access-Control-Max-Age: 1728000

Vary: Accept-Encoding, Origin

Content-Encoding: gzip

Content-Length: 0

Keep-Alive: timeout=2, max=100

Connection: Keep-Alive

Content-Type: text/plain
```

예비요청을 `OPTIONS /resources/post-here/ HTTP/1.1` 로 보내놨고 `Access-Control-Request-Method: POST`를 보면 본요청에서는 POST를 썼다.

헤더를 보니 X-PINGOTHER라는 커스텀 헤더를 이런 식으로 보내고 있다. `Access-Control-Request-Headers: X-PINGOTHER`

리스폰스를 온걸 보면

```
Access-Control-Allow-Origin: http://foo.example

Access-Control-Allow-Methods: POST, GET, OPTIONS

Access-Control-Allow-Headers: X-PINGOTHER
```

셋 다 Access-Control-Allow-로 시작한다. 이후 본요청과 응답이 시행된다.


## Credential request

HTTP Cookie랑 Authentication정보를 인식할 수 있게 해주는 요청이다.
요청시 `xhr.withCredentials =true`값으로 보내면 되는데 Response Header에 꼭 `Access-Control_Allow-Credentials : true`가 포함되어야 한다.

그리고 헤더값 `Access-Control-Allow-Origin`에는 *로 아무거나 다 오케이면 안되고 도메인 이름이 명시되어야 한다.

만약 `Credentials : true`해놓고 `Access-Control-Allow-Origin : *`해놓으면 에러난다.

false로 해놓으면 요청이 거부돼서 리스폰스가 없다.

## Non Credential request

`withCredentails : false` 가 디폴트값이라 Credential요청을 쓰려면 꼭 위와 같이 명시를 해줘야 한다.


---

요약

- Ajax에는 Same Origin Policy라는 원칙이 있음
- 현재 브라우져에 보여지고 있는 HTML을 내려준 웹서버(Origin- 동일도메인,동일port,동일프로토콜)에게만 Ajax 요청을 보낼 수 있음
- CORS가 미 구현된 웹브라우저에서는 다른 도메인간 통신이 불가능
- 우회 방법으로 JSONP, IFRAME IO, CrossDomain Proxy 등이 고안됨 (Get만 허용 및 보안취약, 동기호출안됨 등 문제있음)
- HTML5 부터 다른 도메인 간 통신이 가능한 스펙이 추가되었는데 그게 CORS
- 4가지 방법이 있으므로 골라쓰면 됨
