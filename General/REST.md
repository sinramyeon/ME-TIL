# REST

OpenAPI와 함께 떠오른 REST API(RestFul) 도대체 근데 이게 뭐라고 설명을 해야 할까.

## REPRESENTATIONAL STATE TRANSFER

REST를 개발하고 써놓고 사실 REST 가 뭐냐 하면 어~그 RestFul한 ~ 어쩌구 외에 설명이 안 돼서 이번 기회에 REST를 확실히 정리하려고 한다.

단순히 말하자면 REST는 웹 상의 자료를 HTTP위에서 SOAP이나 쿠키를 통한 세션 트랙킹 같은 별도의 전송 계층 없이 전송하기 위한 아주 간단한 인터페이스를 말한다. 즉, 자료 전송용 틀 같은 것이다.

## REST의 구성

- 자원(RESOURCE) - URI
- 행위(Verb) - HTTP METHOD
- 표현(Representations)

REST는 제어할 자원을 표시한 `URI` 로 수행할 행위인 `HTTP Method` 를 보내어 그 결과를 `표현` 한다. 뭔소리냐면 보내고싶은 자료를 URI라고 하고 그걸 HTTP method를 통해 보내고 결과는 틀에 맞춰서 보여준다는 뜻이다.

### URI

`Uniform Resource Identifier`

통합 자원 식별자(Uniform Resource Identifier, URI)는 인터넷에 있는 자원을 나타내는 유일한 주소이다. URI의 존재는 인터넷에서 요구되는 기본조건으로서 인터넷 프로토콜에 항상 붙어 다닌다.

...가 정의인데, 대강 URI 가 뭐냐면 유일한 자원 주소이다. html이던 뭐던.. 딱 하나 있는 자료이다. naver.com/webtoon/hello/1.html 같은 것들을 URI라고 한다. URL하고 뭐가 다르냐면 URL은 `Uniform Resource Locator`고 URI는 *Identifier*이다. URL이 URI 하위개념이라고 보면 된다.

REST에서는 URI는 반드시 자원을 가리켜야한다. 그러니까 예를 들어 `post/delete/1`로 1번 글을 지우는 REST API를 생각해 보자.

URI패턴이 뭔 정확한 자료로 뭘 하는지 확실히 명시한다.

#### URI설계시 주의점

그냥 아무렇게나 `/post/sakje/1`이라고 하면 안된다. `soinsu_bunhea()`나 `var gilee` 같은 꼴이다. 제발 그러면 안된다.

1) 슬래시 구분자(/)는 계층 관계를 나타내는 데 사용

```
    http://restapi.example.com/houses/apartments
    http://restapi.example.com/animals/mammals/whales
```

2) URI 마지막 문자로 슬래시(/)를 포함하지 않는다.

URI에 포함되는 모든 글자는 리소스의 유일한 식별자로 사용되어야 하며 URI가 다르다는 것은 리소스가 다르다는 것이고, 역으로 리소스가 다르면 URI도 달라져야 합니다. REST API는 분명한 URI를 만들어 통신을 해야 하기 때문에 혼동을 주지 않도록 URI 경로의 마지막에는 슬래시(/)를 사용하지 않습니다.

```
    http://restapi.example.com/houses/apartments/ (X)
    http://restapi.example.com/houses/apartments  (0)
```

3) 하이픈(-)은 URI 가독성을 높이는데 사용

URI를 쉽게 읽고 해석하기 위해, 불가피하게 긴 URI경로를 사용하게 된다면 하이픈을 사용해 가독성을 높일 수 있습니다.

4) 밑줄(_)은 URI에 사용하지 않는다.

글꼴에 따라 다르긴 하지만 밑줄은 보기 어렵거나 밑줄 때문에 문자가 가려지기도 합니다. 이런 문제를 피하기 위해 밑줄 대신 하이픈(-)을 사용하는 것이 좋습니다.(가독성)

5) URI 경로에는 소문자가 적합하다.

URI 경로에 대문자 사용은 피하도록 해야 합니다. 대소문자에 따라 다른 리소스로 인식하게 되기 때문입니다. RFC 3986(URI 문법 형식)은 URI 스키마와 호스트를 제외하고는 대소문자를 구별하도록 규정하기 때문이지요.

    
6) 파일 확장자는 URI에 포함시키지 않는다.

