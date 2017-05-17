##### oracle11g 설치시 'http:/.127.0.0.1:%HTTPPORT%/apex/f?p=4950'

[참고자료](http://stackoverflow.com/questions/27631400/windows-cannot-find-http-127-0-0-1httpport-apex-fp-4950-make-sure-you-t)

---
> 오라클 시작 페이지 접속 하기

1. http://127.0.0.1:8080/apex/f?p=4950:1:1615033681376854 로 접속하세요.
2. ID는 system, 비밀번호는 오라클 설치 시 정한 비밀번호로 로그인 하세요.
3. 빨간색 오라클 시작페이지로 들어가고 주소가 바뀌었습니다!

> 정보 수정하기
1. 오라클 설치 디렉토리 > oracle > app > oracle > product > 11.2.0 > server 로 들어가 보세요.
2. Get_Started.html 을 우클릭해서 속성에 가서 url을 위에서 바뀐 주소로 변경해 보세요.
3. 해결!