# Select 절

> SELECT 열 이름 FROM 테이블명

#### 한가지 행 선택

Select name from GD3

#### 여러 행 선택

Select id, name, age from GD3

#### 모든 행 선택

Select * from GD3

#### 정렬하기

Select id, name, age from GD3 order by id

#### 여러 열로 데이터 정렬

*우선순위로 정렬하고 싶은 열로 정렬이 가능하다.*
Select id, name, age from GD3 order by id, name

#### 오름차순 정렬

기본 정렬은 오름차순 정렬이다. 혹은 Order by 열 이름 ASC 로 표현한다.

#### 내림차순 정렬

Select * from GD3 order by name DESC

#### 중복 제거하기

Select distinct age from GD3

#### 별칭 사용하기

*id컬럼의 이름이 GD3_ID 로 바뀌어 선택된다.*
Select id as GD3_ID from GD3
Select id "GD3_ID" from GD3



