### 프로토콜의 이해

---


[참고한 링크](http://asfirstalways.tistory.com/85)

---

#### 프로토콜이란?

원거리 통신 장비 사이 메시지를 주고받는 양식과 규칙이다.

#### TCP / IP?
##### Transmission Control Protocol / Internet Protocol
* TCP : 정보 패킷 차원에서 다른 인터넷 노드와 메시지를 상호 교환하는데 필요한 규칙을 사용
* IP : 인터넷 주소 차원에서 메시지를 보내고 맏는게 규칙을 사용

#### TCP / UDP?
##### TCP : 전송 과정을 컨트롤하는 프로토콜

1. 높은 신뢰성
2. 가상 회선 연결 방식

##### UDP : 사용자 데이터그램형, 비연결형 전송 프로토콜

1. 비연결형
2. 비신뢰성, 오류검출

#### HTTP
##### HyperText Transfer Protocol
서버가 html로 만들어진 문서를 유저에게 잘 보여주는 것을 목적으로 하는 프로토콜
TCP와 UDP 사용, 클라이언트와 서버 사이 이루어지는 리퀘스트 / 리스폰스 프로토콜

#### 웹 소켓
HTTP 단점을 보완하기 위해 등장. 웹을 통해 클라이언트와 서버 간 신속하고 보안이 유지된 양방향 통신 가능.
HTTP가 적합치 않은 트래픽이 높고 지연시간이 낮은 환경에서 유용하다.

#### 대표적 포트들

* 21 FTP (파일 전송 프로토콜)
* 22 SSH (시큐어 쉘):hamster:
* 23 Telnet(암호화되지 않은 텍스트 통신)
* 25 SMTP (메일 전송 프로토콜)
* 53 DNS (Domain Name System)
* 80 HTTP (HyperText Transfer Protocol)
* 110 POP3 (전자우편 가져오기)
* 115 SFTP
* 135 RPC
* 139 NetBIOS
* 143 IMAP (이메일 가져오기)
* 194 IRC (채팅)
* 443 SSL (암호화 전송)
* 445 SMB (파일 공유)
* 1433 MSSQL
* 3306 MySQL
* 3389 Remote Desktop
* 5632 Pcanywhere
* 5900 VNC

:hamster: Secure Shell, SSH : 네트워크 상 다른 컴퓨터에 로그인하거나 원격 시스템에서 ㅁ여령을 실앻하고 다른 시스템으로 파일을 복사할 수 있도록 해 준다.