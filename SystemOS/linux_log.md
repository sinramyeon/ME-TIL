### 리눅스 로그

---


[참조자료](http://mslee89.tistory.com/42)
[참조자료2](http://windfree.tistory.com/40)

---

#### 리눅스 로그 종류

1. secure 로그
- 로그파일 위치 : /var/log/secure
- telnet, ssh접속에 대한 유저 로그인 인증 기록

2. dmesg 로그
- 로그파일 위치 : /var/log/dmseg
- 부팅시 시스템에 의해 기록되는 로그

3. message 로그
- 로그파일 위치 : /var/log/message
- syslog에 의한 로그인, 설정, 장치 정보의 전체적 로그를 기록

4. utmp 로그
- 로그파일 위치 : /var/rum/utmp
- 현재 시스템에 로그인 한 각 사용자의 상태를 저장하는 바이너리 파일.
- w, who, users, finger 등의 명령어로 내용 확인 가능

5. wtmp 로그
- 로그파일 위치 : /var/log/wtmp
- 로그인, 로그아웃, 시스템의 재부팅에 대한 정보가 담겨있습니다.
- 바이너리 파일로 되어 있으며 last 명령어로 내용을 확인

6. lastlog
- 로그파일 위치 : /var/log/lastlog
- 계정 사용자들이 마지막으로 로그인한 정보. 바이너리 파일로 되어 있으며 lastlog 명령어로 확인 가능

7. boot.log
- 로그파일 위치 : /var/log/boot.log
- 부팅시 서비스 데몬들의 실행 상태를 기록하는 로그이다.

8. cron 로그
- 로그파일 위치 : /var/log/cron
- cron에 예약한 작업이 정상적으로 실행되었는지에 관한 로그

9. anaconda 로그
- 로그파일 위치 : /var/log/anaconda.log
- 리눅스 설치시 install 된 패키지나 과정에 대한 로그 기록

10. su 로그
- 로그파일 위치 : /var/log/sulog한 로그 기록
- su 명령어를 통한 관리자 권한을 사용

11. history 로그
- 로그파일 위치 : /root/.bash_history(root유저일 경우)
- 입력한 실행명령어 목록이 저장되는 로그 

----

#### 로그 보기

`tail [option]....[file]...`
기본 출력은 파일의 마지막 10줄을 보여줌

옵션
- f : 파일의 마지막 10라인을 실시간으로 계속해서 출력
- F : 파일 변동 시 실시간으로 보여주되 로그파일처럼 특정 시간이 지난 후 파일이 변하게 되면 새로운 파일을 오픈해서 보여줌(다시 명령 실행할 필요가 없음)
- n : n만큼의 라인을 출력

예제
`tail -n 20 anaconda-ks.cfg`
anaconda-ks.cfg파일의 마지막부터 20줄을 출력

예제2

`tail -f /var/log/messages`
messages.log파일을 실시간으로 화면에 출력