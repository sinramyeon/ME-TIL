# spark I.O / Performance

---

## spark I/O

파일을 읽고 쓰는 법을 살펴보자.

`sc.textFile("sample_input").collect()` 로 한 파일을 읽어올 수 있으며, `sc.wholeTextFiles("sample_input").collect()` 로는 각 라인 별로 나뉘어진 리스트를 반환받을 수 있다. `wholeTextFiles()` 로 작은 데이터들을 나눠 분석하기에 적합하다. 

### Pickle

pickle한다는 것은 스파크에서 파일을 저장하고 로드하는 하나의 방법이다. 파이썬 pickle모듈을 생각해보면 된다.

우선 파이썬 pickle모듈부터 짚고 넘어간다.

##### pickle 모듈

파이썬에서 리스트나 클래스 자료형으로 파일을 불러오고 싶을 때 쓰는 라이브러리이다.

```
>> List = ['a', 'b', 'c']
>> with open('text.txt', 'w') as f:
       f.write(List)

Traceback (most recent call last):
File "<pyshell#5>", line 2, in <module>
f.write(LIST)
TypeError: must be str, not list
```

사용법은 아래와 같다.

> pickle.dump(데이터, 파일)
> 변수 = pickle.load(파일)

```
>> import pickle
>> List = ['a', 'b', 'c']
>> with open('test.txt','wb') as f:
       pickle.dump(LIST,f)
>> with open('test.txt','rb') as f:
       data=pickle.load(f)         
>> data
['a', 'b', 'c']
```

pickle	모듈을 쓸 때는, 파일형식은 꼭 바이트로 다룬다. (**wb, rb 등...**)
파이스파크에서는 피클 라이브러리를 아래와 같이 다룰 수 있다.

```
sc.parallelize(["a", 1, {"key" : "value"}]).saveAsPickleFile("pickleFile")
# 이렇게 피클형태로 저장 후

sc.pickleFile("pickleFile").collect()
# 같은 방식으로 읽어서 사용하면 된다.
# 만든 그대로의 형식으로 출력된다.
# ["a", 1, {"key" : "value"}]
```

