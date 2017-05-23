### SQL -p104

---

##### 수식 연산 함수

function | Ex | 결과
--- | --- | ---
Round | Round(salary, -2) | 백자리에서 반올림
Trunc | Trunc(salary, -2) | 백자리에서 자름
Mod | Mod(7, 5) | 7을 5로 나눈 나머지

Ex)
```
SELECT last_name, salary, salary/22, ROUND(salary/22, 0)
FROM s_emp
WHERE dept_id = 50;
```

Ex2)
```
SELECT last_name, start_date, SYSDATE,
SYSDATE - start_date"TOTAL DAYS",
TRUNC(MOD((SYSDATE-start_date),7), 0)DAYS
FROM s_emp
```

----

##### NULL에 무엇을 연산해도 NULL입니다.

- NVL(A, B) A가 null일시 B로 대체한다.

----

##### 시간 데이터 가공

function | Example | Description
---- | ---- | -----
sysdate | sysdate | 지금 시간과 날짜
add_months | add_months(start_date, 6) | start_date 에서 6개월 이후 일자
last_day | last_day(start_date) | start_date월의 마지막 일자
next_day | next_day(start_date, 'friday') | start_date 이후의 첫 번째 금요일 날짜
months_between | months_between(sysdate, start_date) | 지금과 start_date간 월 차이

EX)
```
COLUMN ORDERED FORMAT A20
SELECT id, TO_CHAR(date_ordered, 'MM/YY') ORDERED
FROM s_ord
WHERE sales_rep_id = 11;
```

EX2)
```
INSERT INTO s_emp (id, last_name, start_date, dept_id)
VALUES(s_emp_id.NEXTVAL, 'smith',
TO_DATE('070393083000', 'MMDDYYHHMISS'), 50);
```

----

##### 문자열 데이터 가공


Function | Example | Description
--- | --- | ---
\|\| | 'a'\|\|'b' | ab : 문자열 합치기
INITCAP | INITCAP(name) | 첫 문자 대문자로 수정
UPPER | UPPER(name) | 모든 문자 대문자로 수정
LOWER | LOWER(name) | 모든 문자 소문자로 수정
SUBSTER | SUBSTER(name, 1, 4) | NAME의 1자리부터 4개의 문자를 가져온다
LENGTH | LENGTH(name) | NAME 문자열의 길이

EX)
```
SELECT first_name || ' ' || last_name "Employee Names"
FROM s_emp
WHERE dept_id = 41;
```

----

##### GROUP BY 절 사용

```
SELECT column_name FROM table_name WHERE condition GROUP BY expression
```
Group By절을 사용하면 Group By절에 나오는 컬럼 순으로 정렬

EX)
```
SQL> SELECT dept_id, COUNT(*) "Number"
FROM s_emp
WHERE dept_id = 41
GROUP BY dept_id;
```

EX2) HAVING 절
```
SQL> SELECT title, 12 * AVG(salary) "ANNUAL SALARY",
  2     COUNT(*) "Number of Employees"
  3  FROM s_emp
  4  GROUP BY title
  5  HAVING COUNT(*) > 2;
```
HAVING절 사용 시 그룹 함수 관련 조건에 맞는 결과만 가져올 수 있다.