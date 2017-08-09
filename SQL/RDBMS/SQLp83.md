### SQL -P83

---

##### 시스템 권한 System Privilege

데이터베이스 내 특정 Object에 대한 권한이 아니라 Object의 타입 (TABLE, VIEW, CLUSTER 등)에 대해 어떤 작업을 명시한 권한.
DBA가 부여.

`GRANT ROLE TO PUBLIC WITH ADMIN OPTION`

- PUBLIC 모든 사용자에게 권한 부여
- WITH ADMIN OPTION 권한을 부여받는 사람도 다른 사용자에게 해당 권한을 줄 수 있게 함

`GRANT CREATE SESSION, CREATE TABLE TO hero`

- CREATE SESSION 데이터베이스와 연결할 수 있도록 하는 시스템 권한
- CREATE TABLE 테이블 생성 권한

`GRANT ALTER ANY TABLE TO hero`
hero에게 어떤 테이블도 변경할 수 있는 권한 줌

`DESCRIBE dba_sys_privs`

EX)

```
SQL> SELECT *  
  2  FROM sys.dba_sys_privs
  3  WHERE grantee IN ('SCOTT','LAURA');
```

시스템 권한을 받은 사람 중 scott, laura의 권한정보 출력

<br>

---

<br>

##### 오브젝트 권한 Object Privilege

Object권한의 형태

object priileges | table | view | sequence | procedure | snapshots
--- | --- | -- | -- | -- | --
alter | o | x | o | x | x
delete | o | o | x | x | x
execute | x | x| x | o | x
index | o | x | x | x | x
insert | o | o | x | x | x
references | o | x | x | x
select | o | o | o | x | o
update | o | o | x | x | x

> object 권한은 특정 object에 권한을 부여

`GRANT ALL colum ON schema obect TO PUBLIC WITH GRANT OPTION`

EX) SCOTT에게 권한 부여

```
SQL> GRANT SELECT, INSERT(id, first_name, dept_id),
  2     UPDATE(first_name)
  3  ON s_emp TO scott;
```

EX2) OBJECT 권한 부여

```
 SQL> GRANT SELECT, INSERT
  2  ON s_dept
  3  TO scott
  4  WITH GRANT OPTION;
```
SCOTT 이 타 사용자에게 권한부여 가능

EX3) 

```
SQL> GRANT SELECT 
  2  ON s_dept
  3  TO PUBLIC;
```
SELECT 권한 전체에게 줌

- OBJECT 권한에 대한 DATA DICTIONARY 정보

`SQL> DESCRIBE user_tab_privs_made`
USER_TAB_PRIVS_MADE : 자신의 테이블에 대한 모든 Object권한을                        저장

EX)
```
SQL> SELECT *
  2  FROM user_tab_privs_made;

```

`SQL> DESCRIBE user_tab_privs_recd`
USER_TAB_PRIVS_RECD : 자신이 부여받은 모든 테이블에 대한                          Object권한 저장.

<BR>

---

<BR>

##### SYNONYM 을 통한 권한제어

> SYNONYM<BR>
- 데이터베이스내의 Object들에 대한 별명(동의어)이다. 
- 이것은 단순히 별명이기 때문에 별다른 저장 공간이 필요한 것이 아니라 데이터 사전 정보( Data Dictionary )에만 저장되어 진다.
- Synonym은 보안상, 편리함의 이유로 사용되어 진다.
- 예를 들면, 원래 Object의 이름을 숨기고자 할 때
- SQL문을 단순화 할 때 ( 사용자 별명을 사용하지 않는다 )
- 원격지 데이터의 위치정보의 투명성을 보장하기 위해 사용


`CREATE SYNONYM name FOR object`

EX)

```
 SQL> CREATE SYNONYM s_dept
  2  FOR plusdemo.s_dept;
```

EX) 오브젝트 권한 삭제

```
 SQL> REVOKE SELECT, INSERT
  2  ON s_dept
  3  FROM scott;

Revoke succeeded.
```

EX) SYNONYM 삭제

`DROP SYNONUM s_dept`

<BR>

---

<BR>

##### ROLE을 통한 권한 제어

- Role은 시스템 권한과 Object권한의 조합으로 이루어져 있어 이런 권한들의 조합을 어떤 사용자나 다른 Role에 줄 수 있다.
- Role은 특정 사용자에 의해 소유될 수는 없다.
- Role은 특정 저장공간을 사용하지 않고 데이터 사전에만 기록된다.

`CREATE ROLE role`

EX) 롤 만들기
`CREATE ROLE acct_rec`

롤에 비밀번호 부여하기
```
SQL> CREATE ROLE acct_pay
  2  IDENTIFIED BY bicentennial
```

OS에서 비밀번호 관장하기
```
SQL> CREATE ROLE manager
  2  IDENTIFIED EXTERNALLY;
```

EX) 롤에 권한 부여하기
```
 SQL> GRANT CREATE SYNONYM, CREATE TABLE
  2  TO manager;
```

EX) 롤 사용자에게 부여하기
`GRANT manager TO plusdemo`
