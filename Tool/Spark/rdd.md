# RDD란

---

스파크는 RDD라는 자료구조를 이용해서 데이터를 처리한다. 그렇다면 이 RDD라는것은 무엇인가?

> 분산된 변경 불가능한 객체 모임

스파크에서는 모든 작업을 RDD와 함께 한다. 내부적으로 스파크는 자동으로 RDD에 있는 데이터들을 클러스터에 분배하며 클러스터 위에서 수행하는 연산들을 병렬화한다. 

RDD 내부는 파티션이라는 단위로 나뉘어지는데, 스파크에서 이 파티션이 분산 처리의 단위이다. RDD의 변환(transformation)과 액션(action)을 통해 원하는 결과를 얻을 수 있다.

* RDD의 특징 *

- 불변(내부 요소 값 변경 불가)
- [지연 평가](http://dsiceeast.tistory.com/m/69)


---

## RDD 다루기

그렇다면 RDD의 변환(transformation)과 액션(action)이란 무엇일까.

> 트랜스포메이션 : 존재하는 RDD에서 새 RDD를 만들어 냄
<br>
> 액션: RDD를 기초로 결과값을 계산해서 그 값을 갖고 저장하거나 리턴함

### 변환

- filter(요소 필터링)
- map(각 요소에 동일 처리)
- flatmap(각 요소에 동일처리 후 여러 요소를 생성)
- zip(파티션 수가 같고, 파티션에 있는 요소의 수도 같은 두 RDD를 조합. 한쪽은 키, 다른쪽은 쌍으로 만듦)
- shuffle(서로 다른 파티션에 있는 같은 키를 가지는 요소의 자리를 바꿈)
- reduceByKey(같은 키를 가지는 요소를 aggreagtion - 즉 값을 합침)
- join(같은 키를 가지는 요소끼리 join)

### 액션

- saveAsTextFile(결과를 파일로 저장)
- count(요소 수 세기)

---

## RDD 처리

스파크는 보통 여러 머신으로 구성된 클러스터 환경에서 돌아간다. 그래서 클러스터 관리 시스템이 필요한데, 스파크는 세 가지를 지원한다.

- YARN(하둡거)
- Mesos
- Spark standalone(스파크 가 갖고있음. 별도로 처리 없이 이용 가능)

---

## RDD 예제

### RDD 만들기

`sc.parallelize(자료)`

데이터세트를 parallelize()에 넘겨주는 게 제일 쉽다.
아니면 그저 가져오면 된다.

`sc.textFile("위치")`

### RDD 연산


```
import os
from pyspark import SparkConf, SparkContext

# OS에서 내 스파크 홈을 ㅂ                      
if 'SPARK_HOME' not in os.environ:
    os.environ['SPARK_HOME'] = "내 스파크 위치"

conf = SparkConf().setAppName('앱이름').setMaster('local')
sc = SparkContext(conf=conf)

if __name__ == '__main__':
    ls = range(100)
    ls_rdd = sc.parallelize(ls, numSlices=1000)
    ls_out = ls_rdd.map(lambda x: x+1).collect()
    print(ls_out)

    local_out = ls_rdd.filter(lambda x : "local" in x)
    print(local_out)
```

