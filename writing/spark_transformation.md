[스파크 기초 시리즈] 스파크 transformation

그간 써온 스파크 글 중 이 두 가지 게시글이 제일 유용하지 않나 싶다.

초반에 스파크 rdd를 다루는 법에는 액션 과 트랜스포메이션이 있다고 했었다.

이번엔 트랜스포메이션을 알아본다.





Transformation을 적용하려고 할 때, Action과는 달리 lazy하게 적용된다. 

이 말은 즉 transformation은 최대한 미루고 미루다 필요할 때 실행된다는 뜻이다. 

아주 큰 데이터를 다루기 때문에, 아직 필요도 없는데 이미 가져와서야 아무 의미가 없기 때문이다. 스파크의 map은 mapreduce의 map과 똑같은 구조를 갖고 있다.



numbers = sc.parallelize(xrange(10))
numbers.map(lambda x: x*10).collect()
def times_ten(x) :
    return x * 10
numbers.map(times_ten).collect()
# 아직 함수가 적용되지 않았다.
# map은 이런 기능에 쓰이는 것이 아니기 때문에!
persist()
persist()는 rdd에 있는 값을 기억하도록 하는 함수이다. 스파크는 작업이 끝나면 내용을 곧바로 지우기 때문에 계속 필요한 값이라면 persist()를 적용시킨다.



input = sc.parallelize(xrange(1000))
result = input.map(lambda x : x* 85)
result.saveAsTextFile("output.txt")
# 이렇게 한 번 액션을 실행한 후,
result.map(lambda x : "number" + x).saveAsTextFile("output2.txt")
# 다시 그 rdd를 쓰려고 했더니 사라졌다.
위와 같은 오류를 피하기 위해서는 아래와 같이 한다.



input.persist(StorageLevel.MEMORY_ONLY)
# 메모리와 디스크 중 자신이 원하는 레벨에 저장한다.
input.is_cachced
# True	가 반환된다.
input.unpersist()
# 다 사용하고 나면 꼭 다시 날려주어 필요없는 소비를 막는다.
input.ins_cached
# False가 반환된다.
filter()


가장 간단한 기능인 filter()이다. rdd 내용 중 내가 원하는 내용만 거를 수 있도록 도와준다.

numbers = sc.parallelize(xrange(10))
def is_even(x) :
    return (x % 2 ) == 0
numbers.filter(is_even).collect()
# 짝수만 골라낸다.
numbers.collect()
# 짝수만 나오지 않는다. 원래 만든 내용이 나온다. filter로 내가 원하는 값만 골라낸 후 어디엔가 다시 사용할 거라면, 꼭 값을 변수에 저장해 두도록 하자.


filter()이후 RDD가 바뀌지 않는다는걸 기억하자.

flatMap()


flatmap은 하나의 rdd를 여러 rdd 결과로 변환하고 싶을 때 유용하다. map과는 다른데, map은 하나의 rdd를 받아서 하나로 바뀌는것에 유의한다.



flatMap() 하나를 -> 여러개로 쪼개고 싶을 때 map() 하나를 -> 하나로 만들 때



flatmap을 사용할 때의 예제를 들어본다.



text = sc.textFile("text.txt")
text.collect()
words = text.flatMap(lambda x : x.split(" "))
# 띄어쓰기를 기준으로 한 여러개의 rdd로 나누었다.

words.collect()
# 로 단어별로 나뉜 상태를 확인해볼 수 있다.

text.map(lambda x : x.split(" ")).collect()
# 만약 map을 적용했다면?
# map은 리스트 안에 리스트가 있는 형태로 반환된다. (RDD 1개에 1개)

words.count()
# 나뉜 단어의 수를 셀 수 있다.
mapPartitions()


map은 rdd안에서 작동했다, rdd 파티션에서 작동하는게 mapPartition이다. 파티션이란 rdd 안 데이터의 한 chunk이다. mappartition은 데이터 전체 집합을 만드는데 쓸 수 있다. stackoverflow를 참고한다.



