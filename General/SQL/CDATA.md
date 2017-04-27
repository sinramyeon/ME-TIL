#### ibatis와 CDATA
[참조자료](http://hyeonstorage.tistory.com/279)
---

ibatis에서 가끔 부등호같은 <>& 등의 특수문자를 인식 못하는데 그럴 때 부등호를 <![CDATA[]]>로 감싸줘야 한다

```
<![CDATA[

SELECT *
   FROM DUAL
 WHERE A < B
     AND B > C

]]>
```

`SELECT * FROM DUAL WHERE A <![CDATA[ > ]]> B`

---

1. <statement> 타입

\<statement> 타입에는 SQL타입에는 SQL문이 작성되어 있으며 SQL 사용을 위해 호출되는 요소이다.

statement 요소 | 속성 | 하위요소 | 메소드 | 설명
--- | --- | --- | --- | ----
\<statement> | id, parameterClass, resultClass, parameterMap, resultMap, cacheMOdel, resultSetType, fetchSize, xmlResultName, remapReults | 모든 동적 요소 | insert, delete, update... 모든 쿼리 메소드 | 모든 타입의 SQL 문을 사용할 수 있는 "catch all" statement이다. 기본 타입이며, 모든 속성을 사용할 수 있다.
\<insert> | id, parameterClass, parameterMap | 모든 동적 요소, \<selectKey> | insert, update, delete |  기본적으로 insert를 위한 statement 타입이며, insert, update, delete SQL문만 지원. 특이사항으로 하위 요소로 \<selectKey>사용 가능
\<update> | id, ParameterClass, parameterMap | 모든 동적 요소 | insert, update, delete | update 를 위한 statement 타입이며, insert, update, delete SQL문을 지원한다.
\<delete> | id, ParameterClass, parameterMap | 모든 동적 요소 | insert, update | delete 를 위한 statement 타입이며, insert, update, delete SQL문을 지원한다. 
\<select> | id, parameterClass, resultClass, parameterMap, resultMap, cacheModel, resultSetType, fetchSize, xmlResultName, remapResults | 모든 동적 요소 | 모든 쿼리 메소드 | select 조회를 위한 statement 타입이며, \<statement>와 같이 모든 속성을 가질 수 있으며, 모든 쿼리 메소드를 가질 수 있다.
\<procedure> | id, parameterClass, resultClass, parameterMap, resultMap, cacheModel, xmlResultName, remapResults | 모든 동적 요소 | insert, update, delete, 모든 쿼리 메소드 | 프로시저를 위한 statement. 프로시저 사용을 명시적으로 나타낼 때 사용

<br><br>

2. Statement 타입 매핑 방식

```
<statement id="statementName"
[parameterClass="some.class.Name"]
[resultClass="some.class.Name"]
[parameterMap="nameOfParameterMap"]
[resultMap="nameOfResultMap"]
[cacheModel="nameOfCache"]
>

 select * form PRODUCT where PRD_ID = [? | #propertyName#]

 order by [$simpleDynamic$]
</statement>
```
