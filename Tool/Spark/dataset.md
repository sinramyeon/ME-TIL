# spark 공부 시리즈 11 : dataframe, dataset

대략적으로 rdd가 파이썬에서 스칼라보다 2배나 느리다. 데이터 프레임도 느린것 같지만 하여튼지 rdd 쿼리는 jvm과 파이썬 사이 오버헤드 때문에 무진장 느려진다.

데이터프레임을 사용하여서 rdd에서의 불편을 해소해볼 수 있다.


## 데이터프레임[https://github.com/hero0926/HERO_TIL/blob/master/Tool/Spark/dataframe.md]

row를 가지고 있고 스키마를 갖고 있다. sql쿼리로 굴릴 수 있으며, json, hive, parquet를 지원한다. jdbc/odbc와도 연동이 가능하다.


### 데이터프레임 생성


1. SparkSession을 이용하여 데이터를 임포트하는 방식으로 보통 데이터프레임을 생성한다.
2. rdd에서 데이터프레임으로 변경한다.

이 중 rdd에서 데이터프레임으로 변경하는 법을 보겠다. 스파크 rdd는 lazy하다. 무언가의 연산이 일어나기 전까지는 몇천줄을 써도 흥청망청 놀고 있다. 그러므로 rdd를 데이터프레임으로 바꾸는 연동을 하여 사용한다면 속도 증가를 확실히 느낄 수 있을것이다.

- 리플렉션을 이용한 스키마 추측
데이터프레임 스키마는 자동으로 정의되는데, 스키마 정의를 printSchema()로 확인해볼 수 있다.


- 스키마 명시

```
from pyspark.sql.types import *

stringDSVRDD = sc.parallelize([123,"Katie",19,"Brown"])

schema = StructType([
StructField("id",LongType(), True),
StructField("name", StringType(), True),
StructField("age", LongType(), True),
StructField("eyeColour", StringType(), True)
])
StructField 클래스는 name, dataType, nullable로 표시한다. 테이블같다고 생각하면 편하다.



person = spark..createDataFrame(stringCSVRDD, schema)

person.createOrReplaceTempView("person")
```



## spark sql

스파크 문법이 헷갈리고 뭔 메서드가 있었는지 기억이 안나면 sql 문법을 사용 가능하다.


```
from pyspark.sql import SQLContext, Row

hiveContext = HIiveContext(sc)
inputData = spark.rea.json(dataFile)
inputData.createOrReplaceTempView("myStructuredStuff")
myResultDataFrame = hiveContext.sql("select * from bar oder by desc")
```

이런 식으로 메모리에 뷰 테이블을 만들어서 sql처럼 사용할 수 있다. sql을 다루던 개발자들에게는 아주 편하다.

물론 `show(), select(), filter(), groupBy(), rdd().map()`등의 메서드들도 사용 가능하다.

### udf(user defined functions)

특별한 연산을 하고 싶은데 그런게 없으면 이런식으로 자기가 만들어서 쓸 수 도 있다.

```
from pyspark.sql.types import IntegerType

hiveCtx.registerFunction("square", lambda x: x*x, IntegerType()) # 제곱 함수 만들기
df = hiveCtx.sql("select square('field') from tableName")
```

### 예제

```
from pyspark.sql import SparkSession
from pyspark.sql import Row

import collections

# sparkcontext대신 sparksession 을 만들고있다.
# ***윈도우에서 할때는 .config를 넣어준다. (버그인것 같다. 윈도우에서만 필요하게 된다 ㅡㅡ)
spark = SparkSession.builder.config("spark.sql.warehouse.dir", "file:///C:/temp").appName("SparkSQL").getOrCreate()


# csv 파일을 쓸만한 값으로 재정비
def mapper(line):
    fields = line.split(',')
    return Row(ID=int(fields[0]), name=str(fields[1].encode("utf-8")), age=int(fields[2]), numFriends=int(fields[3]))

lines = spark.sparkContext.textFile("fakefriends.csv")
people = lines.map(mapper)


# 만든 rdd를 dataframe으로 바꾸고 임시 뷰 테이블 생성
schemaPeople = spark.createDataFrame(people).cache()
schemaPeople.createOrReplaceTempView("people")

#  만든 그 테이블로 sql 쿼리를 사용가능
teenagers = spark.sql("SELECT * FROM people WHERE age >= 13 AND age <= 19")

# sql쿼리의 결과는 다 rdd라 rdd 연산이 가능함
for teen in teenagers.collect():
  print(teen)

# 물론 sql 쿼리대신 함수를 써도 됨
schemaPeople.groupBy("age").count().orderBy("age").show()

spark.stop()

```

# dataset?

dataframe이 2.0부터 dataset으로 바뀌어가고 있다. 데이터셋은 데이터프레임의 진화형 같은 것이지만... typed data이다.
그런데... **파이썬은 untype 랭귀지 이다...**
그러므로 파이스파크를 쓸때는 골머리 썩히지 말고 그냥 스파크 버전 2.0 이상을 사용한다면 dataset을 dataframe대신 써주면 되겠다.
MLLib이랑 spark streaming은 이 dataset을 기반해서 만들어졌다.

