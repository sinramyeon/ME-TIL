### SQL강의 끝!

---

##### JOIN

- inner join

```
select *
from (select syysdtae to_day, sysdate - 1 yesterday from dual) A
    inner join
    (select sysdate to_day, sysdate +1 tomorrow from dual) B
    on a.yesterday = b.tomorrow
```

inner join 없이도 가능
```
select *
from (select syysdtae to_day, sysdate - 1 yesterday from dual) A,
    (select sysdate to_day, sysdate +1 tomorrow from dual) B
    where a.to_day  = b.to_day
    and a.yesterday <> b.tomorrow
```

- left outer join


```
select *
    from (select sysdate to_day, sysdate =1 yesterday from dual) a
    left outer join
    (select sysdate to_day, sysdate +1 tomorrow from dual) b
    on a.yesterday = b.tomorrow
```

오라클일 경우, +로 표시 가능
```
select *
    from (select sysdate to_day, sysdate =1 yesterday from dual) a,
    (select sysdate to_day, sysdate +1 tomorrow from dual) b
    where a.yesterday = b.tomorrow(+)
```

---

##### DECODE

- 조건에 해당하는 값 추출
- DECODE(COLUMN, 조건1, 조건1일때 값, 조건2, 조건2일때 값 .... , 모두 아닐때 값)
- DECODE(VALUE, IF1, THEN1, IF2, THEN2...., default)
- decode에 사용되는 식은 동등비교만 가능
- default 값은 생략 가능하고, 모든 search값을 만족하지 않을 경우 null 반환

EX)
```
SELECT deptno, DECODE(deptno, 10, "Account",
                              20, "Research",
                              30, "Sales",
                              "Operations") name
        From dept;
```

---

##### CASE

- decode함수가 제공하지 못하는 비교연산 단점을 해결
- 조건 연산자 사용 가능
- CASE 대상값 when 비교값1 then 처리1, when 비교값2 then 처리 2.... else 디폴트처리

EX)
```
SELECT deptno, CASE deptno
    WHEN 10 THEN 'Accont',
    WHEN 20 THEN 'Reaserach',
    WHEN 30 THEN 'Sales',
    ELSE 'Operations'
From dept;
```
