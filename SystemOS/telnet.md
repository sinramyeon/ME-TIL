### telnet 명령어

---

제어판 > 프로그램 및 기능 > Windows 기능 켜기 / 끄기 > 텔넷 클라이언트 체크

`telnet 대상장비 ip port 번호` 로 대상 장비 특정 포트가 열려있는지 여부 확인

1. 포트가 막혀있을 때

`telnet 192.163.2.388 9999`
`Trying 192.163.2.388...`

2. 포트 오픈은 되어잇는데 해상 서비스가 없음

`telnet 192.163.2.388 9999`
`Trying 192.163.2.388...`
`telnet : unable to connect to remote host : connecton refused`

3. 포트 오픈 및 해당 서비스도 정상

`telnet 192.163.2.388 9999`
`Trying 192.163.2.388...`
`Connected to 192.163.2.388`

