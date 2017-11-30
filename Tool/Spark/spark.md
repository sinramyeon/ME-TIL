# spark 란

---

Spark를 집중적으로 다루게 될 듯 하여 Spark에 대한 정리를 해 보려고 한다.

참고할 도서들은 아래와 같다.

[러닝 스파크](http://shopping.interpark.com/product/productInfo.do?prdNo=3785953506&dispNo=008001083&smid1=common_prd)

[아파치 스파크 입문](http://shopping.interpark.com/product/productInfo.do?prdNo=5109587977&dispNo=008001083&smid1=common_prd)

[pyspark 배우기](http://shopping.interpark.com/product/productInfo.do?prdNo=5220183143&dispNo=008001083&smid1=common_prd)

---

## 1. spark가 그런데 무엇인가.

> 빅데이터 처리를 위한 클러스터용 병렬분산처리 플랫폼이다. 오픈소스다.

이런 류 프로그램으로는 맵리듀스, 하둡을 흔히 떠올리게 되는데 스파크가 좌우지간 뛰어나다고 한다. 책에는 그리 서술되어 있다. 아직 체감은 못해봤다.
[클러스터](https://github.com/hero0926/HERO_TIL/blob/master/General/Cluster.mdown) 와 [병렬분산처리](), [오픈소스](https://github.com/hero0926/HERO_TIL/blob/master/General/license.md) 에 대한 설명은 각각의 링크를 참고하자.


## 2. spark는 무슨 특징이 있을까.

> - 반복 처리와 연속으로 이루어지는 변환 처리의 고속화<br>
> - 시행 착오에 적합<br>
> - 서로 다른 처리를 통합해서 이용할 수 있는 환경

### 1. 반복 처리와 연속으로 이루어지는 변환 처리의 고속화

스파크는 데이터셋이 반복 & 변환되는 처리를 빨리 도와주려고 만들어졌다. 전에는 아주 많은 데이터를 쓰려고 [맵리듀스](https://github.com/hero0926/HERO_TIL/blob/master/General/hadoop_spark.md)를 썼다. 흩어져 있는 데이터를 키, 밸류로 묶는 맵과 맵들중 중복 제거 후 원하는것만 뽑아쓰는 리듀스를 이용했다. 그렇지만 특정 데이터의 부분집합을 여러 번 처리하는데는 조금 효율적이지 못했다.

맵리듀스는 연속처리를 할때마다 디스크에 모든 데이터를 출력하지만, 스파크는 연속처리시 불필요한 I/O를 발생하지 않도록 한다.

### 2. 시행 착오에 적합

한 대의 서버 용량보다 큰 데이터셋을 다룰 때는 여러대의 병렬분산처리 머신이 필요하다. 스파크에는 병렬 분산 환경에 관계없이 처리를 할 줄 아는 RDD가 있다. 또, CLI 스파크 쉘이 있어서 편하게 시험해볼 수 있다.

### 3. 서로 다른 처리를 통합해서 이용할 수 있는 환경

배치 처리, 스트림 처리, SQL처리, 머신러닝, 그래프처리를 다 할줄 안다. 혼합해서 쓸 줄도 알고 배치와 스트림 처리를 하나의 환경에서 다룰 수도 있다.


## 3. 스파크의 구성?

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmJuKf2hKruVBF1Bdr_zmLQvRbRvw2TomdkaoZYR80szcSERHE6w)

스파크는 여러 컴포넌트로 구성되어 있다.

- 스파크 코어

작업 스케쥴링, 메모리 관리, 장애 복구, 저장 장치와 연동 등 기본적 기능

- 스파크 SQL

sql, hive table, parquet, json 등 다양한 데이터소스 지원

- 스파크 스트리밍

실시간 데이터 스트림용 컴포넌트

- MLLib

머신러닝 라이브러리

- 그래프X

그래프 그리는 용도

- 클러스터 매니저

하둡의 Yarn, 아파치 mesos 등 다양한 클러스터 매니저 위에서 동작


## pyspark 예제

spark의 hello world인 단어세기 프로그램을 만들어 보자.

```
from pyspark import SparkConf, SparkContext

conf = SparkConf().setMaster("local").setAppName("My App")
sc = SparkContext(conf = conf) # sparkcontext 초기화하기

lines = sc.textfile("파일 위치")
nonempty_lines = lines.filter(lambda x : len(x) > 0) # 공백 거르기

words = nonempty_lines.flatMap(lambda x:x.split('')) # 단어 거르기
wordcount = words.map(lambda x:(x, 1)).reduceByKey(lambda x,y : x+y).map(lambda x : (x[1], x[0]).sortByKey(False))

for word in wordcount.collect() :
    print(word) # 단어 별 몇번 나왔나 세기


wordcount.saveAsTextFile("저장위치") # 결과 저장하기
```
