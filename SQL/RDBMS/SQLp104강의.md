#### SQL -p104 강의

---

- CEIL 올림 함수
- FLOOR 버림 함수
- ROUND 반올림 함수
- TRUNC 절삭 함수

EX) 101
`
SELECT CEIL(100.5) FROM DUAL
`
EX2) 100
`SELECT FLOOR(100.5) FROM DUAL`

EX3) 123.2
`SELECT ROUND(123.17, 1)`

EX4) 123.1
`SELECT TRUNC(123.17, 1)`

---

- NVL NULL을 다른 값으로 치환
- NVL2(A, B, C) A가 NULL이 아니면 B, NULL이면 C를 반환
- NULLIF(A, B) A와 B가 동일하면 NULL, 아닐시 A를 반환

 EX)
 `SELECT empno, NVL(mgr, 0) mgr FROM emp WHERE deptno = 10`
 
 EX2)
 `SELECT empno, NVL2(mgr, 1, 0) mgr FROM emp WHERE detpno = 10`
 
 EX3)
 `CASE WHEN expr1 = expr2 THEN NULL ELSE expr1 END`와 동일
 
 ---
 
 - LPAD(COLUMN, N, STRING) 길이가 N이 되도록 COLUMN 왼쪽을 STRING으로 채운다
 - RPAD(COLUMN, N, STRING) 길이가 N이 되도록 COLUMN 오른쪽을 STRING으로 채운다

EX)0000012345

`SELECT LPAD('12345', 10, '0')AS TEST_SEQ FROM DUAL;`

EX)1234500000

`SELECT RPAD(12345, 10, '0')AS TEST_SEQ FROM DUAL;`

---

- SUBSTR(문자열 or 컬럼, 숫자1, 숫자2)
- 문자열을 숫자1부터 숫자2까지 자른다

EX)2번째부터 10개의 문자열 읽기 > BC오라클문DEF자
`SELECT SUBSTR('ABC오라클문DEF자르기), 2, 10) AA FROM DUAL`

EX)뒤에서 두 글자 추출 > Z9
`SELECT substr('ABC권경안Z9', -2) FROM DUAL;`

---

- GROUP BY + HAVING 절
- 데이터들을 원하는 그룹으로 나눌 수 있다.

```
SELECT [DISTINCT] 칼럼명 [ALIAS명]
FROM 테이블명
[WHERE 조건식]
[GROUP BY 칼럼(Column)이나 표현식]
[HAVING 그룹조건식] ;
```

> COLUMN명이 동일해야 함(GROUP BY와 SELECT의)

EX) WHERE와 같이 쓰는 경우
```
SELECT titles.pub_id, AVG(titles.price)  
FROM titles INNER JOIN publishers  
   ON titles.pub_id = publishers.pub_id  
WHERE publishers.state = 'CA'  
GROUP BY titles.pub_id  
HAVING AVG(price) > 10  
```

---

- UNION 중복을 1로 카운트 = UNION DISTINCT
- UNION ALL 모두 카운트