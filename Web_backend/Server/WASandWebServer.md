### WAS와 Web 서버의 차이는?

<hr/>

[참고한 링크](http://sungbine.github.io/tech/post/2015/02/15/tomcat%EA%B3%BC%20apache%EC%9D%98%20%EC%97%B0%EB%8F%99.html)

<hr/>

##### WAS서버
###### 웹 서버와 웹 컨테이너의 결합. 동적인 데이터를 처리한다.
도입효과 : 안정된 시스템 구성, DB성능 보장, 비용 절감
기술 표준 : J2EE(자바 기반 분산객체 아키텍쳐)

##### Web서버
###### html, jpg, gif등 웹페이지 문서에 반응해 웹페이지를 보여준다. 정적인 데이터를 처리하면 빠르고 안정적이다.

##### Apache는 뭐지?
apache 서버는 아파지에서 만든 http 웹 서버이다.

##### Tomcat은?
WAS의 한 종류. 아파치와의 차이는? *컨테이너가 가능하냐 안하냐.*

##### 두 서버는 그럼 언제 뭘 쓰지?
두 서버 목적 차이가 있으니 두 서버를 연동해서 쓴다.
사용자 요청은 http 웹 서버를 통해 받고 내부 프로그램은 was를 통해 처리하는 식으로...