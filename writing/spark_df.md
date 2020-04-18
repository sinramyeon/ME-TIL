[스파크 기초 시리즈] 스파크 데이터프레임

아까까진 RDD 얘기를 해놓고 갑자기 데이터 프레임은 또 무엇인가?



우리는 스파크를 RDD 로 시작했지만, 2.0부터 스파크는 "데이터프레임" 의 형태를 사용하는 식을 좀 더 권장하고 있다. 데이터프레임이 더 쉽고 알아보기 쉽다고 한다.

심지어 요새는 또 데이터셋이란걸로 쓰라고 한다. rdd는 뭐고 데이터 프레임은 무어고 데이터 셋은 다 무언가?

물론 데이터프레임으로 된 데이터라고 해도, RDD의 형식을 띄고 있기도 하기 때문에 혼용이 아예 되지 않는 것은 아니다.

또 써보면 알겠지만 데이터프레임이 rdd보다 훨씬 빠르다는것을 느낄 수 있다.



스파크로 데이터프레임을 다루는 것을 알아보도록 한다. 우선은 데이터프레임이 뭔지 구경해 보자.



# zeppelin에서 실행중...
from pyspark.sql import SparkSession

spark = SparkSession.builder.appName("hello").getOrCreate()
df = spark.read.json("뭔가의 제이슨(로그 등등)")

# df = dataframe!

df.show()


이렇게 하면 json내에 있는 데이터를 테이블 스키마로 한 위에있는 20개정도 보여준다.

df.printSchema() 로 스키마를 확인해볼 수 있다.

df.columns() 로 컬럼을 확인해 볼 수 있고, 파이썬 리스트로 반환된다.

df.describe() 로 데이터프레임 내용을 확인해 볼 수 있다. 이 메서드로는 데이터프레임이 리턴되는것을 확인해볼 수 있다.

보다시피 rdd나 데이터 프레임이나 똑같은 데이터 객체이다. 속도만 다른!



만약 데이터프레임의 스키마가 아주 복잡한 데이터일 경우에는, 이렇게 다룰 수 있다.

from pyspark.sql.types import (StructField,StringType,
IntgerType,StructType)


data_schema = [StructField('age', IntegerType(), True),]
#                       필드이름,    필드타입,      널가능여부?
StructField('name', StringType(), True)]

final_struct = StructType(fields = data_schema)

df = spark.read.json("파일", schema = final_struct)
df.printSchema()
# 이런식으로 하면 고에서 제이슨을 스트럭트안에 마샬링해 넣던것같은 일이 가능하다!


그렇다면 이렇게 얻은 데이터에서 어떻게 자신이 원하는 내용을 얻을 수 있을까?



df['age'] 로 인덱싱을 해 보면 파이스파크 컬럼을 얻을 수 있다. 이 내용을 sql안 메서드들로 골라올 수 있다.



df.select('age').show() 로 위 컬럼 안 든 데이터를 가져올 수 있다.



df.head(2) 로 위에 있는 두 가지 row를 가져올 수 있다. df.head(2)[0] 의 식으로 파이썬에서 하던 인덱싱을 할 수도 있다. 아주 간편하다.



df.select(['age', 'name']) 의 식으로 다중 컬럼을 선택할 수도 있다.



컬럼을 좀 더 다뤄보자.



df.withColumn("새컬럼", df['age']) 으로 새 컬럼을 만들 수 있다. withColumn() 은 아규먼트로 컬럼 이름과, sql.컬럼을 받기 때문에 주의한다.