What's the difference between an RDD's map and mapPartitions method?



Thee method map converts each element of the source RDD into a single element of the result RDD by applying a function. mapPartitions converts each partition of the source RDD into multiple elements of the result (possibly none).

무슨소린지 싶겠지만 우리 영어 설명 대신 영어 코드를 보면 알아먹는 개발자들은 아래를 보면 이해하실 수 있을것이다.



text = sc.textFile("text.txt", minPartitions = 5)
words = text.flatMap(lambda x: x.split(" "))
def count_words(iterator) :
   counts = {}
   for w in iterator :
       if w in counts :
           counts[w] += 1
       else :
           counts[w] = 1
       yield counts

word_counts = words.mapPartitions(count_words)
word_counts.collect()
# 워드카운트 내용을 확인할 수 있다.
# 각 파티션이 합쳐졌다.
mapPartitionsWithIndex()


맵파티션 인덱스는 기본적으로 맵파티션과 아주 큰 차이가 있는 것은 아니다. 차이점이라고는 파티션 이터레이터에서 인덱스도 같이 준다는 것이다.

Sample()
샘플링을 활용하여 큰 데이터셋을 작게 만들 수 있다. 샘플은 데이터프레임의 샘플 서브셋을 반환한다.

data = sc.parallelize(xrange(10000))
data.count() # 10000
data.sample(False, 0.1).count()
# 샘플링 -> 항상 다른 숫자나 나옴


sample(withReplacement, fraction, seed=None) 로 샘플을 추출할 수 있다. 

withReplacement는 아이템을 두 번 뽑아올것인가 여부이고, 

fraction은 결과 크기를 얼만큼 비율로 보고 싶은지를 기입하면 된다. 

만약 똑같은 샘플을 다시 이용할 예정이라면, seed에 난수를 입력하여 쓴다.

Union()
RDD 두 가지를 합치는 함수이다.

rdd1 = sc.parallelize(xrange(5))
rdd2 = sc.parallelize(xranger(5,10))
rdd1.union(rdd2).collect()
# 0 1 2 3 ... 9 로 나옴!
Intersection()
교집합을 구해준다.

rdd1 = sc.parallelize([1,1,2,3,4,5])
rdd2 = sc.parallelize([1,1,,4,5,6])

rdd1.intersection(rdd2).collect()
# [1, 4] 를 반환
아주 큰 데이터에서는 성능이 조금 떨어질수도 있는 함수이다.

Distinct()
rdd = sc.parallelize(["A","B"]).cartesian(sc.parallelize(range(100)))
# 카테시안 곱으로 a와 99, b와 99으로 가능한 모든 경우를 만들었다.

first = rdd.map(lambda x: x[0])
#첫번째 요소(여기서는 a혹은 b)를 선택

first.distinct().collect()
# 이 중 겹치지 않는 중복이 아닌 것들만 골라옴
['a', 'b']
Cartesian()
RDD 두 개를 하나로 합치기 위한 함수인데, 카테시안 곱의 그 카테시안이다. 모든 멤버들끼리 매칭을 시켜주는 함수이다.(ex : 아이스크림이 5개고 쿠키가 7개라면 모든 다른 경우의 수를 구한다.)

ice = range(5)
cookie = range(7)
[(a,b) for a in ice for b in cookie]
# 이 결과와 동일한 결과를 스파크에서 내고 싶을 때 사용한다.


ice = sc.parallelize(range(5))
cookie = sc.parallelize(range(7))

combination = ice.cartesian(cookie)
combination.collect()
빅데이터 사이에서는 이 카테시안 메소드를 활용할 수 없을 수도 있다. (너무 크기 때문에) 그럴 때는 다른 메서드를 활용한다. 대표적으로는 join이 있다. 카테시안은 너무 큰 데이터에는 활용하지 않도록 주의한다. 또는 작게 나누어 활용하는것을 추천한다.

