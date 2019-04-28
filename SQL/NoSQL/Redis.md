# Redis?

---

### NOSQL / Redis , MongoDB, Memcached...

NoSQL이란? (Not only SQL) - "*RDBMS형태의 관계형 데이터베이스가 아닌 다른 형태의 데이터 저장 기술*"

1. 스키마가 없다(규격이 없다)
2. 관계 정의가 없으니 Join도 없다
3. 트랜잭션 미지원
4. 분산처리가 쉽다.

---

### Redis

Remote dictionary server.
key-value 로 저장하며, 인 메모리 데이터베이스이다.
보통 따른 응답이 필요할 때, 트랜잭션이 쓰고 싶을 때...(Ex - 일처리, 세션 스포어, 리얼타임 순위저장 등) 사용한다.
싱글 스레드 서버로 구현되어 있고 *클러스터 문제가 있음(이후 지원한다고 하는데...)

> *클러스터?*
> 여러 대의 컴퓨터들이 연결되어 하나의 시스템처럼 동작하는 컴퓨터들의 집합을 말한다. 클러스터의 구성 요소들은 일반적으로 고속의 근거리 통신망으로 연결된다. 서버로 사용되는 노드에는 각각의 운영 체제가 실행된다. 컴퓨터 클러스터는 저렴한 마이크로프로세서와 고속의 네트워크, 그리고 고성능 분산 컴퓨팅용 소프트웨어들의 조합 결과로 태어났다. 클러스터는 일반적으로 단일 컴퓨터보다 더 뛰어난 성능과 안정성을 제공하며, 비슷한 성능과 안정성을 제공하는 단일 컴퓨터보다 비용 면에서 훨씬 더 효율적이다.[1] 따라서 열 개 안팎의 중소 규모의 클러스터부터 수천 개로 이루어진 대형 슈퍼컴퓨터에 이르기까지 널리 사용되고 있다.

장점 : 1. 아주 빠르고, 2. 커맨드 레벨의 원자성 오퍼레이션(무슨 소리지?)을 제공하고, 3. 서버사이드 잠금을 지원한다
단점 : 1. IO/Memory overhead, 2. 소스 컴파일형 설치, 3. 데이터 셋이 다 메인 메모리 안에 들어있음

---

### Redis 사용 테스트

그냥 빨리 깔아다 써보자.(https://redis.io/topics/quickstart)
`The suggested way of installing Redis is compiling it from sources as Redis has no dependencies other than a working GCC compiler and libc`

컴파일 해서 쓰라고 한다. 나는 윈도우니까... 않 되.. 돼 ...

설치 후 `PING`을 입력해서 `PONG` 이 돌아오는 걸 보면, 아래와 같이 테스트도 해 보자.

```
redis 127.0.0.1:6379> set mykey somevalue
OK
redis 127.0.0.1:6379> get mykey
"somevalue"
```

레디스 커맨드 키는 여기(https://redis.io/commands) 에서 확인해 볼 수 있다.
