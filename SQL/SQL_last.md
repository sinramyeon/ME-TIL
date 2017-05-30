### SQL 끝!
---

##### JOIN

[참조문서](https://www.w3schools.com/sql/sql_join.asp)

inner join
![inner join](https://www.w3schools.com/sql/img_innerjoin.gif)
두 테이블간 join조건을 만족하는 row만 리턴
```
SELECT column_name(s)
FROM table1
INNER JOIN table2 ON table1.column_name = table2.column_name;
SELECT column_name(s)
FROM table1
INNER JOIN table2 ON table1.column_name = table2.column_name;
```


left join
![left join](https://www.w3schools.com/sql/img_leftjoin.gif)
조인문 왼쪽 테이블 결과 모두 가져오고 오른쪽 테이블과 매칭
매칭되는 결과가 없을시 NULL을 매칭
```
SELECT column_name(s)
FROM table1
LEFT JOIN table2 ON table1.column_name = table2.column_name;
```


right join
![right join](https://www.w3schools.com/sql/img_rightjoin.gif)
조인문 오른쪽 테이블의 결과를 모두 가져오고 왼쪽 테이블과 매칭
```
SELECT column_name(s)
FROM table1
RIGHT JOIN table2 ON table1.column_name = table2.column_name;
```



full outer join
![full outer join](https://www.w3schools.com/sql/img_fulljoin.gif)
합집합. 왼쪽이고 오른쪽이고 매칭이 있는 모든 레코드 리턴
```
SELECT column_name(s)
FROM table1
FULL OUTER JOIN table2 ON table1.column_name = table2.column_name;
```


##### JOIN은 간단히 말하자면 어떤 테이블을 기준으로 다른 테이블에 있는 Row를 찾아오는 것

---

##### OUTER JOIN

어느 한쪽 집합이 기준이 되어 다른 쪽 집합의 연결되는 조건에 상관 없이, **기준이 되는 집합은 무조건 추출**

```
SQL> SELECT s_customer.id "Customer ID",
  2     s_customer.name "Customer Name",
  3     s_ord.id "Order ID"
  4  FROM s_customer, s_ord
  5* WHERE s_customer.id = s_ord.customer_id(+)
```

s_customer 를 기준으로 s_customer.id 가 s_ord.customer_id와 같은 조건을 찾는다.

결과
```
Customer ID    Customer Name                                Order ID
--------------------- --------------------------------------------- -----------------
      201 	Unisports                                                 97
      202 	Simms Atheletics                                    98
     …………….
      207 	Sweet Rock Sports
     …………….   
```

207같은 경우의 s_ord테이블과 상관없이 추출되는 데이터도 생긴다.

---

##### SELF JOIN

From절에 동일한 테이블이 두 번 이상 나와 자신과 자신에 대해 조인

```
SQL> SELECT worker.last_name || ' work for ' || manager.last_name
  2  FROM s_emp worker, s_emp manager
  3* WHERE worker.manager_id = manager.id
```

---

##### JOIN VIEW

두 개 이상의 테이블을 join하여 뷰를 만들 수 있다.

```
SQL> CREATE VIEW empdeptvu 
  2  AS SELECT s_emp.last_name,
  3            s_emp.dept_id,
  4            s_dept.name
  5      FROM s_emp, s_dept
  6      WHERE s_emp.dept_id = s_dept.id;

View created.

```

---

##### 서브쿼리

[참고](http://wiki.gurubee.net/pages/viewpage.action?pageId=26744146)

- 하나의 SQL문안에 포함되어 있는 또 다른 SQL문
- 주 쿼리에 종속된 부 쿼리, 종속성을 갖는 쿼리
- 조건절에서 상수값이나 원래의 값으로 비교할 수 없다.
- 주 쿼리의 속성을 그대로 이어 받아 주 쿼리 내용 참조가 가능하다.

Ex)
```
SQL> SELECT last_name, title
  2  FROM s_emp
  3  WHERE title =
  4          ( SELECT title
  5            FROM s_emp
  6            WHERE last_name = 'Smith' );

```

사원 테이블에서 last_name이 Smith인 사람의 title과 같은 title을 가진 모든 사람 추출

Ex2)
```
SQL> SELECT last_name, title, salary
  2  FROM s_emp
  3  WHERE salary < ( SELECT AVG(salary)
  4                   FROM s_emp );
```

평균 월급보다 작은 급여를 가진 데이터 추출

Ex3)
```
SQL> SELECT id, last_name, salary, dept_id
  2  FROM s_emp
  3  WHERE dept_id IN
  4      ( SELECT id
  5         FROM s_dept
  6        WHERE region_id = 2 );
```

서브쿼리의 결과가 2개 이상인 경우에는 in을 사용