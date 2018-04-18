# Mongo vs Redis, NOSQL 비교


**NOSQL** 에 대해 알아보고, 자주 쓰는 몽고와 레디스가 같은 키밸류 스토어지만 어떻게 다른지 비교해 보도록 하겠다. 

## NoSQL?

RDBMS는 보통 " 관계형 데이터베이스 " 라고 한다. 그렇다면 NoSQL은 " 비관계형 데이터베이스 " 일 것이다. 보통 NoSQL은 키-값이나 컬럼, 문서 형식 등의 데이터 모델을 이용한다.

RDBMS가 아닌 다른 형태로 데이터를 저장하는 기술이라고 생각하면 되겠다.

> NoSQL : Not Only SQL. **비관계형** 데이터베이스.

NoSQL 제품군이 RDBMS와의 다른 점으로는

1.  스키마가 없다. 즉, 데이터 관계와 정해진 규격 ( table - column 의 정의 )이 없다.
2.  관계정의가 없으니 Join이 불가능 하다. ( 하지만 reference와 같은 기능으로 비슷하게 구현은 가능하다. )
3.  트랜잭션을 지원하지 않는다. (주의! - 어떤 건 지원하려고 하는 추세다. 예를 들어 Mongo)
4.  분산처리 (수평적 확장)를 쉽게 제공한다.  
    ( 대부분의 NoSQL DB는 분산처리기능을 목적으로 나왔기 때문에 분산처리 기능을 자체 프레임워크에 포함하고 있다 )

