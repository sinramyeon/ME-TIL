# spark 9 : key-value 페어 RDD

---

스파크는 키, 값 쌍을 갖고 있는 rdd로 다룰 수 있는 메서드들이 많다. rdd에서 필드값을 빼내서 그것을 키로 사용하는 등은 아주 자주 쓰이는 기법이다.

```
pairs = sc.parallelize([("a", 1), ("a", 2), ("b", 3), ("c", 4)], numSlices = 4)
pairs.groupByKey().collect()
# [
('a', <pyspark.resultiterable ...>),
...
]

[x for x in pairs.groupByKey().collect()[0][1]]
# [1, 2]
# []형태 안 () 들로 리턴되므로 인덱싱 가능!
```

### reduceByKey()

데이터셋에서 aggregate를 하고싶을 때 아주 유용한 메소드이다. 전체 데이터셋에 싱글 aggregate 를 할 때는 `reduce()`를 쓰고, 데이터셋의 서브셋에 aggregate를 하고 싶을 때 이 메서드를 사용하면 된다.

```
states = sc.parallelize (["TX", "TX", "CA", "CA", "TX", "CA"])
import operator

states.map(lambda x :(x,1)).reduceByKey(operator.add).collect()
# [("CA",3) , ("TX", 3)]
```

### aggregateByKey()

`reduceByKey()`와 유사하지만 훨씬 유연성이 있다. 그러나 복잡하게 느껴질 수도 있다.
`aggregateByKey(self, zeroValue, seqFunc, combFunc, numPartitions=None)` 세 가지를 넣어 활용할 수 있으며, groupByKey보다 훨씬 효율성이 높고 빠르다고 한다.

```
zero_value=set()
def seq_op(x, y) :
    x.add(y)
    return x

def comb_op(x, y) :
    return x.union(y)

numbers = sc.parallelize([0,0,1,2,5,4,5,5,5]).map(lambda x : ["짝수" if (x % 2 == 0) else "홀수", x])
numbers.collect()
# [
["짝수", 0],
....
]

numbers.aggregateByKey(zero_value, seq_op, comb_op).collect()
# [("짝수", {0, 2, 4,}), ....]
```

### sortByKey()

키별로 소팅을 할 수 있다.

```
paris = sc.parallelize([("B", 1), ("a", 2), ("A", 3), ("d", 4)])
pairs.sortByKey().collect()
# A3 B1 a2 d4순서

pairs.sortByKey(ascending=False).collect()
# d4 a2 B1 A3순서

pairs.sortByKey(numPartitions = 1).glom().collect()
# A3 B1 a2 d4 로 파티션이 [] 하나로 나옴

pairs.sortByKey(numPartitions = 3).glom().collect()
# [], [], [] 로 파티션이 세개로 나옴

pairs.sortByKey(ketfunc =lambda x:x.lower()).collect()
# a2 A3 B1 d4
```

### join()

두 rdd에 대해 이너 조인을 한다.
`join(other, on=None, how=None)`

파라미터

* other – 같이 조인할 데이터
* on – 조인할 컬럼 이름이나 리스트
* how – inner, outer, left_outer, right_outer, leftsemi 중 어느 조인 할건지

```
a = sc.parallelize([(1, "a"), (2, "a")])

b = sc.parallelize([(3 "b"), (4, "b")])

a.join(b).collect()
# 2, ('a','b')
```

### coGroup()

동일 키에 대해 양쪽 rdd를 그룹화 한다.

```
a = sc.parallelize( [(1,"a"),(2,"a")] )
b = sc.parallelize( [(2,"b"),(2,"c"), (3,"d")] )

a.join(b).collect()
# [(2, ('a', 'b')), (2, ('a', 'c'))]

a.coGroup(b).collect()
# resultiterable로 결과가 반환된다.

a.cogroup(b).mapValues(lambda x: [list(x[0]), list(x[1])].collect()
# 이런식으로 풀어서 표현해낼 수 있다.

```

평범히 조인을 쓰면 되지 왜 `coGroup()`을 사용할까? 같은 키를 가진 내용 전체를 한번에 보려고 할 때는 이 메서드가 더 적합하다.
