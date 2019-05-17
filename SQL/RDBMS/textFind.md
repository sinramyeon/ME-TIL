# Like 연산자를 활용한 필터링

1. 뒤에 나오는 문자를 모를 때

SELECT 열 이름 FROM 테이블 WHERE 열 이름 LIKE '찾고싶은문자%'

2. 앞에 나오는 문자를 모를 때

SELECT 열 이름 FROM 테이블 WHERE 열 이륾 LIKE '%찾고싶은문자'

3. 앞뒤에 나오는 문자를 모를 때

SELECT 열 이름 FROM 테이블 WHERE 열 이름 LIKE '%찾고싶은문자%'

4. 뒤에 나오는 문자 한자를 모를 때
(마지막 글자가 생각나지 않을 때?)
SELECT 열이름 FROM 테이블명 WHERE 열 이름 LIKE '찾고싶은문자_'
(마지막 두글자가 생각나지 않을 때?)
SELECT 열이름 FROM 테이블명 WHERE 열 이름 LIKE '찾고싶은문자__'

5. 앞에 나오는 문자 한자를 모를 때
(맨 앞글자가 생각 안날때?)
SELECT 열이름 FROM 테이블명 WHERE 열 이름 LIKE '_찾고싶은문자'

6. 시작과 끝문자만 알 경우
SELECT 열이름 FROM 테이블명 WHERE 열 이름 LIKE '첫글자%마지막글자'

7. 특정단어 빼고 검색하고 싶을 경우

SELECT 열이름 FROM 테이블명 WHERE 열 이름 NOT LIKE '%찾기싫은문자%'

> % 숫자 0 또는 문자를 대체
> _ 한 글자를 대체

---
1. 열 이름 결합하기
여러 열을 결합해 한 열로 출력

SELECT 열이름1||열이름2 FROM 테이블명

2. 문자 삽입
열과 열을 결합할 때 문구를 추가

SELECT 열이름1||'삽입하고 싶은 문자'||열이름2 FROM 테이블명

EX)
`SELECT ID, NAME, CITY, COUNTRY, CITY||'('||COUNTRY||')' AS ADDR FROM CUSTOMERS;`

BERLIN(GERMANY)와 같이 출력하고 싶을 때 사용한다.

3. 공백 제거하기

오른공백 제거
SELECT 열이름1, RTRIM(열이름2) FROM 테이블

왼공백 제거
SELECT 열이름1, LTRIM(열이름2) FROM 테이블

공백제거
SELECT 열이름1, TRIM(열이름2) FROM 테이블

