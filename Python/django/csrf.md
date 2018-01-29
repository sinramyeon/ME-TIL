# CSRF

cross site request fogery. 사이트간 요청 위조라고 부르는데, 사용자가 자기가 원치 않았지만 공격자가 의도한 행위대로 특정 사이트에 요청하게 되는 공격이다.

이러한 취약점을 막는 가장 기본적인 방법은 GET대신 POST를 사용하는 것이다. POST를 쓸 때도 hidden필드에 임의의 키값을 전달하고 그 키값이 맞는지 매번 확인하는게 좋다.

django에서는 1.2부터 자동 CSRF방지 기능을 제공한다.

사용법을 알아보자.

## django csrf token

CSRF 기능은 미들웨어에서 default 로 켜져 있다. 만약 오버라이딩 할 거라면 `django.middleware.csrf.CsrfViewMiddleware` 가 다른 어떤 미들웨어보다 앞에 오도록 신경써서 하자.

기능을 끄는 걸 추천하지 않지만 정 끄고 싶으면 뷰마다 `csrf_protect()` 를 사용할 수 있다.


```
from django.views.decorators.csrf import csrf_protect
from django.template import RequestContext

@csrf_protect
def view_function(request) :
    # 내용...
    context = {}
    return render_to_response("index.html", context, context_instance=RequestContext(request))
```

일반적으로는 템플릿에서 어느 폼이건 POST 를 사용할거면 `csrf_token` 태그를 폼안에 넣는다.

이런식으로 넣으면 된다.
```
<form action="" method="post">{% csrf_token %}
```

내부 url로 이어지는 POST폼안에 넣으면 CSRF 토큰 유출 위험이 있으니 그러지 않는다.

해당 뷰 기능에서 {% csrf_token %}이 올바르게 작동하도록 응답을 렌더링하는 데 RequestContext가 사용되는지 확인해야 한다.  render() 함수, 일반 뷰나 contrib 응용 프로그램을 사용하는 경우에는 모두 RequestContext를 사용하기 때문이다.

만약 특정 뷰에 대해 csrf를 적용하지 않으려면 아래와 같이 csrf_exampt 장식자를 쓴다.

```
from django.view.decorators.csrf import csrf_exempt

@csrf_exempt
def view_function(request) :
    return HttpResponse("")
```

