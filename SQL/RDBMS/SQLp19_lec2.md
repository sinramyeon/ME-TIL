#### SQL -p19 강의2

---
##### DDL, DML, DCL

약어 | 정의 | 예시
---- | ----- | ------
DDL | Data Definition Language | Create, Alter, Drop, Truncate, Grant, Revoke, Comment
DML | Data Manipulation Language | Insert, Update, Delete, Select, Lock table, Explain plan, call
DCL | Data Control Language | Commit, Rollback, Savepoint


---

##### varchar와 varchar2

★  |  varchar | varchar2
---- | ----- | ------
format | String 형태의 가변 길이 값 저장 | String2 형태의 가변길이값 저장
정의 | 앞으로 사용하지 않는 것이 좋다. 별도의 데이터 유형으로 지정될 예정. | 4000바이트의 최대 가변길이 문자열을 저장하는데 사용. 


###### char와 varchar2

★  |  char | varchar2
---- | ----- | ------
쓰임 | 고정길이 문자열을 저장 | 가변길이 문자열을 저장
차이 | 사용자가 고정길이 보다 짧은 값을 저장하면 나머지 공간은 Space(공백)으로 채움. 값을 너무 크게 입력하면 고정 길이에 맞게 잘림 | 적은 값을 입력 시 실제 저장공간도 적게 줄어듦
쓰는곳 | 주민번호, 성별... | 내용
검색법 | char(5)일 경우 where post = '50___' | where post='50'

----


##### Unique Key(UK)
- 테이블 내 해당 컬럼 값은 항상 유일해야 한다
- Primary Key와 유사하나 **NULL값을 중복 허용**
- 내부적으로 Unique Index를 만들어 처리한다
- 테이블 내에서 UK는 여러번 지정 가능하다
<br><br>
<br><br>
##### Primary Key(PK)
- 테이블 내 해당 컬럼 값은 항상 존재해야 하며 유일해야 한다.
- 테이블에 대한 기본 키를 생성한다.
- 기본키는 테이블 당 하나만 존재하며 반드시 하나의 컬럼으로만 구성되는 것은 아니다.
- **NULL 안됨**, 이미 테이블에 존재하는 중복 데이터 다시 입력 안됨.
- Unique Index가 자동으로 생성
- NOT NULL + UNIQUE
<br><br>
<br><br>
##### Foreign Key(FK)
- 해당 컬럼 값은 참조되는 테이블 컬럼 값 중 하나와 일치하거나 NULL을 가진다
- 두 테이블의 데이터 간 연결을 설정하고 강제 적용
- 참조하고자 하는 컬럼이 PK또는 UK가 잡혀있어야 사용 가능
- 부모 테이블을 삭제하려면 자식 테이블부터 삭제해야 함