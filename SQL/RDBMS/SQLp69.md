### SQL -p69

---

##### 테이블 생성하기

`Create table table_name ...`

데이터타입 | 설명
--- | ---
CHAR(s) | s 길이의 고정길이 문자 타입(s:1~255)
VARCHAR2(s) | s길이의 가변길이 문자 타입(s:1~2000)
DATE | 날짜, 시간 타입
LONG | VARCHAR2와 비슷
NUMBER(p,s) | 숫자 타입
RAW | binary large object data type(그림, 음악 등에 사용). 2GB까지 가능

##### 테이블 조작하기

데이터 무결성 Constraints

Constraint | 설명
--- | ---
NOT NULL | 해당 컬럼에 NULL 허용하지 않음(DEFAULT : NULL 허용)
UNIQUE | 해당 컬럼 또는 컬럼들의 조합이 테이블 내에서 유일
PRIMARY KEY | 고유키로 정의
REFERENCES TABLE(COLUNM_NAME) | 다른 테이블과의 참조관계 설정
CHECK | 주어진 조건에 맞는 데이터만 입력 가능

```
SQL> CREATE TABLE s_emp (
   Id  		VARCHAR2(7),
   last_name  		VARCHAR2(25)
       CONSTRAINT s_emp_last_name_nn NOT NULL,
   first_name  		VARCHAR2(25),
   userid      		VARCHAR2(8),
   start_date  		DATE,
   comments   		VARCHAR2(255),
   manager_id 		NUMBER(7),
   title		VARCHAR2(25),
   dept_id		NUMBER(7),
   salary		NUMBER(11,2),
   commission_pct         NUMBER(4,2)
     CONSTRAINT s_emp_id_pk PRIMARY KEY (id),
     CONSTRAINT s_emp_userid_uk UNIQUE (userid),
     CONSTRAINT s_emp_commission_pct_ck
   CHECK (commission_pct IN (10, 12.5, 15, 17.5,
                                    20)));

```

---

##### 테이블 구조 변경하기

`ALTER TABLE table_name`

`ALTER TABLE s_region ADD (comments VARCHAR2(255));`

```
ALTER TABLE table_name
MODIFY  (column_name datatype
          [, column_name datatype] … )
```

```
ALTER TABLE s_emp
ADD CONSTRAINT s_emp_manager_id_fk
FOREIGN KEY (manager_id) REFERENCES s_emp(id);

Table altered.
```

##### 테이블 제약사항 확인

> USER_CONSTRANTS
> 해당 User의 모든 Object들에 대한 모든 제약조건(Constraint)가 있는 데이터 사전 뷰( Data Dictionary View )이다.


`DESCRIBE user_constraints`

```
SQL> SELECT constraint_name, constraint_type,
  2      search_condition, r_constraint_name
  3  FROM user_constraints
  4* WHERE table_name = 'S_EMP'

```

제약조건 종류

> C : CHECK Constraint<br>

> P : PRIMARY KEY Constraint<br>

> U : UNIQUE KEY Constraint<br>

> R : FOREIGN KEY Constraint<br>

<br>

> USER_CONS_COLUMNS<br>

> 제약조건(Constraint)에 참가한 컬럼을 저장하고 있는 데이터 사전 뷰(Data Dictionary View)이다.
>
<br>
 
`DESCRIBE user_cons_columns`
```
SQL> SELECT constraint_name, column_name
  2  FROM user_cons_columns
  3  WHERE table_name = 'S_EMP';
```

---

##### VIEW 사용

`CREATE VIEW view_name AS query WITH CHECK OPTION CONSTRAINT constraint`

뷰 컬럼 이름 추가
```
SQL> CREATE VIEW empvu41 (id_number, employee, job)
  2  AS SELECT id, last_name, title
  3  FROM s_emp
  4  WHERE dept_id = 41;

View created.
```

```
SQL> CREATE VIEW salvu41
      2  AS SELECT id, first_name FIRST, last_name LAST, 
      3      salary "MONTHLY_SALARY”
      4  FROM s_emp
      5  WHERE dept_id = 41;

View created.
```

뷰 with check option
- 뷰를 통해서 UPDATE, DELETE 수행 시, 더 이상 뷰를 통해서 볼 수 없는 데이터에 대해서는 UPDATE, DELETE를 수행하지 못하게 한다.

EX)

```
SQL> CREATE VIEW empvu20
  2  AS SELECT *
  3  FROM s_emp
  4  WHERE dept_id = 20
  5  WITH CHECK OPTION;

View created.


SQL> UPDATE empvu20
  2  SET dept_id = 42
  3  WHERE id = 16;

0 rows updated.
```

CHECK OPTION 으로 dept_id가 20인것만 볼 수 있게 해놨으므로 수정이 안된다.

