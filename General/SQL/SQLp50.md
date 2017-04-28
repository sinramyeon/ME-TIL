### SQL -p50

---

##### 연산자 우선순위

우선순위 | 연산자
---- | ----
1 | 비교 연산자들(=, <>, !=, IN, LIKE, IS NULL, BETWEEN, AND...)
2 | NOT
3 | AND
4 | OR

`SELECT column_name FROM table WHERE condtion AND condition`

Ex)
salary 가 1000보다 크거나 같고 dept_id가 44인것 또는 dept_id가 42인 것
```
SELECT  last_name, salary, dept_id
FROM  s_emp
WHERE  salary >= 1000  AND  dept_id = 44 OR dept_id = 42;

```

salary가 1000이상 이면서 dept_id가 44이거나 42인 데이터
```
SELECT last_name, salary, dept_id
FROM  s_emp
WHERE  salary >= 1000 AND (dept_id = 44 OR dept_id = 42 );

```

---

#### Order by

DESC 내림차순, ASC 오름차순(표시하지 않을 시 자동)

```
SELECT  last_name, salary
FROM  s_emp
WHERE  dept_id = 45
ORDER BY  salary  DESC;

```

Ex)
두 개 이상의 항목으로 order by
```
SELECT  region_id , id, name
FROM  s_dept
ORDER BY  region_id, id;
```

dept_id별로는 오름차순, salary별로는 내림차순

```
SELECT  last_name, dept_id, salary
FROM  s_emp
WHERE  title = ‘Stock Clerk’
ORDER BY  dept_id , salary  DESC;
```

---


#### Data Dictionary

##### Database에는 일반 User가 생성,관리하는 테이블 뿐만 아니라 Database가 자체적으로 생성,관리하는 테이블들이 있다. 이것들을 “Data Dictionary”라고 한다. ( View 포함 )

User가 생성한 모든 Objects에 대한 정보를 가지고 있는 Data-Dictionary View
`DESCRIBE user_object`

----

#### SQL * Plus

명령어 | 설명
---- | ----
L[ist] | 전체 sql
A[ppend] text | 현재 줄에 text 추가
C[hange]/old text/new text | 현재 줄의 old text 를 new text로 변환
DEL | 현재 줄 삭제
I[nput] | 현재 줄 다음에 한 줄 삽입
n text | N번째줄을 text로 대체

 Ex)
 
 ```
 SQL> L 1
   1 * SELECT name
SQL> A  “Region”
SQL> L
   1 * SELECT name “Region”

 ```
 
 ```
         SQL> L 1
   1 * SELECT name “Region”
SQL> C/name/id
   1 * SELECT id “Region”
```

```
SQL> L
   SELECT id “Region”
    FROM s_region
SQL> I
   3   WHERE name = ‘Asia
```

Ex2)

Run 명령 : 해당 SQL을 먼저 보여주고서 실행

```
 RUN
SELECT  id “Region”
FROM  s_region
WHERE  name = ‘Asia’;

```

/ 명령 : 메모리로 불러온 SQL문이나 메모리에 있는 명령 실행

---

#### SQL 파일로 저장하기

`SAVE filename [replace]`

Ex)현재 Default Directory에 ‘research.sql’ 파일이 생성

```
SQL> SELECT first_name, last_name, title
    2  FROM s_emp
    3  WHERE dept_id = 41
    4  
SQL> SAVE research
   Created file research

```

불러오기

`GET filename`

실행하기

`START filename`
`@filename`

----

#### 데이터 조작

- 데이터 insert
`INSERT INTO s_dept VALUES(11, 'finance', 2)`

```
SQL> INSERT INTO s_emp (id, last_name, first_name,
      2  userid,start_date,manager_id, dept_id )
      3  VALUES (26, 'Hering', 'Elizabeth', 'ehering',
      4                   SYSDATE, 2, 32 );
```
- Insert 절에 나열된 column 순서와 values 절 순서를 같게
- 만약 column을 생략할 시 describe 명령에서 나타난 테이블 컬럼 순서대로 넣으세요

- 특별한 값 넣기


Ex)
Null Keyword 사용하기
```
SQL> INSERT INTO s_customer
  2  VALUES (216, 'Sports on Wheels', NULL, NULL,  3  NULL, NULL,NULL, NULL, 'GOOD', 12, 2, NULL );

```

공백"" 사용하기
```
SQL> INSERT INTO s_customer
  2  VALUES (217, 'Wheels-R-Us', '', '', '', '', '', '', 'GOOD', 12, 2, NULL);
```

INSERT절 항목 제외해버리기
```
SQL> INSERT INTO s_customer (id, name,
    2     credit_rating, sales_rep_id, region_id)
    3     VALUES (218, 'Fitness Center', 'EXCELLENT', 12, 2);

```

Ex2)
시퀀스 확인하기
```
SQL> SELECT  object_name
  2  FROM  user_objects
          3  WHERE  object_type = 'SEQUENCE';

```


시퀀스를 사용한 고유 키 생성, 입력
```
SQL> INSERT INTO  s_dept (id, name, region_id )
  2  VALUES  (s_dept_id.NEXTVAL, 'Education', 1);

```
- Sequence_name.NEXTVAL : 현재값의 다음값을 주고 자신은    다음값을 가진다.
- Sequence_name.CURRVAL : 현재값을 준다.
- 한번 부여된 Sequence값은 되돌릴 수 없다.
- 재 정의하여 초기값을 재 설정할 수 있다


Ex3)
함수 사용하여 넣기
```
SQL> INSERT INTO s_emp (id, first_name, last_name,
    2     userid, start_date, salary)
    3     VALUES (27, '', 'Student', USER, SYSDATE, NULL);
```

- USER : 현재 유저 이름
- SYSDATE : 현재 시스템 일자, 시간