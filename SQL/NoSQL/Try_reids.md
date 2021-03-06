# 레디스 튜토리얼(try.redis.io)

---
레디스는 키-밸류 저장소입니다.

`SET server:name "fido"`
`GET server:name`
`"fido"`

---

- DEL으로 키에 관련된 밸류를 지울 수 있습니다.
- INCR 로 값 자동 증감이 가능합니다.

`SET connections 10`
`INCR connections`
`11`
`DEL connections`
`INCR connections`
`1`

---

INCR 로 값 증감 시 유저별 동시 작업 등의 문제가 생길 수 있다?

```Calling the INCR command in Redis will prevent this from happening, because it is an atomic operation. Redis provides many of these atomic operations on different types of data.```

> atomic operation?
> 기능적으로 분할할 수 없거나 분할되지 않도록 보증된 조작. 원자와 같이 분할할 수 없다는 것을 비유하여 이렇게 부른다. 원자 조작은 끼어들기가 불가능하며, 만일 중지되면 동작 개시 직전의 상태로 시스템을 복귀시킬 것을 보증하는 복구(백업과 복원) 기능이 제공된다.

프로그래밍적으로 보면,

> 필요한 부분은 멀티스레드 프로그램에서 공유 자원들에 대해 여러 스레드가 서로 동시에 액세스 하는 경쟁상태(race Condition)을 막시 위한 하나의 방법이다. 쉽게 말에 동기화를 위한 하나의 방법이다.

Redis는 이러한 작업을 위하여 아래와 같은 명령어를 사용한다.

- EXPIRE, TTL

`set resource:lock "Redis"`
`expire resource:lock 120`

위 명령어로 resource:lock은 120초 이후에 삭제된다.

`ttl resource:lock`
`(지워지기까지 남은 시간)`

- TTL 명령어로 값이 지워지기까지 남은 시간을 확인할 수 있다.
`-2` 가 출력될 시 키가 더 이상 존재하지 않음을 뜻한다.
`-1` 이 출력될 시 키가 절대 지워지지 않을 것임을 뜻한다.

`set resource:lock "Redis2`
`ttl resource:lock`
`-1`

---

- Redis 의 list
RPUSH, LPUSH, LLEN, LRANGE, LPOP, RPOP의 명령어를 지원한다.

- RPLUSH로 리스트 끝에 값 추가하기

`RPLUSH friends "Alice"`
`RPLUSH friends "John"`

- LPLUSH로 리스트 시작에 값 추가하기

`LPLUSH friends "Sam"`

- LRANGE로 리스트의 하위 집합 얻기(보통 LRANGE 리스트 시작할 인덱스, 마칠 인덱스)로 사용. 파이썬의 리스트 슬라이싱을 떠올리세요~

`LRANGE friends 0 -1` -> 전체 보기
`"Sam, Alice, John`
`LRANGE friends 0 1`
`Sam Alice`
`LRANGE friends 1 2`
`Alice John`

- LLEN = len(list)

`LLEN friends => 3`

- LPOP / RPOP : 시작/끝 값 지우기

---

- Redis의 set
- list와 비슷하지만 순서가 없고 중복이 없다(python의 set을 떠올려 보세요)

`SADD, SREM, SISMEMBER, SMEMBERS and SUNION.`

- SADD 는 값 추가
- SREM 은 값 삭제
`SADD superpowers "x-ray vision`
`SADD superpowers "spider web`
`SREM superpowers "spider web"`


- SISMEMBER는 값 존재 여부 반환(1 있어요 / 0 없음)
```
SISMEMBER superpowers "flight" => 1
SISMEMBER superpowers "reflexes" => 0
```

- SMEMBERS 는 리스트 내 값 반환


`SMEMBERS superpowers => 1) "flight", 2) "x-ray vision"`

- SUNION 는 셋을 여러 개 합해서 모든 값을 반환

```
SADD birdpowers "pecking"
SADD birdpowers "flight"
SUNION superpowers birdpowers => 1) "pecking", 2) "x-ray vision", 3) "flight"
```

---

- Sorted Sets
- 정렬이 가능한 Set이다.

```
    ZADD hackers 1940 "Alan Kay"
    ZADD hackers 1906 "Grace Hopper"
    ZADD hackers 1953 "Richard Stallman"
    ZADD hackers 1965 "Yukihiro Matsumoto"
    ZADD hackers 1916 "Claude Shannon"
    ZADD hackers 1969 "Linus Torvalds"
    ZADD hackers 1957 "Sophie Wilson"
    ZADD hackers 1912 "Alan Turing"
```

`ZRANGE hackers 2 4 -> 1) "Claude Shannon", 2) "Alan Kay", 3) "Richard Stallman`


---

- Hash
- dict겠지 뭐... string fields랑 value가 있고 객체 사용에도 아주 적격


```
    HSET user:1000 name "John Smith"
    HSET user:1000 email "john.smith@example.com"
    HSET user:1000 password "s3cret"

```

데이터 얻기 HGETALL:

`    HGETALL user:1000`

한번에 필드 지정

```
    HMSET user:1001 name "Mary Jones" password "hidden" email "mjones@example.com"
```
싱글 필드만 따로 지정 가능

```
    HGET user:1001 name => "Mary Jones"
```

hash 안 numeric 값들은 strings와 똑같이 취급되며 atomic하게 값을 증가시킬 수 있습니다.
- HINCRBY

```
  HSET user:1000 visits 10
    HINCRBY user:1000 visits 1 => 11
    HINCRBY user:1000 visits 10 => 21
    HDEL user:1000 visits
    HINCRBY user:1000 visits 1 => 1
```