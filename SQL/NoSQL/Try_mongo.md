# Mongo 심플 쿼리


- 현재 db 콜렉션 보기 `show collections`
- 테이블 개수 찾기 `db.users.count()`

- 검색 `db.foo.find()`
- 좀 보기좋게 검색 `db.foo.find.pretty()`
- 한개만 검색 `findOne({_id:1}) `
- 정렬 검색 `find().sort({“-_id”, -1}) // descending`
- objectID 로 머리아파하지 말기 `ObjectId().getTimestamp()`


- 필드의 범위로 검색 예제

gt, gte
gt : 크다, gte : 크거나 같다

lt, lte
lt : 작다, lte : 작거나 같다

파이썬으로 치자면 이런 느낌이다.

```
for doc in actor_collection.find({'actor_rate' : {'$gte' : 40000}}):
    print(doc)
```

Find는 전부 찾아버리니까 충격적인 여파(최악 : 터짐)를 불러오므로 aggregate를 사용해 보자.

```
db.channel_msg.aggregate([ // []안에 넣을것
    
    {$match : {}}, // where
    {$project : {}}, // select
    {$group : {}}, // group by
    {$sort : {}} // order by
    
    ])    
 ```

 대강 sql과 비교하자면 아래와 같다.

```
 where -> $match
 group by -> $group
 having -> $having
 select -> $project
 order by -> $sort
 limit -> $limit
 sum() -> $sum
 count() -> $sum
 join() -> $lookup
```

 이 aggregate는 pipeline을 짜는데 아아아아아아아주 기이이이이인 쿼리잉이 되기 십상이니 괄호 오류등으로 멘탈이 부숴지지 않게 조심해야 한다.

 aggregation pipeline 연산자들은 공웹에서 보고 레퍼런스를 찾아서 쓰자.. 외우지도 못하겠음 https://docs.mongodb.com/manual/reference/operator/aggregation/

 또 몽고는 row단위가 아니라 데이터베이스 단위 락을 사용하니까 write lock을 주의해야 한다.(write 하는데 한 오퍼레이션만 적용됨) insert를 미친듯이 많이 한번에 하면 다른작업을 못한다는 뜻이다.

 3.2부터 aggregate내에서 $group 하고 $project안에서 $sum을 ㅆㄹ 수 있게 되었다. 원래는 $group에서만 썼었다.

 ```
 db.sales.aggregate(
   [
     {
       $group:
         {
           _id: { day: { $dayOfYear: "$date"}, year: { $year: "$date" } },
           totalAmount: { $sum: { $multiply: [ "$price", "$quantity" ] } },
           count: { $sum: 1 }
         }
     }
   ]
)
```

$sum사용시 값이 없는곳은 0으로 리턴하게 하려면 $qty를 사용한다.

```
db.sales.aggregate(
   [
     {
       $group:
         {
           _id: { day: { $dayOfYear: "$date"}, year: { $year: "$date" } },
           totalAmount: { $sum: "$qty" },
           count: { $sum: 1 }
         }
     }
   ]
)
```

몽고db 컬렉션은 기본적으로 _id에 인덱스가 존재한다. 따로 _id를 안정해주면 objectid를 _id로 정한다. ObjectId는 같은 document 내에서 유일함이 보장되는 12 byte binary data다. timestamp를 기준으로 생성되니 아래 사이트들을 이용하자.


[시간변환(unix > 인간)](https://www.epochconverter.com/)
[날짜별 objectID 조회하기](https://steveridout.github.io/mongo-object-time/)
[Timestamp 만드는데](http://www.timestampgenerator.com/1533081599/#result)






