## Tomcat 설정

---

#### url 설정

myApp이라는 wepapp을 myApp.war로 배포시, 톰캣에서는 기본적으로 http://hostname:8080/myApp 같은 url로 접근한다.
이 접근루트를 http://hostname:8080 으로 변경하자.

그러려면 아래의 방법이 있다.

1. war rename

myApp.war를 ROOT.war 로 이름을 바꾼다.
maven에서는 pom.xml을 아래와 같이 바꾼다.
```
<build>
  <finalName>ROOT.war</finalName>
</build>
```

2. server.xml 변경

tomcat의 context root 설정을 변경한다.
```
<Host name="localhost" appBase="webapps" unpackWARs="true" autoDeploy="false">
  <Context path="" docBase="myApp" reloadable="false > </Context>
</Host>
```

3. apps를 webapps 외부에 위치시키기

app을 webapps밖에다 설정하고 context path를 수정한다.

conf/server.xml 수정
```

<Host name="localhost" appBase="webapps" unpackWARs="true" autoDeploy="false" deployOnStartup="false" >
 <Context path="" docBase="${catalina.home}/myApp" reloadable="false" > </Context>
</Host>
```

---

#### Admin, manager 애플리케이션 설정

http://localhost:8080/admin, http://localhost:8080/manager/html 로 접속하면, admin과 manager 기능이 시작된다.
위 로그인 정보를 설정하는 과정은 다음과 같다.

1. CATALINA_HOME/cont/tomcat-users.xml 파일을 연다
2. Admin 어플리케이션에 접근하기 위해서는 admin role, Manager 애플리케이션에 접근하기 위해서는 manager role이 필요하다.
3. 같은 계정이 두개 다 가질 수도 있다.

```
<?xml version='1.0' encoding='utf-8'?>
<tomcat-users>
  <role rolename="tomcat"/>
  <role rolename="role1"/>
  <role rolename="manager"/>
  <role rolename="admin"/>
  <user username="tomcat" password="tomcat" roles="tomcat"/>
  <user username="role1" password="tomcat" roles="role1"/>
  <user username="both" password="tomcat" roles="tomcat,role1"/>
  <user username="admin" password="admin" roles="admin,manager"/>
</tomcat-users>
```

로 username, password, role을 설정 해 준다.
이 내용대로 로그인 할 수 있다.

