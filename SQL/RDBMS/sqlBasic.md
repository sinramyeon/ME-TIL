# SQL이란?

> 데이터베이스 : 여러 사람에게 공유되어 사용될 목적응ㄹ 가지고 구조적 방식으로 관리되는 데이터의 집합

DBMS(Database Management System) 으로 괄니한다. oracle, mysql, mssql, postgresql 등이 있다.

DBMS의 특성은 아래와 같다.

1. 실시간 접근성 Realtime processing 지속적이고 비정형적 질의에 대해 실시간 처리가 가능하다
2. 계속적 변화 Continouous evolution 동적인 데이터베이스는, 삽입, 삭제, 갱신 등이 계속해서 일어난다.
3. 동시 공용 Concurrent sharing 다수의 사용자가 동시에 이용할 수 있어야 한다.
4. 내용에 의한 참조 Contents reference 레코드이ㅡ 값에 따라 참조를 가능하게 한다.

---
#### 스키마
테이블에 데이터가 저장되는 방식을 정의
스키마는 데이터에비스 안에 존재하는 자료의 구조 및 내용과 자료들의 논리적, 물리적 특성에 대한 정보를 표현하는 데이터베이스의 논리적 구조

#### 기본 키 Primary key
각 행을 고유하게 해 주는 열

---

SQL의 종류

1. 데이터 정의어 DDL
데이터가 저장된 공간이 테이블, 테이블 형식이 스키마일때 DDL은 이 구조를 정의하는 언어이다.
Create, Drop, Alter
2. 데이터 조작어 DML
데이터 저장, 수정, 삭제, 조회 언어
Insert, Delete, Update, Select
3. 데이터 제어어 DCL
사용자의 권한 제어를 위해 사용
Grant, Revoke

---



