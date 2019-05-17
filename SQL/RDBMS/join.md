# 테이블 합치기

## 열 합치기

### 내부 조인

1. from/where 절 사용
```
SELECT 테이블명1.열 이름1, 테이블명2.열 이름2
FROM 테이블명1, 테이블명2
WHERE 테이블명1.KEY = 테이블명2.KEY
```
2. from/where 절 과 별칭 사용
```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1, 테이블명2 AS 별칭2
WHERE 별칭1.KEY = 별칭2.KEY
```
3. inner join 사용
```
SELECT 테이블명1.열 이름1, 테이블명2.열 이름2
FROM 테이블명1 INNER JOIN 테이블명2
ON 테이블명1.KEY = 테이블명2.KEY
```
4. inner join 과 별칭 사용
```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1 INNER JOIN 테이블명2 AS 별칭2
ON 별칭1.KEY = 별칭2.KEY
```

inner join 은 교집합이라고 생각하면 편하다.

> 조인 조건을 지정하지 않은 채 두 테이블을 조인하면 곱집합이 된다.
> ex) select tmp1.cust_id, tmp1.coust_nm, tmp2.order_id from customers tmp1, orders tmp2;


### 외부 조인

1. LEFT OUTER JOIN

```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1 LEFT OUTER JOIN 테이블명2 AS 별칭2
ON 별칭1.KEY = 별칭2.KEY
```
왼쪽 테이블을 기준으로 조인, 왼쪽 테이블에는 존재하지만 오른쪽 테이블에는 존재하지 않는 키값의 경우 NULL로 반환

2. RIGHT OUTER JOIN
```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1 RIGHT OUTER JOIN 테이블명2 AS 별칭2
ON 별칭1.KEY = 별칭2.KEY;
```

오른쪽 테이블을 기준으로 조인, 오른쪽 테이블에는 존재하지만 왼쪽 테이블에는 존재하지 않는 키값이 존재하면 NULL로 반환

3. FULL OUTER JOIN
```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1 FULL OUTER JOIN 테이블명2 AS 별칭2
ON 별칭1.KEY = 별칭2.KEY
```

FULL OUTER JOIN 키워드는 왼쪽과 오른쪽의 모든 행이 반환됨(LEFT+RIGHT JOIN)

## 행 합치기

```
SELECT 열 이름 1, 열 이름 2 FROM 테이블명1 WHERE 조건절
UNION
SELECT 열 이름 1, 열 이름 2 FROM 테이블명2 WHERE 조건절
OREDR BY 1
```

1. UNION문으로 합쳐지는 SELECT 열의 숫자는 반드시 동일해야 한다.
2. 각 데이터 타입이 일치해야 한다.
3. 중복된 값이 있으면 한 가지만 표시된다.

```
SELECT 열 이름 1, 열 이름 2 FROM 테이블명1 WHERE 조건절
UNION ALL
SELECT 열 이름 1, 열 이름 2 FROM 테이블명2 WHERE 조건절
ORDER BY 1
```

UNION ALL은 중복을 포함한 모든 값을 나타낸다.



