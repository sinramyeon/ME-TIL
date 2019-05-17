# 하위 쿼리

### FROM절의 하위 쿼리

1. 조건에 맞는 대상자 선정 후 요약할 때
```
SELECT 열 이름1, 열 이름
FROM(
SELECT * FROM 테이블명
WHERE 조건절
) AS 별칭
WHERE 조건절
```

별칭을 주지 않으면 에러 발생.

2. 테이블 조인을 할 때
```
SELECT 별칭1.열 이름1, 별칭2.열 이름2
FROM 테이블명1 AS 별칭1 LEFT OUTER JOIN
(
SELECT 열 이름 1, 열 이름2
FROM 테이블명2
WHERE 조건절
) AS 별칭2
ON 별칭1.KEY=별칭2.KEY
```

### WHERE 조건절 하위쿼리

```
SELECT 열 이름 1, 열 이름 2
FROM 테이블명1
WHERE 열 이름 IN ( SELECT 열 이름 FROM 테이블명2 WHERE 조건절)
```
