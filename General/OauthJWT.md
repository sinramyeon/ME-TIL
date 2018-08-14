# OAuth

> OAuth - 인증 수단

보통 구글로 가입...페이스북으로 가입 등의 기능이 OAuth라고 알고 있을 것이다. 물론 틀린 말은 아니다! 그렇지만 OAuth란게 당최 어떤 것인지데 대해 좀더 자세히 알아보자.

> 외부 서비스에서도 인증을 가능하게 하고, 그 서비스의 API를 이용하게 하는 기능

간단하게 인증(Authentication)과 권한(Authorization)을 획득하는 것이 곧 OAuth이다. 내가 직접 회원가입 탈퇴 로그인을 갖다 만드는것도 좋지만 얼굴책, 네입어, 고글 등에서 이미 만들어놓은 API를 통해 인증을 하는 것이다.

대강 OAuth의 역할을 네 가지로 나누어진다.

1. Resource Owner : 사용자
2. Resource Server : API 서버
3. Client : 서드파티 어플리케이션
4. Authorization Server : 인증서버

대략적 워크플로우는 아래와 같다.

![](https://i.imgur.com/7T48KvR.png)

OAuth2는 OAuth1 에서 인증을 강화하고 구현을 간소화한 차이가 있다.

대략적 차이는 아래와 같다.

- 인증 절차 간소화
기능의 단순화, 기능과 규모의 확장성 등을 지원하기 위해 만들어 졌다.
기존의 OAuth1.0은 디지털 서명 기반이었지만 OAuth2.0의 암호화는 https에 맡김으로써 복잡한 디지털 서명에관한 로직을 요구하지 않기때문에 구현 자체가 개발자입장에서 쉬워짐.

- 용어 변경
Resource Owner : 사용자 (1.0 User해당)
Resource Server : REST API 서버 (1.0 Protected Resource)
Authorization Server : 인증서버 (API 서버와 같을 수도 있음)(1.0 Service Provider)
Client : 써드파티 어플리케이션 (1.0 Service Provider 해당)

- Resource Server와 Authorization Server서버의 분리
커다란 서비스는 인증 서버를 분리하거나 다중화 할 수 있어야 함.
Authorization Server의 역할을 명확히 함.

- 다양한 인증 방식(Grant_type)
Authorization Code Grant
Implicit Grant
Resource Owner Password Credentials Grant
Client Credentials Grant
Device Code Grant
Refresh Token Grant

Go에서의 구현 예제이다.

```
import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	newappengine "google.golang.org/appengine"
	newurlfetch "google.golang.org/appengine/urlfetch"

	"appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var c appengine.Context = appengine.NewContext(r)
	c.Infof("Logging a message with the old package")

	var ctx context.Context = newappengine.NewContext(r)
	client := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(ctx, "scope"),
			Base:   &newurlfetch.Transport{Context: ctx},
		},
	}
	client.Get("...")
}
```

---

# JWT

그러면 또 어딘가에서 자주 들어본 JWT 에 대해서 알아보자. JWT란 JSON자체를 인증 토큰으로 사용하는 것이다.

> JSON Web Token

무슨 뜻이냐면, 이런 JSON을 토큰 그 자체로 사용한다는 뜻이다.

```
{ "name" : "seolhwa" }
```

그러면 따로 사용자 정보를 조회할 필요가 없다. 왜냐면 토큰 자체가 정보값이니까.

> JWT는 JSON 객체에 요구사항을 작성하고, 어떠한 암호화 방식을 사용해서 문자열로 인코딩을 한 후, HTTP header에 추가함으로써 사용자 인증을 요청한다.

대략적 워크플로우는 아래와 같다.

![](https://cdn.auth0.com/content/jwt/jwt-diagram.png)

먼저 사용자가 id와 password를 입력하여 로그인을 시도 -> 서버는 요청을 확인하고 secret key를 통해 Access token을 발급 -> JWT가 요구되는 API를 요청할 때는 클라이언트가 Authorization header에 Access token을 담아서 보냄 -> 서버는 JWT Signature를 체크하고 Payload로부터 user 정보를 확인해 데이터를 리턴

JWT가 가진 특징은 아래와 같다.

1. 정보가 담긴 데이터( JSON 객체 )를 암호화 해서 HTTP 헤더에 추가 (보안성 증가)
2. 권한을 부여하기 위한 데이터가 JWT안에 모두 담겨있음 (OAuth 처럼 인증 서버에서 토큰에 대한 정보를 찾을 필요가 없음)

대신, 토큰이 곧 정보이므로 토큰 유효시간을 항상 설정해주는게 좋다.
또 토큰 보안을 위해 HMAC기법(데이터 암호화 해싱)을 사용하는게 좋다.

JWT의 구조와 생성법을 알아보자.

![](https://t1.daumcdn.net/cfile/tistory/99E82B395A655D1131)

JWT공홈에 가면 이런식으로 이루어진다는 것을 쉽게 만들어보고 확인할 수 있다.

- 헤더 ( header )

헤더에는 typ와 alg 속성을 명시합니다.

- typ : 토큰의 타입을 명시합니다.
- alg : 해싱 알고리즘을 명시하는데, 이 알고리즘은 서버에서 토큰을 검증할 때 사용되는 signature에서 사용됩니다.


- 내용 ( payload )

내용에는 토큰에 대한 정보를 작성합니다.

정보는 속성, 값으로 표현되고 이를 claim이라 하는데, registered, public, private claims을 작성할 수 있습니다.


registered claim : 미리 정의된 claim으로써, 토큰에 대한 정보를 작성합니다.

iss  : 토큰 발급자

exp : 토큰의 만료시간

sub : 토큰 제목

aud : 토큰 대상자

그 밖에 여러 claim들이 있으며, 더 많은 정보는 JWT의 공식문서를 참고해주세요

public claim : 공개적인 claim을 명시하는데, 충돌방지를 위해 URI 형식으로 작성합니다.

private claim : 서버와 클라이언트가 협의한 claim을 명시합니다.


- 서명 ( signature )

서명에서는 헤더(header)의 인코딩 값과 내용(payload)의 인코딩 값을 "."으로 연결하여 합친 후 비밀키로 해싱합니다. 일종의 암호화하는 작업이라고 생각하시면 됩니다.


이제 header, payload, signature의 각 값들을 "."으로 합치면 하나의 JWT가 생성됩니다.

이렇게 생성된 JWT를 HTTP header에 추가하여 서버에 요청을 하면 서버에서는 이를 디코딩하여 사용자 인증을 하게 됩니다.

[출처] http://victorydntmd.tistory.com/115

Go에서 iris 프레임워크 내에서 jwt 인증을 구현했다.

```
package main

import (
	"github.com/kataras/iris"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func myHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("%s", user.Signature) // claim에 들어가는 유저 사인
}

func main() {
	app := iris.New()

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil // 여기에 시크릿 키를 넣는다.
		},
	})

	app.Use(jwtHandler.Serve) // 이런식으로 사용한다.

	app.Get("/ping", myHandler)
	app.Run(iris.Addr("localhost:3001"))
} 
```
