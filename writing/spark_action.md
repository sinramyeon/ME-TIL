[스파크 기초 시리즈]스파크 액션

스파크 액션이란, RDD내용을 바탕으로 데이터를 가공하지 않고 원하는 결과를 얻는 조작이다.

트랜스포메이션에선 RDD내용을 씹고 뜯고 맛보고 즐겼다면 여기에선 그러지 않는다.

자주 쓰는 액션계 함수들을 정리해 보도록 하겠다.



collect()
그간 잘 써왔으나 다시 한번 collect()를 짚고 가자, collect()를 통해서 RDD안에있느 데이터를 드라이버 프로그램 속으로 리스트 모양으로 넣을 수 있다.

members = sc.parallelize(xrange(10))
numbers.collect()
# [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
작은 데이터셋을 다룰 때는 괜찮지만, 사이즈가 너무 큰 데이터를 갖고 하다간 죽어 버리기 쉽다. 그러므로 꼭 샘플링을 하거나 다른 작업을 거친다.

huge_rdd = sc.parallelize(xrange(1992348214834829))
huge_rdd.sample(withReplacement = False, fraction = 0.00001, seed=1).collect()
# 이런식으로 사이즈를 줄여 쓰자.
count()
얼마나 많은 아이템이 rdd에 있는지 반환해 준다.

members = sc.parallelize(xrange(10))
members.count()
# 10
아래의 메서드들은 2.1.0 현재 아직 실험단계라고 한다.

큰 자료에서 대강의 개수를 구하고 싶다면(예를들어 10만3천8십3개 가 아니가 10만개 정도~로 구하고 싶다면) countApprox를 사용한다.

numbers.countApprox(timeout=200, confidence=0.5)

더욱 빨리 대강의 개수만 세고 싶다면, countApproxDistinct	를 사용한다.

numbers.countApproxDistinct(relativeSD = 0.04)

relativeSD 값을 변화를 줌으로써 결과값이 달라진다.



api문서를 참고하자.



countApprox(timeout, confidence=0.95)

Approximate version of count() that returns a potentially incomplete result within a timeout, even if not all tasks have finished.

>>> rdd = sc.parallelize(range(1000), 10)
>>> rdd.countApprox(1000, 1.0)
1000

countApproxDistinct(relativeSD=0.05)

Return approximate number of distinct elements in the RDD. The algorithm used is based on streamlib’s implementation of “HyperLogLog in Practice: Algorithmic Engineering of a State of The Art Cardinality Estimation Algorithm”, available here.

Parameters:	relativeSD – Relative accuracy. Smaller values create counters that require more space. It must be greater than 0.000017.

>>> n = sc.parallelize(range(1000)).map(str).countApproxDistinct()
>>> 900 < n < 1100
True
>>> n = sc.parallelize([i % 20 for i in range(1000)]).countApproxDistinct()
>>> 16 < n < 24
True


first()

first()이름을 보면 알겠지만 rdd에서 한 아이템만 가져다 쓰고 싶을때 이 메서드를 쓰면 된다.(처음에 있는 아이템이 나온다.)

members = sc.parallelize(xrange(10))
members.first()
# 0
빈 rdd에서 꺼내려고 들면 ValueError를 발생시키니 항상 빈 값 체크를 해 주자.

take()
take()로 rdd에서 어떤 값도 드라이버로 가져올 수 있다.

members = sc.parallelize(xrange(10))
members.take(1)
# 0
members.take(10)
# [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
# first()나 collect()와도 동일한 효과를 낼 수 있지만,
# 정확히 첫 번째 값이나 전체 값이 필요하다면 take대신 해당 함수를 사용하자.
takeSample()
rdd에서 랜덤한 샘플을 뽑아 드라이버로 넘겨준다. 타겟 샘플 사이즈를 지정해줘야 하며, 예제는 아래와 같다.

numbers = sc.parallelize(xrange(10))
numbers.takeSample(withReplacement=False, num=1, seed=1)
# seed를 줬으니 항상 같은값이 골라진다.
# withReplacement 로 같은 아이템을 여러번 고를 것인지 아닐지에 대해 지정한다.
takeOrdered()
numbers = sc.parallelize([1,2,3,9,6,7,8,5,1])
numbers.take(4)
[1,2,3,9]
numbers.takeOrdered(4)
[1,1,2,3]
추측대로의 함수이다. 그러나, 작은 데이터를 다룰 때만 빠르기 때문에 아주 큰 데이터에는 사용하지 않는다. 그럴 때는 sort()를 사용하는게 더 빠르다.

내용이 많을때 : sort()
내용이 적을때 : takeOrdered()
saveAsTextFile()
스파크에서 이용한 정보를 디스크에 저장할 때 아주 유용하다. 데이터 셋 파티션 하나당 한 파일로 얻을 수 있다. 그래서 저장 전 리파티션을 하는 것이 관례이다.

numbers = sc.parallelize(xrange(1000), numSlices=5)
numbers.saveAsTextFile("test")
이렇게 저장하면, 5가지 파일과 _SUCCESS폴더랑 함께 저장되게 된다.

특별한 포맷으로 저장하고 싶을 시 파라미터를 넘겨준다.

saveAsTextFile("test.gz", compressionCodecClass="org.apache.hadoop.io.compress.GzipCodec"

저장이 gz형태로 되는 것을 확인할 수 있다.

countByKey()
키, 밸류 rdd에서 작동하며, 키를 기준으로 개수를 센다.

pairs = sc.parallelize([("a", 1),("b", 2),("b", 3)])
# 키 a에는 하나 b에는 2개가 있다.
# 1대1 매핑이 아니다

pairs.countByKey()
# defaultdict(int, {'a':1,'b':2}) 로 나온다.
# 값 자체는 안나오는걸 유의한다
forEach()
각 rdd 구성 요소마다 작업이 필요할 때 forEach()를 사용한다.

def add_to_queue(x, queue=[]) :
    queue += [x]
    return queue


add_to_queue(3)
[3]

add_to_queue(4)
[3, 4]

numberes = sc.parallelize(xrange(100))
numberes.foreach(add_to_queue)
# 이렇게 적용 시,
add_to_queue(5)
# [3,4,5] 로 나옴.
# 왜 이런 일이 생겼을까???
######## 
# 나온 3,4,5 의 큐는 드라이버 프로그램에서 온 것이다.
# add_to_queue()함수는 executor안에서 실행되고 자기 자신만의 영역을 가진다.
# 각 executor는 add_to_queue함수 복사본을 갖고 있는 것이다.
# 그렇지만, 더하는 작업이 끝나면 무시되는게 문제다.
########
스파크는 드라이버 프로그램 뒤로 변수 변경값을 자동으로 넘겨주지 않기에 주의해야 한다. 스파크 자체는 큐를 어떻게 묶을지 모르기 때문에, 각 executor에서 받아오게 된다. 

만약 executor에서 데이터를 컴바인해오고 싶다면, reduce()나 accumulator를 사용해야 한다.

accumulator에 값을 넣을 때 foreach()를 제일 많이 쓴다.



드라이버 프로그램과 executor를 혼동하거나 하면 안 된다.