df.withColumn("두배', df['age']*2) 와 같이 컬럼 내 데이터들을 손봐서 새 컬럼을 추가하는 데 사용할 수 있다.



df.withColumnRenamed("age", "renamed_age").show() 로 컬럼 이름을 바꾼 후 내용을 확인해볼 수 있다.



sql 모듈에서는 sql을 다루는 쉬운 메서드들이 많다.



df.createOrReplaceTempView("새로만들 테이블") 로 테이블을 만들듯 sql 임시 뷰를 만들 수 있다.



results = spark.sql("select * from 새로만들 테이블") 식으로, sql을 다루듯 사용할 수 있으므로 별다른 많은 공부를 요구하지 않는다. 아주 쉽다. 아래와 같이 제플린에서 사용한다면 될 것이다.



df.createOrReplaceTempView("temp")
results = spark.sql("select * from people where age = 30").show()


간단한 데이터프레임 연산을 배워보자.



from pyspark.sql import SparkSession as ss

spark = ss.builder.appname("야호").getOrCreate()

df = spark.read.csv("읽어올 파일", inferSchema = True, header=True)
# csv파일을 만약 다룬다면, header라는것은 csv맨 첫줄(컬럼 이름)을 뜻한다.

df.printSchema()
# 스키마 확인

df.show()
# 데이터 확인

df.head(3)[0]
# 처음 줄

df.filter("close < 500").show()
# close 라는 컬럼이 500 이하인것만 선택하여 보여줌

df.filter("close < 400").select("Open").show()
# close 라는 컬럼이 400 이하인 것 중 Open컬럼을 보여줌

df.filter(df["close"] < 500).show()
# sql 문법이 아닌 이런식으로 컬럼자체를 골라서 표현할수도 있다.(무엇이 더 빠른지는 모르겠는데 거의 비슷한 속도인것 같다. 둘다 0. xxxx s가 걸림)

# df.filter(df["close"] < 200 and df["open"] > 200).show()
# 이런식으로 컬럼을 불리언같이 쓸 수는 없다!

df.filter((df["close"] < 200) & (df["open"] > 200)).show()
# 항상 이런식으로 사용해야 한다.
# not 을 찾고싶으면 표현식 앞에 ~ 를 붙인다.
# ~(df["open"] >200) 식으로 표현하면 된다.

df.filter(df["low"] == 197.16).show()
# 특정한 값에 해당하는 내용을 찾는다.



이런식으로 .filter() 와 .show() 를 이용하여 특정 값에 해당하는 데이터를 빠르고 쉽게 확인할 수 있다. 만약 그냥 이렇게 보는것만이 아니라, 이 값을 기반으로 한 연산이 필요할때는 어떨까?



df.filter().collect() 의 collect() 메소드를 활용한다.

result = df.filter(df["low"] == 198).collect()
result[0]
# 첫번째 Row를 반환

result
# row 를 반환

row = result[0]
row.asDict()
# dict 자료형으로 바꿔 파이써닉하게 필요값을 활용할 수 있다.

spark dataframes groupby / aggregate

from pyspark.sql import SparkSession as ss

spark = ss.builder.appname("aggs").getOrCreate()
df = spark.read.csv("sth.csv"), inferSchema = True, header=True)

df.printSchema()
#스키마 확인

df.groupBy("Company")
# 해당 컬럼의 그룹된 객체 반환

df.groupBy("company").mean().show()
# mean, sum, max, min 등으로 간편히 평균, 합, 최대값, 최솟값을 알 수 있다.

df.agg({"Sales" : "sum"}).show()
# 딕셔너리를 활용하여 선택할 수 있다.
# agg 함수에 원하는 기능과 컬럼을 전달한다.


Groupby를 이용하면, 그룹을 지어(sql 과 같이) pyspark 를 이용할 수 있다.



group_data =df.groupBy("Company")
group_data.agg( {"Sales" : "max"} ).show()
# group by 에 max 를 전달한것과 똑같은 내용이 나오지만, 훨씬 간단하다.(속도는 잘 모르겠다. 똑같이 0.0000...s 가 걸렸다..)
# count, max 함수를 직접 쓰지 않아도 이렇게 aggregate를 간편히 할 수 있다.


다른 다양한 sql functions를 알아보자.



from pyspark.sql.functions import countDistinct, avg, stddev

df.select(countDistinct("Sales")).show()
# 원하는 함수, 컬럼을 패싱해 고통스럽게 sql을 쓰고있지 않아도 된다.

