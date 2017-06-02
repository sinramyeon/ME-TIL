#### Jboss / WildFly

Redhat사의 제품으로 무료 Jboss는 7까지, 그 이후 버전은 WildFly로 무료 제공된다. Jboss EAP는 유상이다.
[참고](http://opennaru.tistory.com/44)

- auto deploy 설정

`wildfly-8.2.0.Final\standalone\configuration\standalone.xml`
내용을

```
<subsystem xmlns="urn:jboss:domain:deployment-scanner:1.1">
<deployment-scanner path="deployments" relative-to="jboss.server.base.dir" scan-interval="5000" auto-deploy-exploded="true"/>
</subsystem>
```
로 바꾼다.



- 이슈 해결

> egovframe 3.5.1버전에서 Jboss7버전 구동시 구동이 되지 않는 에러가 발생하며 spring-modules-validation-0.9.jar 관련 에러 메시지가 뜰 때


```
spring-modules-validation-0.9.jar 안에 있는 /META-INF/valang.tld 파일을 다음과 같이 편집해서 적용하시면 되실 것 같습니다.


org.springmodules.validation.valang.javascript.taglib.ValangCodebaseTag

에 대하여 <body-content>None</body-content>

만약 이게 안될시

<body-content>empty</body-content> 로 변경
```

> Jboss7은 jdk 1.7까지만 지원하므로 1.8이상을 쓰려면 WildFly 8부터 사용해야 한다

- WildFly 설정방법

1. localhost:8080 으로 접속하면 와일드플라이 접속화면이 뜬다.

![접속화면]("wildfly_start.PNG")

2. Administration Console에 들어가서 유저 추가를 해준다.

![접속화면]("wildfly_start.PNG")

wildfly-8.2.0.Final\bin 안에 add-user.bat 을 켭니다.
`add-user bat start`

user type
a : 매니지먼트 유저(관리자)
b : 어플리케이션 유저(사용자) 인데 종류를 고른 후,
username과 password, group(이건 필수가 아님)을 골라
유저를 추가하고, 다시 서버를 재접속합니다.

*비밀번호에 알파벳, 숫자, 특수문자가 다 들어가야 함 ㅡㅡ*

3. http://127.0.0.1:9990 로 접속해 admin 페이지를 확인 가능하다.

- WildFly db 설정방법


![db]("wildfly_db2.PNG")

이클립스 내 database.properties를 설정하고,
datasource 설정 context 파일 내 jndi를 선언한다.


![db]("wildfly_db1.PNG")

http://127.0.0.1:9990 화면의 Configuration > Connector > Datasources > Add로 들어갑니다.

![db]("wildfly_db3.PNG")

![db]("wildfly_db4.PNG")

위와 같은 순서로 내용을 입력 후 추가하고, Enable 해주면 됨
