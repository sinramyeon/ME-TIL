### rpm

---

#### 1. rpm 이 뭐양

rpm, yum이란, 리눅스의 패키지 인스톨 프로그램이자 인스톨 파일.

redhat package manager로서 setup.exe같은거라고 생각하면 편하다.

#### 2. rpm 설치

설치
`rpm -ivh 패키지명`

확인
`rpm -qa | grep 패키지명`

제거
`rpm -ev 패키지명`

패키지 의존성 무시하고 삭제하기

```
error: Failed dependencies:
java is needed by (installed) jna-3.2.4-2.el6.x86_64
java >= 1.5.0 is needed by (installed) libvirt-java-0.4.7-1.el6.noarch
```
이런 에러가 뜰 때는

`rpm -e --nodeps 패키지명` 으로 삭제.

#### 3. rpm 관련 명령어

파일이 속한 패키지 찾기
`rpm -qf 파일`

rpm 패키지 정보 보기
`rpm -qi 설치된 패키지명`
`rpm -qip 파일명.rpm`