그러나 pickle을 적용할 수 없는 네스티드 클래스 등에도 적용하지 않게 주의한다. 이것은 파이썬 [문서](https://docs.python.org/3/library/pickle.html)에서 확인해보는게 좋다.


### 하둡 인풋포맷

hadoop input format으로 된 자료형도 물론 지원한다. 하둡인풋포맷이 자바기반이라 조금 헷갈리게 보일 수도 있다.

```
rdd = sc.newAPIHadoopFile(
path = "test.txt",
inputFormatClass="org.apache.hadoop.mapreduce.lib.input.TextInputFormat",
valueClass="org.apache.hadoop.io.Text",
keyConverter=None, # 필요한 포맷으로 키와 밸류를 포맷할 수 있다.
valueConverter=None,
conf={}, # 맵으로 하둡conf에 필요한 내용을 입력해도 된다.
batchSize=0 # 파이썬 자바 연결용 사이즈인데 그냥 0으로 안전히 두자...
)

rdd.collect()
```

### 하둡 아웃풋포맷

파이스파크는 역시 하둡으로 출력할 수 있다. 그러나 pickle이나 text로 파일을 빼내는것보다는 복잡하다.

```
rdd.saveAsHadoopFile(
path = "output", # 어디에 저장할 것인지?
outputFormatClass = "org.apache.hadoop.mapreduce.TextOutputFormat",
keyClass = None,
valueClass = None,
keyConverter = None, # 어떻게 밸류값을 변환할건지 정해줄 수 있다.
conf = None, # map형태로 설정을 넣는다
compressionCodecClass = "org.apache.hadoop.io.compress.GzipCodec" # 압축해야 하면 압축 형식을 정해준다.
)
```

하둡은 자바기반이므로 자바 생각을 하면서 api 문서를 참고해도 괜찮다.

---

## Performance

스파크 사용 시 파이썬 특유의 걱정거리기도 한 성능 문제를 다뤄보도록 한다.

### broadcast()를 쓰자.

일반적으로 변수를 선언해서 연산에 이용할 수 있다.

```
important_data = 21
rdd = sc.parallelize(xrange(5))
rdd.map(lambda x:x + important_data).collect()
# [21, 22, 23 24, 25]

rdd.map(lambda x:x + important_data).repartition(2).map(lambda x:x + important_data).collect()
# [42, 43, 44, 45, 46]
```

위와 같은 연산이 있다고 할 때, `repartiton()` 을 쓰느라고 두가지 맵을 셔플시키고 있다. 물론 여기에서는 imortant_data값이 작은 값이니까 별 상관이 없었지만... 만약 이게 아주 큰 데이터였다면 스파크에서 매칭시키느라고 성능이 엄청 떨어지고 말 것이다.

이럴 때를 위해서 `broadcast()` 를 이용한다.

```
important_data = sc.broadcast(21)
important_data
# pyspark.broadcast.Broadcast 값

rdd.map(lambda x:x + important_data.value).collect()
# 위와 같이 .value 로 꺼내 쓴다.
```

물론 작은 데이터에서는 속도에서 큰 차이가 없겠지만.... 큰 데이터를 다루고자 할 때는 작은 습관 하나로 큰 기다림을 막을 수도 있다.

물론 모두 사용한 후, `important_data.unpersist()`로 캐시 메모리를 지워주는 것을 잊지 말자.

### accumulators

어큐뮬레이터는 메인드라이버만 보고있는 항상 내가 가져다 쓸 수 있는 변수이다. 저번에 한번 정의를 익혔으니, 이번에는 어떻게 성능 활용에 도움을 줄 수 있을 지 살펴본다.

```
evens = sc.accumulator(0)
numbers = sc.parallelize(xrange(10))

def inc_and_report(x) :
    global evens
    if (x % 2 == 0) : #짝수
        evens += 1
    return x+1
```

global evens를 선언해서 지역변수를 만드는 일을 방지했다.

```
numbers.map(inc_and_report).collect()
# [1, 2, 3, 4, 5, 6...10]

evens.value
# 5
```

어큐뮬레이터를 참고하자 몇개의 짝수가 나왔는지 얻을 수 있었다.

```
numbers.map(inc_and_report).collect()
# [1, 2, 3, 4, 5, 6...10]
# 한번 더 실행했음
evens.value
# 10
```

accumulator변수이므로 값을 기억하고 있다.
다시 계산을 돌리는 일을 방지하기 위하여 기억하고 있어야 할 변수들은 accumulator를 이용하여 기억해 둔다.

```
evens = sc.accumulator(0)
numbers.map(inc_and_report)
evens.value
# 0
```

또한 spark에서 지연 평가를 한다는 것을 기억하자. map만 적용시켜서, 스파크는 아직도 실행을 안 했다.
다시 정한 accumulator의 값인 0만을 반환한다.
이 두가지 점을 잘 기억한다면, 쓸데없는 여러번의 계산을 방지해 줄 것이다.

**accumulator**를 보면서 느꼈을수도 있는데 **빌트인으로는 숫자만 입력된다.** 그래서 다른 값을 넣으려면 custom하게 만들어야 한다. [api문서](https://spark.apache.org/docs/2.2.0/rdd-programming-guide.html#accumulators)에도 써있다. 처음 사용할 때는 지나치기 쉬운 사항이니 기억해 두자.

```
ids = sc.parallelize(xrange(20))
ids.distinct().collect()
# [0, 8, 16, 1 ... 20]

from pyspark import AccumulatorParam

class SetAccumulatorParam(p) :
    def zero(self, initialValue) :
        s = set()
        s.add(p)
        return s
    def addInPlace(self, v1, v2) :
        return v1.union(v2)


ids_seen = sc.accumulator(None, SetAccumulatorParam())

def inc_and_note(x) :
    global id_seen # accumulator 만들기
    ids_seen += x
    return x+1


ids.map(inc_and_note).collect()
```

### 기타 잡기

1. RDD 를 하나 만들고 다시쓴다?

`memory, memory_and_disk` 옵션으로 따로 보관해 둔다.

2. 무조건 groupByKey()만 안쓰면 된다?

확실히 정말 느리다. `sql, pandas` 옵션을 선택하는게 정신 건강에 이롭다.
`groupByKey()`대신에 아래를 활용한다.
	`reduceByKey()` : 타입이 같을때
	`aggregateByKey()` :  타입이 같지 않아도 사용할 수 있음

3. 파티션 하나당 2gb의 사이즈 제한이 있음을 기억한다.

4. rdd대신 dataframe에 익숙해 진다?

jvm 안에 데이터를 넣어두고 파이썬은 쿼리에만 신경쓰게 한다.
py4j등을 활용해서 자파이썬 형태로 짬뽕 시킨다.
