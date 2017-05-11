### SQL -p50 강의

----

##### 인덱스에 따른 속도

대량의 데이터를 처리할 경우, **OR연산 DESC 연산**은 인덱스를 타지 않아 성능 저하에 많은 영향을 준다.

그래서 OR는 IN으로, DESC는 - 로 역수를 취해 ASC 정렬을 시키는 방식 등으로 튜닝을 한다.

###### 속도 비교

`IN(201701, 201702) > LIKE 2017% > between 201701 and 201702 > DT > 201701 and DT < 201702`

---

##### 오라클 날짜 타입

기본적으로 SYSDATE를 가져오는 쿼리

`SYSDATE`

DATE 타입을 CHAR 타입으로 변환

```
//24시간
TO_CHAR(SYSDATE, 'YYYYMMDDHH24MISS')
//12시간
TO_CHAR(SYSDATE, 'YYYYMMDDHH12MISS')
//밀리언초도 하고싶어
TO_CHAR(SYSDATE, 'YYYYMMDDHH12MISSFF4')
```

CHAR 를 DATE 타입으로 변환

`TO_DATE('20131224', 'YYYYMMDD')`

년만 가져오기

`TO_CHAR(SYSDATE, 'YYYY')`
`TO_CHAR(SYSDATE, 'YY')`

월과 분 구분하기

월
`TO_CHAR(SYSDATE, 'MM')`

분
`TO_CHAR(SYSDATE, 'MI')`

시간 가져오기
`TO_CHAR(SYSDATE, 'HH24')`

날짜에 시간을 더하고 빼기

하루 전
`SYSDATE-1`

1시간 전
`SYSDATE-(1/24)`

1분 전
`SYSDATE-(1/1440)`

1초 전
`SYSDATE-(1/86400)`

---

##### 정해진 자료형의 표시법

오라클의 경우 : CHECK
MYSQL의 경우 : ENUM

```
 create table con_dept(
              deptno NUMBER(2) primary key,
              dname varchar(10) DEFAULT '홍길동',
              loc varchar(4) check(loc in('서울', '부산')));
```

```
CREATE TABLE employee_person (
    id int unsigned not null primary key, 
    address varchar(60), 
    phone int, 
    email varchar(60), 
    birthday DATE, 
    sex ENUM('M', 'F'), 
    m_status ENUM('Y','N'), 
    s_name varchar(40), 
    children int
);
```

---

##### TRUNCATE, DELETE, DROP?

###### TRUNCATE

테이블의 모든 행을 삭제할 수 있다. 인덱스도 같이 지워진다.
**속도가 빠르고 ROLLBACK 이 불가능하다.**

`TRUNCATE TABLE emp`


###### DELETE

테이블의 행을 삭제할 수 있다.
TRIGGER가 걸려있다면 각 행이 삭제될 때 실행된다.
이전에 할당되었던 영역은 삭제되어 빈 테이블이나 CLUSTER에 그대로 남아 있다. 그러므로 롤백 가능.
**속도가 느리다!** 또, **시퀀스가 기억된다.**

`DELETE FROM emp`

###### DROP

테이블과 그 공간 객체를 기양 다 지워버린다.
**네버** 되돌릴 수 없다.

`DROP emp`