```
    http://restapi.example.com/members/soccer/345/photo.jpg (X)
```

REST API에서는 메시지 바디 내용의 포맷을 나타내기 위한 파일 확장자를 URI 안에 포함시키지 않습니다. Accept header를 사용하도록 합시다.

```
    GET / members/soccer/345/photo HTTP/1.1 Host: restapi.example.com Accept: image/jpg
```

### HTTP METHOD

REST에서는 어떤 자원에 행할 수 있는 행위를 네 가지로 나타내며 이를 CRUD 라고 한다.
CRUD는 Create, Read (Retrieve), Update, Delete 의 첫 글자를 각각 따온 것이다.
CRUD는 각각 HTTP Method 의 POST, GET, PUT, DELETE 에 대응된다. PUT을 유의하자.

- POST : POST를 통해 해당 URI를 요청하면 리소스를 생성한다.
- GET : GET를 통해 해당 리소스를 조회하고 해당 도큐먼트에 대한 자세한 정보를 가져온다.
- PUT : PUT를 통해 해당 리소스를 수정한다.
- DELETE : DELETE를 통해 리소스를 삭제한다.


### Representation

자원의 URI에 특정 행위를 요청하면 그 결과로 Representation 을 응답 받는다.
어떤 자원의 Representation은 html, xml, text, json, rss 등 다양한 형태로 표현될 수 있다. postman을 사용해서 곧잘 해보았을 것이다. 이걸 해놓고도 REST가 뭔지 몰랐다니 환멸이 난다.

응답 상태 코드는 아래와 같이 온다. 잘 설계된 REST API는 URI만 잘 설계된 것이 아닌 그 리소스에 대한 응답을 잘 내어주는 것까지 포함되어야 한다.

200 - 정상
201 - 클라이언트가 어떠한 리소스 생성을 요청, 해당 리소스가 성공적으로 생성됨(POST를 통한 리소스 생성 작업 시)
400	- 클라이언트의 요청이 부적절 할 경우 사용하는 응답 코드
401	- 클라이언트가 인증되지 않은 상태에서 보호된 리소스를 요청했을 때 사용하는 응답 코드(로그인 하지 않은 유저가 로그인 했을 때, 요청 가능한 리소스를 요청했을 때)
403 - 유저 인증상태와 관계 없이 응답하고 싶지 않은 리소스를 클라이언트가 요청했을 때 사용하는 응답 코드(403 보다는 400이나 404를 사용할 것을 권고. 403 자체가 리소스가 존재한다는 뜻이기 때문에)
405	- 클라이언트가 요청한 리소스에서는 사용 불가능한 Method를 이용했을 경우 사용하는 응답 코드
301 - 클라이언트가 요청한 리소스에 대한 URI가 변경 되었을 때 사용하는 응답 코드
(응답 시 Location header에 변경된 URI를 적어 줘야 합니다.)
500	- 서버에 문제가 있을 경우 사용하는 응답 코드


---

### 그런데 REST는 왜 쓰는지

가장 큰 장점인 *쉽다* 라는 이유부터, 표준이 없다는 단점에도 불구하고 표준으로 이용된다는 것. 웹의 장점을 최대한 활용할 수 있는 네트워크 기반의 아키텍쳐이므로.

### SOAP는 무엇인지

정의를 찾아보면,

SOAP

> SOAP(Simple Object Access Protocol)은 일반적으로 널리 알려진 HTTP, HTTPS, SMTP 등을 통해 XML 기반의 메시지를 컴퓨터 네트워크 상에서 교환하는 프로토콜이다.

REST

> REST(Representational State Transfer)는 월드 와이드 웹과 같은 분산 하이퍼미디어 시스템을 위한 소프트웨어 아키텍처의 한 형식이다. REST 원리를 따르는 시스템은 종종 RESTful이란 용어로 지칭된다


라고 나온다. REST에 대해선 방금 배웠으니, SOAP는 간단히 설명하면 웹서비스 내 모든 데이터를 XML로 표현한것이라고 생각하면 된다. 개념적으로 REST보다 어려울 수 있지만 분산 컴퓨팅 환경을 다루기 위한 설계에 걸맞는다.