pipe()
pipe(self, command, env={}) 로 커맨드 창에 입력할 명령어를 스파크에도 적용시킨다.

numbers = sc.parallelize(xrange(11))
numbers.pipe("grep 1").collect()
# [1, 10]
# 1이 들어간거만 나옴
i/o가 가능한 모든 언어로 적용할 수 있다.

sc.parallelize(['1', '2', '', '3']).pipe('cat').collect()

['1', '2', '', '3']
Coalesce()
coalsce() 함수 이해를 위해서는, 스파크가 어떻게 자료를 저장하는지에 대해서 알아야 한다.

 스파크는 빅데이터를 다루기 위해 만들어졌기 때문에, 한번에 한곳에 다 저장하지 않는다. 

대신 내 클러스터 전체에 나눠서 파티션을 만들어서 저장한다. 

coalsce() 함수를 이용해서, 같은 executor안에 있는 파티션들을 관리해서 파티션 수를 효율적으로 줄일 수 있게 된다.

rdd = sc.parallelize(xrange(10000), numSlices=100)
# 파티션이 100개인 rdd
rdd2 = rdd.coalesce(10)
# 파티션을 줄였다.


그럼 파티션을 몇 개 까지 만들 수 있을까. 자신의 컴퓨터에 달렸다. 세 파티션 위에서 다섯개 프로세서를 가질 수 없듯이... 클러스터의 씨피유당 2-4 개 파티션을 만드는게 좋다고 api에서는 말한다. 물론, 만들다 보면 한 1000개씩 만들고 있기는 하다.

task overhead 가 이런 이유로 발생하게 된다. 그러므로 씨피유와 클러스터당 2-4개 파티션이 적당하다는걸 기억하자.

Repartition()
리파티션을 하기전에 파티션이 무엇인지부터 알아보자. 

스파크는 cluster안 executor들에 rdd를 위치시킨다. 스파크는, 어떤 데이터가 어디 노드에서 끝날지 정하기 위해 파티셔너를 쓰고 있다. 

잘 짜진 파티셔너는, 각 노드마다 같은 양의 데이터를 분배한다. 

만약 한 노드에만 엄청 많이 넣으면, 퍼포먼스 문제가 생기고 말 것이다. 이런걸 longtail performance issue라고도 부른다.

repartition() 함수로 파티션 커맨드를 다시 재정리할 수 있다. 스파크 파티션 관련 글을 참조 해볼 수 있다. (참고로 scala 기반으로 글을 써 주셨다.)



numbers = sc.parallelize(xrange(1000), numSlices = 1)
numbers.repartition(100)
# 파티션 1개로 만들어둔걸 100개 파티션짜리로 바꿨다.

# 만약 큰 데이터셋이었다면 많은 트래픽을 발생할 코드일 것이다.
# 그냥 트래픽 없이 줄이기나 하고 싶다. 하면 coalesce()를 사용하는게 낫다.
공식 문서에도 아래와 같이 적혀 있다.



repartition(numPartitions)

Return a new RDD that has exactly numPartitions partitions.

Can increase or decrease the level of parallelism in this RDD. Internally, this uses a shuffle to redistribute data. If you are decreasing the number of partitions in this RDD, consider using coalesce, which can avoid performing a shuffle.
RepartitionAndSortWithinPartitions()
메소드 대로 파티션 내에서 다시 정렬후 다시 재 파티션을 하는... 거의 쓸일이 없을것같은 메서드라고 생각한다. 리파티션 후 바로 소팅하는게 더 효율적이므로 활용한다고 한다.

pairs = sc.parallelize([[1,1], [1,2], [2,3]])
paris.repartitionAndSortWithinPartitions(2).glom().collect()
# glom()?
# 파티션 안에 있는 모든 엘리먼트를 가져다가 리스트에 넣고 그 리스트를 rdd로 돌려줌
결과는 아래와 같이 나온다.

[[(2,3)], [[1,1], [1,2]]]
보기 조금 헷갈려 보이지만 키별로 정렬되어 나왔다.
