# Where 절

> Select 열 이름 from 테이블명 where 비교할 열 이름 = 숫자;
> Select 열 이름 from 테이블명 where 비교할 열 이름 = '문자';

*Order by 절은 문장의 끝에 사용한다.*

비교연산자 | 설명
----- | ----
A=B | A와 B가 같다
A<>B | A와 B가 같지 않다
A!=B | A와 B가 같지 않다
A^=B | A와 B가 같지 않다
BETWEEN A AND B | 두 값 사이에 있다
IS NULL | NULL값
IS NOT NULL | NULL값이 아님

> SELECT 열 이름, 열 이름2 산술연살자 열 이름3 AS 새로운 이름 FROM 테이블명;
> SELECT 열 이름 FROM 테이블명 WHERE (열 이름 1 산술연산자 열 이름2) 비교연산자 비교할 대상

---
COALESCE 함수 문법

`COALESCE(EXPRESSION 1, EXPRESSION 2, ... , EXPRESSION N)`

EXPRESS1이 NULL이 아니면 EXPRESSION 1을 리턴하고 아니면 2.. 순서로 리턴하는 함수.

NULL값을 0 으로 치환할 때 많이 사용

* DBMS마다 다르긴 한데 아래와 같음 

- ZEROIFNULL(열이름) 해당 열에 NULL값이 포함되면 숫자 0으로 바꾸는 함수
- NVL2(열 이름, 표현식1, 표현식2) 해당 열이 NULL이면 표현식 2로 나타냄
