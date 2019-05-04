문자 함수 예제

함수 | 설명 | 비고
--- | --- | ---
LOWER | 소문자로 변환 |
UPPER | 대문자로 변환 |
LENGTH | 문자 길이 |
SUBSTR | 문자 값 중 원하는 길이만큼만 나타냄 |
RTRIM | 오른쪽 공백을 잘라냄 |
LTRIM | 왼쪽 공백을 잘라냄 |
TRIM | 모든 공백을 없앰 |
REPLACE | 문자열을 대체 |
COALESCE | 조건에 따른 치환 |
INITCAP | 첫 글자는 대ㅜㅁㄴ자, 나머지 글자는 소문자로 변환 |

---

숫자 함수 예제

함수 | 설명 | 비고
--- | --- | ---
ROUND | 소숫점 자릿수 지정 반올림 |
TRUNC  | 해당 소수점 자리에서 잘라낼 때 사용 |
MOD(M, N) | M을 N으로 나눈 나머지 |
ABS | 절대값 | 
SIGN | 양수면 1, 음수면 -1, 0이면 0 반환 |
SQRT | 제곱근 |
COS | 코사인 |
SIN | 사인 |
PI | 파이값 |
TAN | 탄젠트값 |

Ex)

소숫점 둘째자리 반올림 ROUND(열 이름, 1)
소수점 첫재짜리 반올림 ROUND(열 이름, 0)
정수 첫째자리 반올림 ROUND(열 이름, -1)
---

날짜함수 예제


함수 | 설명 | 비고
--- | --- | ---
ADD_MONTHS | 지정한 날짜에 개월 수를 더한값을 출력 |
SYSDATE, NOW, GATEDATE | 현재 시스템의 날짜를 반환 |
LAST_DAY | 해당 월의 마지막 날짜 반환 |
MONTH_BETWEEN | 지정된 월 간의 월 수를 반환 |

사실 몇 가지 함수는 특정 오라클 등에서만 제공하기도 하고, 함수 종류가 다르기도 하다.
이런 경우에는 주력 DB를 두고 함수 정리 표를 비교해서 사용하면 된다.
---

COUNT 함수


함수 | 설명 | 비고
--- | --- | ---
COUNT | 행의 수 | null 값 포함한 행 수 COUNT(*), null값을 제외한 전체 행의 수 COUNT(열 이름), 중복을 제외한 행 수 COUNT(DISTINCT 열 이름)
SUM | 행의 합계 | 전체 합계 SUM(열 이름)
AVG | 행 평균 |
MAX | 행 최댓값 |
MIN | 행 최솟값 |
STDENV | 행의 표준편차 |
VARIANCE | 행의 분산 |

EX)
열 평균 계산

AVG 함수 사용시 NULL값을 가진 열이 생략되어 전체 평균값이 잘못될 수 있다.
NULL값을 포함해서 계산하고 싶으면 `AVG(COALESCE(열 이름,  0))` 을 사용한다.

---

조건문

> SELCT 열 이름 1,
> CASE WHEN 조건 1 THEN 결과 1
>      WHEN 조건 2 THEN 결과 2
>      ELSE 결과 3 END AS 새로운 열 이름
> FROM 테이블명;

오라클의 경우 DECODE 함수로도 사용한다.

1. 구문형식

Decode( expr, search1, result1, search2, result2, …… , default_result)



2. 설명

디코드 함수는 첫번째 파라미터로 들어오는 expr 표현식을 검사하여, 
이 값이 search1에 해당할 경우 result1을 반환하고, 
search2에 해당하면 result2를 반환합니다. 

※ default_result는 생략될 수 있고, 모든 search값을 만족하지 않을 경우 null을 반환하게 됩니다.

※ DECODE에 사용되는 식은 동등비교만 가능합니다. (true / false)


