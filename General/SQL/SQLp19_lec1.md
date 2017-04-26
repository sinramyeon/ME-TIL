#### SQL -p19 강의

---

##### Oracle의 DB Link 생성 및 사용법

DB Link란? : 클라이언트 또는 현재의 데이터베이스에서 네트워크 상 다른 데이터베이스에 접속하기 위한 접속 설정

` CREATE [PUBLIC] DATABASE LINK <link_name> CONNECT TO <username> IDENTIFIED BY <password> USING <원격 db alias>`

- PUBLIC : 공용 데이터베이스 링크를 생성함.
- PUBLIC 을 사용하지 않을 시 링크 생성한 자만 사용 가능
- link_name : 데이터베이스 링크의 이름 지정
- 원격 db alias : 네트워크 접속에 사용할 오라클 데이터베이스 네트워크 서비스명

---

> ORA-12154: TNS:could not resolve the connect identifier specified. 해당 오류 발생시 사용
>


```
CREATE PUBLIC DATABASE LINK LINK_NAME
 CONNECT TO <연결하고자 하는 user>
 IDENTIFIED BY <연결하고자 하는 user password>
  USING '(DESCRIPTION =   
                (ADDRESS_LIST =   
                  (ADDRESS = (PROTOCOL = TCP)(HOST=127.0.0.1)(PORT = 1521))  
                )  
                (CONNECT_DATA =   
                  (SERVICE_NAME = ORCL)  
                )  
              )'
```

----

사용 방법
`SELECT COUNT(*) FROM TABLE_NAME@DATA_LINK`

---

생성시 권한 오류

SYSDBA 계정으로 권한 추가 : PUBLIC 명 시 여부 체크

```
//PUBLIC DB LINK 생성  권한
GRANT CREATE PUBLIC DATABASE LINK TO SCOTT;

//생성권한 제거
REVOKE CREATE PUBLIC DATABASE LINK FROM SCOTT;

//PUBLIC 링크 제거 권한
GRANT DROP PUBLIC DATABASE LINK TO SCOTT;

//PUBLIC 링크 제거
DROP PUBLIC DATABASE LINK "링크명" ;
```

---

DB 링크 사용시 에러 (INVALID USERNAME 등)

1. 11g에서 대소문자를 구분하지 않도록 설정
`ALTER SYSTEM SET SEC_CASE_SENSIVITVE_LOGON = FALSE`

2. DB링크 생성 시 접속 사용자 패스워드 부분에 쌍따옴표를 붙여본다
`CREATE DATABASE LINK DBLINK CONNECT TO SUNSHINY INDENTIFIED BY "PASS" USING '<TNS정의명>'`