등이 있는데, 트랜잭션 지원에 대해서는 [이 기사](http://it-developer.tistory.com/2417)에서 보듯 조금씩 갈린다고도 생각하면 된다. 

그럼 이 NoSQL을 대체 왜, 언제 쓰는걸까? 아주 많은 양의 데이터를 효율적으로 처리가 필요할 때, 데이터의 분산처리, 빠른 쓰기 및 데이터의 안정성이 필요할 때 사용한다. 특정 서버에 장애가 발생했을 때에도 데이터 유실이나 서비스 중지가 없는 형태의 구조이기 때문에, NoSQL을 사용하는 것이다.

트랜잭션을 지원하지 않고 (may not provide full ACID (atomicity, consistency, isolation, durability) 있는데도 분산형이고 데이터 안정성(fault tolerant) 구조라는 점은 아주 흥미로운 점이라고 할 수 있다.

그럼 이런 NoSQL 의 종류로는 뭐가 있을까? 

1. 키-밸류 스토리지 형
키 밸류형 : redis, memcached, Oracle Coherence,
열 지향 와이드 컬럼 스토어 : Cassandra, HBASE, Cloud Datastore

2. 문서 형
MongoDB, Couchbase, MarkLogic, PostgreSQL, MySQL, DynamoDB MS-DocumentDB

3. 그래프 형
Neo4j

너무 많기때문에 우선은 대표적으로 몽고와 레디스만 알아보도록 하겠다.

## MongoDB?

몽고DB는 도큐먼트 지향 데이터 베이스이다. 무슨 소리냐? 라고 하면 기본적으로 bson데이터구조(Json과 비슷한 구조라고 생각하면 편하다. 비선인지 비손인지 뭐라고 읽는지 궁금함) 로 저장한다는 뜻이다. 문서 지향 데이터베이스인 몽고에서는 문서를 기본 저장 단위로 이용하면서 내장 문서와 배열을 이용해서 복잡한 계층구조를 하나의 레코드(열)로 표현할 수 있다. 

몽고 디비의 특징 중 가장 유명한 것은 **스키마가 없다**. 언제 어디서든 필드 추가 제거는 자유롭게 할 수 있다. 고통스러운 데이터베이스 설계 과정 없이 개발을 아주 쉽게 할 수 있게 된다.  

또한 몽고는 스케일 아웃Scale out 도 가능하고 스키마가 없어 필요할 때 마다 필드를 자유자재로 변경 가능하다.

또 아주 고성능이다. RDBMS보다 몇십 몇백배가 빠르다고 하는데, 정확히 몇배인지는 몰라도 적어도 사용자가 느끼기에는 편리할만큼 빠르다.

**대신!**

조인과 트랜잭션을 지원하지 않으며(*트랜잭션에 대해서는 18년부터 변경의 조짐이 있다.*) 여러 제약조건에 대한 처리도 없다. 이런 복잡한 것이 없으니까 빠르다라고 생각해도 될 것이다.

몽고디비 명령어 몇 가지를 보고 정리하겠다.

```
  
database보기

show  dbs  
  
  
"youngsung"라는 이름의 database 선택하기(youngsung이라는 db가 생성되지 않아도 가능)  

use  youngsung  
  
  
선택한 databases의 collections들 확인하기

show  collections

  
  
"kys" collcetion에 document 생성.(선택한 db와 collection이 없다면 이 시점에 함께 생성)

db.kys.**insert**( {name:"aaa", age:**30**} )

  
  
데이터 검색하기.  
조건 없이 해당 collection의 모든 document검색.

db.kys.**find**();  
  
  

name 이 "aaa"인 document들을 검색.

db.kys.**find**( {name:"aaa"} );

  
  
기존의 RDBMS의 SQL문을 어느정도 아는 상태라면 아래의 사이트에서 SQL문 <-> mongodb 쿼리 변환을 보면 쉽게 사용가능.

http://docs.mongodb.org/manual/reference/sql-comparison/

  
  
출처: [http://dokydoky.tistory.com/432](http://dokydoky.tistory.com/432) [#dokydoky]
```


## Redis?

Redis는 빠른 오픈 소스 인 메모리 키 값 데이터 구조 스토어이다. Redis는 다양한 인 메모리 데이터 구조 집합을 제공하므로 다양한 사용자 정의 애플리케이션을 진짜 손쉽게 생성할 수 있다. 주요 Redis 사용 사례로는 캐싱, 세션 관리, pub/sub 및 순위표를 생각해보면 된다.

> Redis : Remote Dictionary Server. BSD 라이선스

여기까지 들으면 우리는 Memcahed를 떠올리지 않을 수 없다.

Memcached의 기본적인 특징은 아래와 같다.

1. 처리 속도가 빠르다.

- 당연히 데이터가 메모리에만 저장되므로 빠르다. 즉, 속도가 느린 Disk를 거치지 않는다.  

2. 데이터가 메모리에만 저장된다.

- 당연히 프로세스가 죽거나 장비가 Shutdown되면 데이터가 사라진다.  

3. 만료일을 지정하여 만료가 되면 자동으로 데이터가 사라진다.

- 이름에서도 느껴지듯이 Cache이다

4. 저장소 메모리 재사용

- 만료가 되지 않았더라도 더이상 데이터를 넣을 메모리가 없으면  LRU(Least recently used) 알고리즘에 의해 데이터가 사라진다.  

  

그래서, 보통 대형 포털들에서 Static page, 또는 검색 결과 등을 캐쉬하는데 많이 사용한다.

  
그렇다면 Memcached와 Redis의 차이는 무엇일까?

Memcached | Redis
--- | ---
처리 속도가 빠르다. 당연히 데이터가 메모리에만 저장되니 빠를 수 밖에 없다. 디스크 따위는 안 거침. | 처리 속도가 빠르다. 디스크와 메모리에 저장되는데도 멤캐시드랑 차이가 없다.
데이터가 메모리에만 저장된다. 프로세스가 죽거나 장비 끄는 즉시 데이터와는 안녕 해야 한다. | 데이터는 메모리와 디스크에 저장되서 불의의 경우에도 데이터 복구가 가능하다.
만료일을 지정해서 익스파이어시 데이터는 캐시처럼 안녕히 사라진다. | 동일
저장소 메모리를 재사용한다. 익스파이어 전에도 더이상 데이터를 넣을 메모리가 없으면  **LRU(Least recently used)** 알고리즘을 따라 데이터는 사라져 버린다. | 저장소 메모리 재사용 없음. 명시적으로만 데이터 제거 가능
문자열만 지원 | 문자열, Set, Sorted set, Hash, List등 온갖 데이터 타입을 지원한다.

[위 표의 출처](http://blog.daum.net/smufu/4)

레디스 구조를 이미지화 한 [조대협](http://bcho.tistory.com/654) 님의 블로그에서 구조를 머리속에 그려볼 수 있다.

![](http://cfile1.uf.tistory.com/image/202A37504FFBDA60262DD2)

레디스를 이용하는 장점은 아래와 같다.

1. 리스트, 배열과 같은 데이터를 처리하는데 유용하다.

- value 값으로 문자열, 리스트, Set, Sorted set, Hash 등 여러 데이터 형식을 지원.

- 따라서 다양한 방식으로 데이터를 활용할 수 있다.

- 리스트형 데이터 입력과 삭제가 MySQL에 비해서 10배정도 빠르다고 한다.


2. 여러 프로세스에서 동시에 같은 key에 대한 갱신을 요청할 경우,

- Atomic 처리로 데이터 부정합 방지 Atomic처리 함수를 제공(원자성을 잘 지킨다)  

3. 메모리를 활용하면서 영속적인 데이터 보존

- 명령어로 명시적으로 삭제, expires를 설정하지 않으면 데이터가 삭제되지 않는다.

- 스냅샷(기억장치) 기능을 제공하여 메모리의 내용을 *.rdb 파일로 저장하여 해당 시점으로 복구할 수 있다.  

4. Redis Server는 1개의 싱글 쓰레드로 수행되며, 따라서 서버 하나에 여러개의 서버를

띄우는 것이 가능하다.

- Master - Slave 형식으로 구성이 가능함

- 데이터 분실 위험을 없애주는 것이 바로 위 Master - Slave 방식이다.  
  
[개발새발하는 개발 블로그](http://codingmania.tistory.com/18)  에서 잘 정리해 주셨다.


레디스 명령어를 간단히 구경해보면 아래와 같다.

```
**- keys**  : 현재의 키값 들을 확인하는 명령어.

127.0.0.1:6379> keys *

(empty list or set)

[http://redis.io/commands/keys](http://redis.io/commands/keys)

  

**- set**  : 키/값을 저장하는 명령어.

127.0.0.1:6379> set key value

OK

127.0.0.1:6379> keys *

1) "key"

[http://redis.io/commands/set](http://redis.io/commands/set)

  

**- get**  : 키에 해당하는 값을 가져오는 명령어.

127.0.0.1:6379> get key

"value"  

[http://redis.io/commands/get](http://redis.io/commands/get)

  

**- del**  : 키와 해당하는 값을 삭제하는 명령어. 여러개의 키값을 지우는 dels 가 없다.

127.0.0.1:6379> del key

(integer) 1

127.0.0.1:6379> keys *

(empty list or set)

[http://redis.io/commands/del](http://redis.io/commands/del)

  
  
출처: [http://firstboos.tistory.com/entry/redis-간단-명령어-정리](http://firstboos.tistory.com/entry/redis-%EA%B0%84%EB%8B%A8-%EB%AA%85%EB%A0%B9%EC%96%B4-%EC%A0%95%EB%A6%AC) [散策 의 정리공간]
```

## Mongo vs Redis

그렇다면 우리는 대체 언제 레디스를 쓰고 언제 몽고를 써야 하는가? 둘 간의 유의한 차이점은 무엇인가?

#### Redis

심플의 극치를 달리는 형식으로, 키-밸류 스토어를 사용한다. 관계형 데이터베이스와는 아주 다른 형태로 NoSQL을 처음 쓴다면 헷갈릴수도 있다. 그렇지만 `플랫폼개발팀 : "김설화"` 와 같은 모를래야 모를 수 없는 키와 값이 따라오는 구조와 바보도 알 수 있는 `GET 플랫폼개발팀 > 김설화` 와 같은 직관적 명령어로 러닝 커브는 그리 높지 않을 것이다.

세컨더리 인덱싱이 가능하다.  Sorted set을 이 인덱스를 만들어서 이용할 수 있다. 이외에도 다양한 인덱싱이 가능하다.

인메모리 DB이다.

####  Mongo

JSON-like 문서형 스키마프리 데이터 구조를 갖췄기 때문에, 일정히 미리 정해진 양식 따위는 없다. 관계형 데이터베이스의 많은 기능을 구현하고 있는데, 쿼리 랭귀지와 일관성에서는 관계형 데이터베이스와 아주 유사하다.

인덱스 없이는 속도가 느려지므로 인덱스를 꼭 만들자.

디스크-메모리 맵 파일이다. 인덱스는 램안에 있다.

그렇다면 언제 Redis를 쓰고 언제 Mongo를 쓰는가? 팀 내 개발자들이 몽고를 더 많이 할 줄 알 때? 아니다. 어플리케이션이 무엇을 요구하는지에 따라 골라쓰면 된다.

예를 들어 쿼리를 많이 써야 하면 Mongo를 쓰는게 더 좋을 것이다. Redis를 쓰면 많은 개발자의 고통이 따르겠지만 Mongo는 같은 쿼리를 좀 더 쉽게 이용할 수 있을 거다. 그렇지만 속도만이 모든것이라고 하면 Redis도 좋은 선택이 될 것이다.

몽고가 좀 더 심플하고 낮은 러닝커브(이미 RBMS 를 깨나 쓴다 하는 개발자들에게만!) 를 제공한다고 한다면 레디스는 아주 유연하다고 생각할 수 있다. 이것 둘 중 하나를 고르는 일은 나같은 사람에게는 권외의 일 같지만 많은 [포스트](https://blog.panoply.io/redis-vs-mongodb)에서 이 둘을 흥미롭게 잘 비교해 주었다.