df.select(avg("Sales").alias("판매량평균").show()
# zeppelin에서 보기 쉽게 알리아스 설정을 해줄수도 있다.

df.select(stddev("Sales")).show()

from pyspark.sql.functions import format_number

# 파이썬에서 좀 더 깔끔하게 표현하기 위해 고쳐본다.

sales_std = df.select(stddev("Sales").alias("std"))
sales_std.select(format_number("std", 2))
#                        컬럼이름, 보여줄 나머지자릿수
# 이렇게 테이블 내 내용을 고쳐서 보여줄수도 있다.

df.orderBy("Sales").show()
# Sales 컬럼을 큰것부터 차례로 정리해서 보여준다.
# desc 로 하고 싶다면?

df.orderBy(df["Sales"].desc()).show()
# 조금 손에 익지 않지만 order by desc 연산 대신 파이스파크써닉하게 이렇게 쓰자.
# 컬럼자체를 넘겨준 후에 desc 메서드를 적용하여 쓰면 된다.

데이터에 빈 값이 있을때
데이터를 다루다 보니 빈 값이 있다. 이럴 때는

null 로 바꾼다
행 채로 버린다
다른 값으로 채운다
세 가지 옵션 중 선택할 수 있다. 이걸 pyspark로 해보겠다.



from pyspark.sql import SparkSession as ss

spark = ss.builder.appName("빈값").getOrCreate()
df = spark.read.csv("빈값이 있는 자료", header=True, inferSchema=True)

df.show()
빈 값이 있는 자료가 있다고 할 때, 아래와 같이 다룰 수 있다.



df.na.drop() 으로 그냥 null이 있는 곳은 버릴 수 있다.



df.na.drop(thresh=2).show() 로 thresh를 지정하면, null이 2개 이상인 곳은 버린다. tresh 옵션이 필요하다면 사용한다.



df.na.drop(how="any") 가 기본으로, 모든 null 을 drop 하지만, how="all" 을 넘겨주면, 모든 값이 null 일때 드랍한다.



df.na.drop(subset=["Sales"]) 로 특정 컬럼이 null 인 경우만 골라낼 수도 있다. 이렇게 하면 Sales 컬럼이 null 인 경우는 지운다.



df.na.fill("Fill value").show() 로 빈 값에 값을 넣을 수 있다. df.na.fill(0) 등으로 0을 넣거나 하는데, 여기서 알아둘 것이 있다.



pyspark에서는 python이 형 명시가 없는 것을 고려해 string을 넣으면 문자열 빈값에 넣어지고, int를 넣으면 숫자 빈값에 넣어준다. 0을 넣는다고 문자 null에도 0이 들어가있지 않는다. 아주 편리하다.



만약 그래도 못믿겠다, 또는 특정한 컬럼에만 값을 채워두고 싶다고 한다면 아까 썼던 subset을 활용한다.



df.na.fill("빈 이름", subset = ["name"])과 같이 name 컬럼이 비었다면 "빈 이름" 을 넣을 수 있다.



from pyspark.sql.funciton import mean

mean_val = df.select(mean(df["Sales"])).collect()
mean_sales = mean_val[0][0]
# sales mean 값

df.na.fill(mean_sales, ["Sales"])
위와 같이 비어있는 값에 계산값을 넣는 예제를 떠올려볼 수 있다.

date / timestamp


만약 파이썬 날짜시간모듈에 조금 덜 익숙하다면 이곳 과 python3 api 를 참조한다.

from pyspark.sql import SparkSession as ss

spark = ss.builder.appName("dates").getOrCreate()
df = spark.read.csv("a.csv", header=True, inferSchema=True)


스파크에서 datetime을 다뤄보자.



from pyspark.sql.functions import dayofmonth, hour, dayofyear, month, year, weekofyear, format_number, date_format 으로 온갖 모듈을 불러온다.



df.select(dayofmonth(df["date"])) 식으로 해당 날짜형식 컬럼의 달만 얻어올 수 있다.



항상 select() 에 함수를 arg로 넘겨줄 수 있음을 기억하자. 일급 함수를 기억해 보자.



df.withColumn("Year", year(df["date"])) 로 새 컬럼을 만들어 있는 년도를 모두 반환할 수 있다. 뷰처럼 활용하고 싶다면 withColumn() 을 기억하자.



 newdf = df.withColumn("Year", year(df["date"]))
 newdf.groupBy("Year").mean().select(["year", "avg(close)"]).show()
이렇게 하면 년도별 close값의 평균을 보여준다.
