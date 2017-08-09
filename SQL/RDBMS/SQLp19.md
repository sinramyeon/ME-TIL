#### SQL -p19

---

##### SQL : database와 통신할 수 있는 유일한 언어.

##### 데이터 가져오기

테이블 구조 알아보기

` DESC 테이블이름`

테이블 데이터 가져오기

```
SELECT column_information
FROM table
WHERE condition
ORDER BY expression or keyword\
```

Ex)

> 테이블의 모든 컬럼 데이타 가져오기
`SELECT * FROM s_region`

Ex2)

> 테이블의 특정 컬럼 데이타 가져오기
`SELECT colum_name, colum_name... FROM table`


Ex3)

> 중복된 Row를 제거하고 유일한 Row 가져오기
`SELECT DISTINCT colum_name, colum_name... FROM table`

> 유일한 부서 이름 가지고 오기
`SELECT DISTINCT name FROM s_dept`


Ex4)

> Colum 이름을 원하는 이름으로 별명(Alias)사용
`SELECT DISTINCT name "Different Departments" FROM s_dept`
`SELECT last_name Employees FROM s_emp`

Ex5)

> 주어진 조건에 맞는 데이터만 가져오기
`SELECT colum_name, colum_name... FROM table WHERE condition`
`WHERE dept_id=14`

> Equal 조건 사용시 대소문자를 구별합니다.
`SELECT first_name, last_name FROM s_emp WHERE last_name="KIM"`

> Equal은 DATE Type 도 비교합니다.

- Date Type 상수를 사용할때는 항상 Single Quotes ''
- 대소문자 구분 없음
- Date Type 상수 형태는 항상 NLS_DATE_FORMAT 형태(DD-MON-YY 등)형태로

`SELECT userid, start_date FROM s_emp WHERE start_date = '04-MAR-17'`

Ex6)

> 비교 연산자 <>

`SELECT last_name, title, salary FROM s_emp WHERE title <> 'Stock Clerk'`
*<>, != 같지 않음*

> 비교 연산자 BETWEEN

`SELECT first_name, last_name, start_date FROM s_emp WHERE start_date BETWEEN '09-MAY-91' AND '17-JUN-91'`

> 비교 연산자 NOT BETWEEN (위의 반대겠죠?)

> 비교 연산자 IN

`SELECT last_name, title, dept_id FROm s_emp WHERE title IN('Warehous Manager', 'Stock Clerk')`

`SELECT id, name, region_id FROM s_dept WHERE region_id NOT IN(4,5)`


Ex7)

> Wildcard 문자 사용

`SELECT last_name FROM s_emp WHERE last_name LIKE "M%"`

`SELECT last_name, start_date FROM s_emp WHERE start_date LIKE "%91"`

`SELECT last_name, userid, FROM s_emp WHERE userid LIKE "%a%"`

`SELECT last_name FROM s_emp WHERE last_name LIKE 'S____h'`

Ex8)

> NULL 비교
> NULL != 공백, 0

`SELECT id, name, credit_rating FROm s_customer WHERE sales_rep_id IS NULL`

`IS NOT NULL 이 반대겠지?`

