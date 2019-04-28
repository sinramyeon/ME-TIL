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

** COALESCE ˌkəʊəˈles**  (더 큰 덩어리로) 합치다

`COALESCE(EXPRESSION 1, EXPRESSION 2, ... , EXPRESSION N)`

EXPRESS1이 NULL이 아니면 EXPRESSION 1을 리턴하고 아니면 2.. 순서로 리턴하는 함수.

NULL값을 0 으로 치환할 때 많이 사용

* DBMS마다 다르긴 한데 아래와 같음 

- ZEROIFNULL(열이름) 해당 열에 NULL값이 포함되면 숫자 0으로 바꾸는 함수
- NVL2(열 이름, 표현식1, 표현식2) 해당 열이 NULL이면 표현식 2로 나타냄

---


### 논리연산자

AND 연산자
교집합
> SELECT 열 이름1, 열 이름2 FROM 테이블명 WHERE 조건절 AND 조건절

OR 연산자
합집합
> SELECT 열 이름1, 열 이름 2 FROM 테이블명 WHERE 조건절1 OR 조건절2

*우선순위는 OR연산자보다 AND연산자가 더 앞선다*

IN 연산자
> SELECT 열 이름 FROM 테이블명 WHERE 비교할 열 이름 IN (조건1, 조건2);
IN연산자는 여러 조건 중 하나만 만족해도 출력

NOT IN
> SELECT 열 이름 FROM 테이블명 WHERE 비교할 열 이름 NOT IN (조건, 조건2)

*IN이 OR보다 빠르다.*
*IN 연산자 안에 다른 SELECT문장을 사용할 수 있다*

