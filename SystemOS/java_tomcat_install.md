### linux에 java, tomcat을 설치해보자.

---

> #### CentOS7 기반, rpm과 yum 없이 깔아볼까요?

##### 1. java 설치

`java -version` 으로 자바 설치 여부 확인
`rpm -qa | grep java`로도 확인 가능.

설치된 자바가 없으면 `wget 다운로드할 주소` 로 jdk.tar.gz를 받고
`tar -xvf jdk.tar.gz`로 압축 해제

`vi etc/profile`을 실행하여, 최 하단에

```
JAVA_HOME=/opt/jdk1.8.0_131
CLASSPATH=.:$JAVA_HOME/lib/tools.jar
PATH=$JAVA_HOME/bin
export JAVA_HOME CLASSPATH PATH
```

을 입력해 변경한 후.

`source /etc/profile`로 profile 변경내역을 적용

다시 한번 `java -version` 으로 버전 확인

##### 2. tomcat 설치

` wget http://apache.tt.co.kr/tomcat/tomcat-8/v8.0.26/bin/apache-tomcat-8.0.26.tar.gz `로 톰캣 다운로드

`tar -zxf /다운받은 위치/apache-tomcat-8.0.26.tar.gz `로 압축 해제

톰캣 폴더에 사용자 계정 권한 부여 `chown userid:/다운받은 위치/apache-tomcat-8.0.26-src`

톰캣 포트 변경

`vi /다운받은 위치/apache-tomcat-8.0.26-src/conf/server.xml`

```
<Connector port="변경할 포트" protocol="HTTP/1.1"
               connectionTimeout="20000"
               redirectPort="8443" URIEncoding="UTF-8" />
```

방화벽 설정

`firewall-cmd --permanent --zone=public --add-port=변경할 포트/tcp`

`firewall-cmd --reload`

톰캣 시작 / 정지
`apache-tomcat-8.0.26-src/bin/startup.sh `
`apache-tomcat-8.0.26-src/bin/shutdown.sh`

접속

`http://아이피주소:포트번호`