---

#### 시퀀스

`CREATE SEQUENCE sequence_name INCREMENT BY n START WITH n MAXVALUE n`


```
SQL> CREATE SEQUENCE s_dept_id
  2  MAXVALUE 9999999
  3  INCREMENT BY 1
  4  START WITH 51
  5  NOCACHE
  6  NOORDER
  7  NOCYCLE;

Sequence created.
```

* ORDER/NOORDER : 병령서버에서 의미 있음 *

시퀀스 사용 예제
```
SQL> INSERT INTO s_dept
  2  VALUES (s_dept_id.NEXTVAL, 'Finance', 3);

1 row created.
```

- .NEXTAL 현재값 다음 값을 되돌려주고 자신은 다음 값으로
- .CURRVAL 시퀀스 현재값을 되돌려 준다

---

#### 인덱스(INDEX)

어떤 데이터가 어디에 있다는 위치정보를 가진 주소록 

인덱스 생성 원리
> 전체 테이블 스캔 -> PGA 내 SORT AREA 정렬 -> BLOCK 기록 -> 정렬 작업하기 때문에 **시간이 많이 걸리고** 데이터가 정렬되어 들어감

> 자동인덱스 : 프라이머리 키 또는 UNIQUE 제한규칙에 의해 자동적으로 생성되는 인덱스
> 수동인덱스 : CREATE INDEX 명령을 실행하서 만드는 인덱스

인덱스를 생성하는 것이 좋은 칼럼

1. WHERE절이나 JOIN조건 안에서 자주 사용되는 칼럼
2. NULL값이 많이 포함되어 있는 컬럼
3. WHERE절이나 JOIN조건에서 자주 사용되는 두 개 이상의 컬럼들

<BR><BR><BR>

- INDEX는 하나 혹은 두개 이상의 컬럼과 ROWID로 구성.
- INDEX는 INDEX를 구성하는 컬럼값으로 정렬되어 있으며 모든값이 동일할 경우 ROWID로 정렬되어 있다.
- ROWID는 INDEX를 읽고서 테이블을 읽기 위한 몇가지 정보로 구성되어 있다.
- INDEX는 테이블과는 물리적으로 다른 저장장소에 저장된다.
- 특정 테이블에 INDEX가 존재한다고 해서 항상 INDEX를 사용해서 테이블을 엑세스하는 것은 아니다. Database(정확히 말하면, 옵티마이져(Optimizer))가 SQL의 구성상태 및 INDEX의 구성상태등의 여러가지 조건을 따져서 Database가 사용여부를 결정을 한다.
- 대부분의 RDBMS에서 체택하고 있는 INDEX는 B-Tree(Balanced Tree)아키텍쳐이다.

테이블 액세스 방법
- By ROWID : 테이블 특정 ROW 에 대한 위치 정보(ROWID)를 가지고 직접 테이블에 액세스
- FULL-TABLE-SCAN : 테이블에 존재하는 모든 ROW에 대해 순차적으로 액세스
- BY INDEX : 먼저 INDEX를 읽고서 INDEX 의 ROWID를 가지고 테이블을 RANDOM 액세스

EX ) ROWID로 액세스
```
SQL> SELECT last_name
  2  FROM s_emp
  3  WHERE rowid = '000001B6.001A.0002';
```
EX2)FULL-TABLE-SCAN
```
SQL> SELECT last_name
FROM s_emp;
```
조건절이 전혀 없기 떄문에 Full Table Scan

EX3)BY INDEX
```
SQL> SELECT last_name
FROM s_emp
WHERE last_name = ‘Smith’;
```

인덱스 생성
`CREATE INDEX index_name ON table_name ( column_name [, column_name] … )
`

EX)
```
SQL> CREATE INDEX s_emp_last_name_i
  2  ON s_emp(last_name);
Index created.
```

유니크 인덱스 생성
UNIQUE 인덱스는 중복값들을 포함하지 않고 사용할 수 있다.
`CREATE UNIQUE INDEX index_name ON table_name ( column_name [, column_name] … )
`

EX)
```
SQL> CREATE UNIQUE INDEX s_customer_phone_uk
  2  ON s_customer(phone);
Index created.
```
고객테이블 내 전화번호를 중복되지 않게 함.

인덱스 삭제
데이터에는 아무런 영향을 미치지 않음.
인덱스는 변경ALTER 불가
`SQL> DROP INDEX emp_empno_ename_indx; `


USER_INDEXES
- 해당 사용자가 만든 모든 INDEX 에 대한 정보를 가진 데이터 사전 정보

`DESCRIBE user_indexes;`
```
SQL> SELECT index_name, uniqueness
  2  FROM user_indexes
  3  WHERE table_name = 'S_EMP';

```

---

