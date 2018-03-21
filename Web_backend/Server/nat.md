# NAT란?

> NAT : Network Address Translation

NAT에 대하여 알아보자. 보통 natsd서버니 nat이라니 불리는데, 그 정체가 무엇일까?

### NAT의 정의

Network address translation - 네트워크 주소 변환, 줄여서 NAT는 IP패킷의 TCP UDP 포트 숫자랑 소스 그리고 목적지 IP주소를 재기록하면서 라우터를 통하여 네트워크 트래픽을 주고 받는 기술이다. 즉, 주소 하나당 컴퓨터 하나밖에 인터넷이 안 되니, 여러대 컴퓨터로 쓰려고 ip를 배분해주는 기술이 곧 NAT이다.

> Private Network에 위치하는 단말이 Public Network(인터넷)과 통신이 가능하도록 상호 간에 연결 시켜 주는 기능

### NAT의 동작 원리

NAT는 1:1의  주소  매핑을  수행하기  때문에 NAT라우터로  들어온 inside->outside 패킷(또는  그  반대)만이  주소전환의  대상이  된다. 간단히  설명하자면 IP의  헤더  부분을  체크하여 NAT 테이블에  의해  해당  주소로  바꾼다음checksum을  다시  계산하여 IP의  헤더를  바꾸는  방법으로  동작한다. Application layer에서  까지 NAT의  주소전환이  반영이  되려면 NAT는 IP 주소의  참조내용을  담고  있는 application 데이터  부분을  새로운  주소로  변환해야  한다.

**\[출처\]**  [NAT란??](http://blog.naver.com/changjaeho/120008095774)|**작성자**  [호기심](http://blog.naver.com/changjaeho)

### NAT의 도입 장점

(1) 사설 IP 주소를 가진 여러대의 단말들이 하나의 공인 IP 주소를 통해 인터넷과 연결됨으로써 공인 IP 주소를 절약
(2) 사내망을 사설 IP 주소화 하여 외부로 부터의 친입/공격을 차단할 수 있습니다. (방화벽의 개념).
(3) ISP변경에 따른 내부 IP변경을 최소화합니다.

### NAT의 도입 단점

(1) end-to-end간의  추적(IP trace)어려워 진다
(2) NAT 라우터를  거치는  모든  패킷을 scan하므로 switching path delay가  커진다

#### natsd 서버? gnatsd?

NATS를 위한 다양한 오픈소스 중 낫츠디 라고 불리우는 [gnatsd](https://github.com/nats-io/gnatsd) 를 뜻하는 것이다.
