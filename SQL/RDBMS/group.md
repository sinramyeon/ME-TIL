# 그룹화

> 열 이름으로 그룹화
> SELECT 그룹화할 열 이름 1, 열 이름 2, 집계 함수 FROM 테이블명 WHERE 조건절 GROUP BY 열 이름 1, 열 이름 2

ex)
집계 함수만 사용시
`SELECT AVG(ENG), AS ENG_SCORE, AVG(MATH) AS MATH_SCORE FROM CLASS_SCORE`
ENG_SCORE | MATH_SCORE
--- | ---
62 | 78
그룹 함수도 사용시
`SELECT GENDER, AVG(ENG) AS ENG_SCORE, AVG(MATH) AS MATH_SCORE FROM CLASS_SCORE GROUP BY GENDER`
GENDER | ENG_SCORE | MATH_SCORE
---- | ---- | ----
남자 | 48 | 73
여자 | 76 | 83

*group by 절은 where 뒤, order by 앞!*

> SELCT SEG, COUNT(*) AS CNT, SUM(CARD_FLG) AS CARD_FLG FROM PPC_201312 GROUP BY SEG ORDER BY SEG;

*그룹화될 열에 NULL값 포함시 NULL값도 그룹화됨*

# 그룹화된 데이터 필터링

>SELECT 그룹화할 열 이름, 집계 함수 FROM 테이블명 WHERE 조건절 GROUP BY 열 이름 1 HAVING 집계 함수 조건;

행에 where를 사용했다면, 그룹화된 변수 필터링에서는 having을 사용한다.

> SELECT FROM WHERE GROUP BY HAVING ORDER BY 순서로 사용한다.

ORDER BY가 항상 문장 마지막에 사용된다고 기억하면 편하다.

#### WHERE, HAVING, GROUP 순서?

WHERE 조건은 데이터가 그룹화되기 전에 필터링하고, HAVING절은 데이터가 그룹된 후 필터링한다

```
SELECT CUST_NM, SUM(SALES_AMT) AS SALES_TOT
FROM PROD_SALES
WHERE SALES_AMT >= 50000
GROUP BY CUST_NM
HAVING SUM(SALES_AMT) >= 10000;
```
