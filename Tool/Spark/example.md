[ENG]


# spark analysis

to Analyse our chatbot tour log data, i should know the configuration of data.

the tour log data is saved in this format
`viewfs:///platform/italk/wholelog/svc_code=tour/year=2018/month=01/day=30`

the tour log looks like

```

[Row(service='italk', 
room_id=None, 
user=None, 
channel={
    'disp_id': 'overseas_air', v
    'disp_nm': '해외항공/공통',
    'disp_no': '9100'
    
},
message={
    'type': 'commit',
    'msg_list': '[{
        "type":"text",
        "body":"항공기 탑승이나 수화물 등 공항서비스 관련 궁금한 사항을 선택해주세요~\\r\\n"}]'
    
},

writer={
    'user_id': 'tour',
    'app_os': '',
    'app_ver': '',
    'user_nm': 'alfred',
    'mem_no': '1',
    'user_tp': 'system', 
    'app_id': 'tour'
    
}, 

room={
    'jobseq': '395265',
    'id': 'c914a60c',
    'status': 'auto'
    
},

zipsa=None,

logdate={
    'date': '2018-01-03',
    'time': '00:17:52',
    'ts': '1514906272'
    
}, 

hour='00',
timestamp='1514906272',
logtype='message')]
"""

# room.status means
"""
    'auto'        : '자동대화', // 10
    'no-zipsa'    : '분배대기', // 14
    'night'       : '야간상담', // 15
    'wait'        : '미응대',  // 19
    'ing'         : '대화중',  // 20
    'holding'     : '응대대기', // 30
    'no-answer'   : '무응답',  // 31
    'end-passive' : '대화종료', // 40
    'giveup'      : '대화포기', // 41
    'end-auto'    : '자동종료', // 42
    'transfer'    : '대화이관'  // 50
"""


# how to get a average waittime
"""
room={
    'jobseq': '395265',
    'id': 'c914a60c',
    'status': 'auto'
    
},

zipsa=None,

logdate={
    'date': '2018-01-03',
    'time': '00:17:52',
    'ts': '1514906272'
    
}, 
```




## 1.  get count of each consulting

in chatbot, each message is saving in log and have a unique id in counsult. it has  `timestamp` value, so we can get a count and average wait time.

```
from pyspark.sql.types import *
from pyspark.sql import functions as F
from pyspark.storagelevel import StorageLevel
from datetime import datetime, timedelta
import calendar
from datetime import date
from pyspark import SparkConf, SparkContext


yesterday = datetime.now() + timedelta(days=-1)

def getLeadZeroFormat(no):
	"""
	for formatting sql date form
	"""
    return (lambda x: '0'+str(x) if x < 10 else str(x))(no)
    
sample_path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(yesterday.year),"/month=",getLeadZeroFormat(yesterday.month),"/day=",getLeadZeroFormat(yesterday.day)])

def getDayLogPath(year, month, day) :
	
    path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(year),"/month=",getLeadZeroFormat(month),"/day=",getLeadZeroFormat(day)])
    return path

def getMonthLogPath(y, m) :
    pathList = []
    month = date(y, m, 1)
    last_day_of_month = calendar.monthrange(y, m)[1]
    
    for i in range(1, last_day_of_month+1) :
        day_path = getDayLogPath(y, m, i)
        pathList.append(day_path)
        
    return pathList
```

setup a boilerplate code first, and from now on let's get yesterday's log from hadoop.

```
from pyspark.sql.types import *
from pyspark.sql import functions as F
from pyspark.storagelevel import StorageLevel
from datetime import datetime, timedelta

yesterday = datetime.now() + timedelta(days=-1)

target_path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(yesterday.year),"/month=",getLeadZeroFormat(yesterday.month),"/day=",getLeadZeroFormat(yesterday.day)])

df = spark.read.parquet(target_path)

raw = df.select(
    df.channel.disp_id.alias('disp_id'), 
    df.channel.disp_nm.alias("disp_nm"),
    df.message.msg_list.alias("message"),
    F.concat_ws('-', df.logdate.date, df.room.id, df.room.jobseq).alias('msg_idx'), # from pyspark.sql import functions as F
    df.room.status.alias('status'), 
    df.writer.user_tp.alias("user_type"),
    df.writer.user_age.alias("user_age"),
    df.writer.app_os.alias('app_os'), 
    df.writer.user_id.alias('user_id'),
    df.writer.user_nm.alias('user_nm'),
    df.logdate.date.alias('date'),
    df.hour.alias('hour'),
    df.timestamp.alias('ts')
)

raw.persist(StorageLevel.MEMORY_ONLY_SER)

raw = raw.filter(raw.user_age != "00") #
raw.show()
```

if you use `.collect()` always, it can cause performance degradation. so you shouldn't use the `.collect()` much.

from this data, now we can get a max and min time for waiting.

```
def pickmax(x,y):
    if x['timestamp'] > y['timestamp']:
        return x
    else:
        return y
    
def pickmin(x,y):
    if x['timestamp'] < y['timestamp']:
        return x
    else:
        return y

def pickauto(raw) :
    raw.filter(raw.status == "auto")
    return raw
    
autordd = pickauto(raw)
print(autordd)
"""
DataFrame[disp_id: string, disp_nm: string, message: string, msg_idx: string, status: string, user_type: string, user_age: string, app_os: string, user_id: string, user_nm: string, date: string, hour: string, ts: string]
"""        
                
max_auto = raw.filter(raw.status == 'auto') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)

max_auto.unpersist()
max_auto.persist(StorageLevel.MEMORY_ONLY_SER)

min_auto = raw.filter(raw.status == 'auto') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_auto.unpersist()
min_auto.persist(StorageLevel.MEMORY_ONLY_SER)

max_wait = raw.filter(raw.status == 'wait') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_wait.unpersist()
max_wait.persist(StorageLevel.MEMORY_ONLY_SER)

min_wait =  raw.filter(raw.status == 'wait') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_wait.unpersist()
min_wait.persist(StorageLevel.MEMORY_ONLY_SER)

max_ing = raw.filter(raw.status == 'ing') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_ing.unpersist()
max_ing.persist(StorageLevel.MEMORY_ONLY_SER)

min_ing = raw.filter(raw.status == 'ing') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_ing.unpersist()
min_ing.persist(StorageLevel.MEMORY_ONLY_SER)

max_night = raw.filter(raw.status == 'night') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_night.unpersist()
max_night.persist(StorageLevel.MEMORY_ONLY_SER)

min_night = raw.filter(raw.status == 'night') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_night.unpersist()
min_night.persist(StorageLevel.MEMORY_ONLY_SER)

max_giveup = raw.filter(raw.status == 'giveup') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_giveup.unpersist()
max_giveup.persist(StorageLevel.MEMORY_ONLY_SER)

min_giveup = raw.filter(raw.status == 'giveup') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
```

after this, we can make a temp view to get things from sql.
to make a table, you can code it like

```
def getDiffTs(val):
    x = val[1]
    y = x[0]
    z = x[1]
    
    dp = y['disp_id']
    if y['disp_id'] == 'common':
        dp = z['disp_id']

    a = int(y['timestamp'])
    b = int(z['timestamp'])
    
    date_a = datetime.fromtimestamp(a)
    date_b = datetime.fromtimestamp(b)
    
    if a >= b:
        diff = date_a-date_b
    else:
        diff = date_b-date_a
    return [val[0], dp, z['app_os'], z['user_age'], z['date'], z['hour'], int(diff.total_seconds())]
 
autoToIng = max_auto.join(min_ing).map(getDiffTs)

joins_df = spark.createDataFrame(autoToIng, ["msg_idx","disp_id","app_os", "user_age", "date","hour","waitsec"])
print(joins_df.take(1))

joins_df.createOrReplaceTempView("test")
```

you can check the table view from code `
joins_df.createOrReplaceTempView("test")`

Let's count yesterday's all log in specific  condition.

```

from pyspark.sql.functions import countDistinct


ing =raw.filter(raw.status=="ing").select("msg_idx").distinct().count()
no_answer = raw.filter(raw.status == 'no-answer').select("msg_idx").distinct().count()
holding = raw.filter(raw.status == 'holding').select("msg_idx").distinct().count()

notEndedYetCount = ing+no_answer+holding
print("notEndedYetCount", notEndedYetCount)

giveup = raw.filter(raw.status == "giveup").select("msg_idx").distinct().count()
end_passive = raw.filter(raw.status == "end-passive").select("msg_idx").distinct().count()
distributedCount = notEndedYetCount + giveup+end_passive
print("distributedCount", distributedCount)

no_zipsa = raw.filter(raw.status == "no-zipsa").select("msg_idx").distinct()
print("no_zipsa", no_zipsa.count())
wait= raw.filter(raw.status == "wait").select("msg_idx").distinct()
print("wait", wait.count())
manualCount = no_zipsa.count() + wait.count() + distributedCount
print("manualCount", manualCount)

end_auto =  raw.filter(raw.status == "end-auto").count()

autoCount = giveup+ end_auto
print("autoCount", autoCount)


todayCount = autoCount+manualCount
print("todayCount", todayCount)


responseRate = distributedCount/manualCount
print("responseRate", round(responseRate*100, 3), "%")
```


the result seems like

```
notEndedYetCount 572
distributedCount 594
no_zipsa 93
wait 466
manualCount 1153
autoCount 22
todayCount 1175
responseRate 51.518 %
```

##  2.  Count user diversity

shame for me, i count a foriegn user for who's name includes alphabet.
but in korea, almost 100% of Korean people use korean for their name.

when i checked for result, all non-korean name user was foreigners.

```
import re
internationalUser = raw.filter(raw.user_nm.rlike("[^\W_]")) 

internationalUserCount = internationalUser.count()
wholeUserCount = raw.count()
domesticUserCount = wholeUserCount - internationalUserCount


internationalUserRate = internationalUserCount/wholeUserCount
print("internationalUserRate", round(internationalUserRate*100, 3), "%")

print("internationalUserCount", internationalUserCount)
print("domesticUserCount", domesticUserCount) 
```

the result was like this.

```
internationalUserRate 2.624 %
internationalUserCount 312
domesticUserCount 11579
```

## 3. get most busy hour


```
hourlyRaw =raw.select("hour").rdd

hourlyRawCounts = hourlyRaw.map(lambda hour : (hour, 1)).reduceByKey(lambda a, b : a+b)
hourlySorted = hourlyRawCounts.map(lambda x : (x[1], x[0])).sortByKey()
results = hourlySorted.collect()

for x in results :
    count = x[0]
    hour = x[1]
    
    if(hour) :
        print(hour, ": ", count, "회")
```

am 9-11 was extremely busy in korean utc.

```
Row(hour='04') :  4 회
Row(hour='05') :  18 회
Row(hour='02') :  35 회
Row(hour='07') :  39 회
Row(hour='03') :  43 회
Row(hour='00') :  72 회
Row(hour='06') :  76 회
Row(hour='01') :  80 회
Row(hour='21') :  150 회
Row(hour='08') :  171 회
Row(hour='19') :  195 회
Row(hour='20') :  215 회
Row(hour='23') :  238 회
Row(hour='22') :  265 회
Row(hour='18') :  339 회
Row(hour='17') :  581 회
Row(hour='12') :  889 회
Row(hour='14') :  1056 회
Row(hour='13') :  1071 회
Row(hour='16') :  1177 회
Row(hour='11') :  1194 회
Row(hour='10') :  1267 회
Row(hour='15') :  1317 회
Row(hour='09') :  1399 회
```

## 4. use nlp

```
from konlpy.tag import Twitter
import sys
import operator
from collections import Counter

from pyspark.sql.types import StringType

placeRaw = raw.select("disp_nm")
```

i do want to get the percentage of domestic and international chats in bot.

```
from konlpy.tag import Twitter
import sys
import operator
from collections import Counter

from pyspark.sql.types import StringType

placeRaw = raw.select("disp_nm")

collectedRaw = placeRaw.rdd.map(lambda x : (x,1)).reduceByKey(operator.add).map(lambda x : ["국내" if ("국내" in x[0].disp_nm) else "해외", x])



placeCount = collectedRaw.countByKey()
print(placeCount)
```

i know the `placeRaw.rdd.map(lambda x : (x,1)).reduceByKey(operator.add).map(lambda x : ["국내" if ("국내" in x[0].disp_nm) else "해외", x])` code seems complicated, but it's just a key and value calculation.

countByKey() returns  defaultdict and it looks like `defaultdict(<class 'int'>, {'해외': 49, '국내': 5})`  -> domestic for 5 and international for 49.

now i want to check which country is most famous for tour app.

```
from konlpy.tag import Kkma
from konlpy.utils import pprint
from konlpy.tag import Twitter


def get_tags(text, ntags=100) :
    
    kkma = Kkma()
    twitter = Twitter() 
    return_list = []
    
    nouns = twitter.nouns(text)
    count = Counter(nouns)
    
    for n, c in count.most_common(ntags):
        temp = {'tag': n, 'count': c}
        return_list.append(temp)
    return return_list


internationalRaw = placeRaw.rdd.map(lambda x : (x,1)).filter(lambda x : ~("국내" in x[0].disp_nm))


placeKeys = internationalRaw.keys().keys().collect()
text = "".join(placeKeys)

tags = get_tags(text)


reduceText = "해외, 항공, 공통, 해외여행, 실시간, 임박, 호텔, 숙박, 담당, 기타, 팀, 지역, 패키지, 출발, 자유, 만들기, 현지, 긴급, 일, 법인, 단체, 나라"
reduceArray = reduceText.split(", ")

for tag in tags:
    noun = tag['tag']
    count = tag['count']
    
    if noun not in reduceArray :
        print('{} {}'.format(noun, count))   

```

the result is like 

```
동남아 1139
일본 1080
유럽 666
미주 535
일본국 278
대양주 184
중국 136
필리핀 79
북부 47
.. etc
```

the whole process tooks about 3.45mins

## further Fixes

1. I should do word classification with python library. I'm still newbie at nlp.
2. i can use pandas than spark for speed. or i can think consider about cython.

---

# spark 분석 적용하기 ~ 톡집사 투어 로그 ~

여태 공부한 spark를 실제 업무에 적용해 보도록 하겠다. 간단한 예로 톡집사 투어 로그를 분석해보겠다.

톡집사 투어 로그는 매일매일 쌓이고 있고 아래와 같은 형식으로 백업되고 있다.

`viewfs:///platform/italk/wholelog/svc_code=tour/year=2018/month=01/day=30`

투어 로그는 아래와 같이 구성된다.

```

[Row(service='italk', 
room_id=None, 
user=None, 
channel={
    'disp_id': 'overseas_air', v
    'disp_nm': '해외항공/공통',
    'disp_no': '9100'
    
},
message={
    'type': 'commit',
    'msg_list': '[{
        "type":"text",
        "body":"항공기 탑승이나 수화물 등 공항서비스 관련 궁금한 사항을 선택해주세요~\\r\\n"}]'
    
},

writer={
    'user_id': 'tour',
    'app_os': '',
    'app_ver': '',
    'user_nm': 'alfred',
    'mem_no': '1',
    'user_tp': 'system', 
    'app_id': 'tour'
    
}, 

room={
    'jobseq': '395265',
    'id': 'c914a60c',
    'status': 'auto'
    
},

zipsa=None,

logdate={
    'date': '2018-01-03',
    'time': '00:17:52',
    'ts': '1514906272'
    
}, 

hour='00',
timestamp='1514906272',
logtype='message')]
"""

# 대화상태(room.status)
"""
    'auto'        : '자동대화', // 10
    'no-zipsa'    : '분배대기', // 14
    'night'       : '야간상담', // 15
    'wait'        : '미응대',  // 19
    'ing'         : '대화중',  // 20
    'holding'     : '응대대기', // 30
    'no-answer'   : '무응답',  // 31
    'end-passive' : '대화종료', // 40
    'giveup'      : '대화포기', // 41
    'end-auto'    : '자동종료', // 42
    'transfer'    : '대화이관'  // 50
"""


# 대기시간
"""
room={
    'jobseq': '395265',
    'id': 'c914a60c',
    'status': 'auto'
    
},

zipsa=None,

logdate={
    'date': '2018-01-03',
    'time': '00:17:52',
    'ts': '1514906272'
    
}, 
```

위 정보를 기반으로 아래와 같은 값을 계산해보려고 한다.

1.  대화상태별 상담 수
2.  응답율
3. 외국인, 국내 유저 수
4. 시간별 상담 수
5. 여행 장소 분석 (국내, 해외 등...)


## 1. 대화상태별 상담 수

각 상담별로 타임스탬프가 `timestamp` 로 찍히므로, 이 값을 기준으로 대기시간을 구해 보도록 한다.

```
from pyspark.sql.types import *
from pyspark.sql import functions as F
from pyspark.storagelevel import StorageLevel
from datetime import datetime, timedelta
import calendar
from datetime import date
from pyspark import SparkConf, SparkContext


# [샘플] 어제 로그파일 가져오기
yesterday = datetime.now() + timedelta(days=-1)

def getLeadZeroFormat(no):
	"""
	sql 날짜 포맷을 맞추기 위한 함수
	"""
    return (lambda x: '0'+str(x) if x < 10 else str(x))(no)
    
sample_path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(yesterday.year),"/month=",getLeadZeroFormat(yesterday.month),"/day=",getLeadZeroFormat(yesterday.day)])

def getDayLogPath(year, month, day) :
	"""
	특정 일 로그파일 가져오기
	"""
    path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(year),"/month=",getLeadZeroFormat(month),"/day=",getLeadZeroFormat(day)])
    return path

def getMonthLogPath(y, m) :
	"""
	특정 달 로그파일 가져오기
	"""
    pathList = []
    month = date(y, m, 1)
    last_day_of_month = calendar.monthrange(y, m)[1]
    
    for i in range(1, last_day_of_month+1) :
        day_path = getDayLogPath(y, m, i)
        pathList.append(day_path)
        
    return pathList
```

위와 같이 보일러플레이트를 준비하고(제플린 안에서는 스파크콘텍스트를 따로 만들 필요가 없어 만들지 않았다) 분석을 원하는 날짜의 로그를 가져온다. 편의를 위해 어제 날짜를 기준으로 분석해 본다.

```

from pyspark.sql.types import *
from pyspark.sql import functions as F
from pyspark.storagelevel import StorageLevel
from datetime import datetime, timedelta

# 어제날짜
yesterday = datetime.now() + timedelta(days=-1)

target_path = "".join(["viewfs:///platform/italk/wholelog/svc_code=tour/","year=",str(yesterday.year),"/month=",getLeadZeroFormat(yesterday.month),"/day=",getLeadZeroFormat(yesterday.day)])

df = spark.read.parquet(target_path)

# 어제날짜 내용 중 내게 필요한 값만 선택
raw = df.select(
    df.channel.disp_id.alias('disp_id'), 
    df.channel.disp_nm.alias("disp_nm"),
    df.message.msg_list.alias("message"),
    F.concat_ws('-', df.logdate.date, df.room.id, df.room.jobseq).alias('msg_idx'), # from pyspark.sql import functions as F
    df.room.status.alias('status'), 
    df.writer.user_tp.alias("user_type"),
    df.writer.user_age.alias("user_age"),
    df.writer.app_os.alias('app_os'), 
    df.writer.user_id.alias('user_id'),
    df.writer.user_nm.alias('user_nm'),
    df.logdate.date.alias('date'),
    df.hour.alias('hour'),
    df.timestamp.alias('ts')
)

raw.persist(StorageLevel.MEMORY_ONLY_SER)

raw = raw.filter(raw.user_age != "00") #
raw.show()
```

실수로 `.collect()`를 해버리면 아주아주 아주많이 느려지기 때문에, `collect()`연산은 소량 데이터에만 적용하자. 이렇게 하면 대강의 데이터의 모습이 어떻게 생겼는지 확인할 수 있다.

![로그그림]()

로그 데이터는 지금 이렇게 생겼다.

이 로그 데이터를 기반으로, 각 상태에서의 대기시간의 최대 / 최소값을 구할 수 있다.

```
def pickmax(x,y):
	"""
	제일 오래걸린 대기시간 찾기
	"""
    if x['timestamp'] > y['timestamp']:
        return x
    else:
        return y
    
def pickmin(x,y):
	"""
	제일 짧은 대기시간 찾기
	"""
    if x['timestamp'] < y['timestamp']:
        return x
    else:
        return y

def pickauto(raw) :
	"""
	대화가 시작된(자동대화) 인 상태를 찾으면, 이 사람이 언제 앱에 들어왔는지 알 수 있다
	"""
    raw.filter(raw.status == "auto")
    return raw
    
autordd = pickauto(raw)
print(autordd)
"""
DataFrame[disp_id: string, disp_nm: string, message: string, msg_idx: string, status: string, user_type: string, user_age: string, app_os: string, user_id: string, user_nm: string, date: string, hour: string, ts: string]
"""        
                
max_auto = raw.filter(raw.status == 'auto') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)

max_auto.unpersist()
max_auto.persist(StorageLevel.MEMORY_ONLY_SER)

min_auto = raw.filter(raw.status == 'auto') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_auto.unpersist()
min_auto.persist(StorageLevel.MEMORY_ONLY_SER)

max_wait = raw.filter(raw.status == 'wait') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_wait.unpersist()
max_wait.persist(StorageLevel.MEMORY_ONLY_SER)

min_wait =  raw.filter(raw.status == 'wait') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_wait.unpersist()
min_wait.persist(StorageLevel.MEMORY_ONLY_SER)

max_ing = raw.filter(raw.status == 'ing') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_ing.unpersist()
max_ing.persist(StorageLevel.MEMORY_ONLY_SER)

min_ing = raw.filter(raw.status == 'ing') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_ing.unpersist()
min_ing.persist(StorageLevel.MEMORY_ONLY_SER)

max_night = raw.filter(raw.status == 'night') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_night.unpersist()
max_night.persist(StorageLevel.MEMORY_ONLY_SER)

min_night = raw.filter(raw.status == 'night') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
    
min_night.unpersist()
min_night.persist(StorageLevel.MEMORY_ONLY_SER)

max_giveup = raw.filter(raw.status == 'giveup') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmax)
    
max_giveup.unpersist()
max_giveup.persist(StorageLevel.MEMORY_ONLY_SER)

min_giveup = raw.filter(raw.status == 'giveup') \
    .rdd.map(lambda x: (x['msg_idx'], {
        'disp_id' : x['disp_id'], 
        'app_os' : x['app_os'], 
        'user_age' : x['user_age'],
        'date' : x['date'], 
        'hour' : int(x['hour']), 
        'timestamp' : int(x['ts'])
    })).reduceByKey(pickmin)
```

차후 이렇게 생성한 각 대화상태별 최대/최소값 풀에서 temp view를 만들어서, 각 상태별 평균 대기시간 등을 구할 수 있다.

```
def getDiffTs(val):
	"""
	각 상태별 최대최소 시간차를 구한다.
	"""
    x = val[1]
    y = x[0]
    z = x[1]
    
    dp = y['disp_id']
    if y['disp_id'] == 'common':
        dp = z['disp_id']

    a = int(y['timestamp'])
    b = int(z['timestamp'])
    
    date_a = datetime.fromtimestamp(a)
    date_b = datetime.fromtimestamp(b)
    
    if a >= b:
        diff = date_a-date_b
    else:
        diff = date_b-date_a
    return [val[0], dp, z['app_os'], z['user_age'], z['date'], z['hour'], int(diff.total_seconds())]
 

#  대기 시간 : 상담신청에서 실제 상담사가 상담을 시작하기 전까지의 시간이다.

autoToIng = max_auto.join(min_ing).map(getDiffTs)

joins_df = spark.createDataFrame(autoToIng, ["msg_idx","disp_id","app_os", "user_age", "date","hour","waitsec"])
print(joins_df.take(1))

joins_df.createOrReplaceTempView("test")
```

이렇게 생성한 테이블은 아래와 같이 생겼다.

![너무커서짤린그림]()

이제 각 대화 상태 별 개수 또한 세어볼 수 있다.

```

from pyspark.sql.functions import countDistinct

# raw -> 전체대화(어제)

# 미종료대화
# 미종료 = 대화중ing + 무응답no-answer + 응대대기holding	

# 1. null처리 fillna
# 2. 대화중 26? 15? 41? 과 비슷한값이 나와야함
#ing = raw.filter(raw.status == 'ing').filter(df.room_id.isNotNull()).agg(countDistinct(df.room_id))


ing =raw.filter(raw.status=="ing").select("msg_idx").distinct().count()
no_answer = raw.filter(raw.status == 'no-answer').select("msg_idx").distinct().count()
holding = raw.filter(raw.status == 'holding').select("msg_idx").distinct().count()

notEndedYetCount = ing+no_answer+holding
print("notEndedYetCount", notEndedYetCount)


# 분배대화 = 대화포기giveup + 대화종료end-passive + 미종료

giveup = raw.filter(raw.status == "giveup").select("msg_idx").distinct().count()
end_passive = raw.filter(raw.status == "end-passive").select("msg_idx").distinct().count()
distributedCount = notEndedYetCount + giveup+end_passive
print("distributedCount", distributedCount)


# 수동대화 = 분배대기no-zipsa + 미응대wait + 분배대화
# 이쪽에 문제가 있는듯함.
# 응대율 = 분배대화/수동대화

no_zipsa = raw.filter(raw.status == "no-zipsa").select("msg_idx").distinct()
print("no_zipsa", no_zipsa.count())
wait= raw.filter(raw.status == "wait").select("msg_idx").distinct()
print("wait", wait.count())
manualCount = no_zipsa.count() + wait.count() + distributedCount
print("manualCount", manualCount)


# 자동대화 = 대화포기giveup + 자동종료end-auto
end_auto =  raw.filter(raw.status == "end-auto").count()

autoCount = giveup+ end_auto
print("autoCount", autoCount) #223인데 대강 비슷


# 당일인입 = 자동대화 + 수동대화

todayCount = autoCount+manualCount
print("todayCount", todayCount)



# 응대율 구하기

responseRate = distributedCount/manualCount
print("responseRate", round(responseRate*100, 3), "%")
```

결과를 구해보면 아래와 같이 나오고 있다.

```
notEndedYetCount 572
distributedCount 594
no_zipsa 93
wait 466
manualCount 1153
autoCount 22
todayCount 1175
responseRate 51.518 %
```

실제 결과와의 차이가 있는 이유는 아래와 같이 예상하고 있다.

- 실시간으로 가져와서 보는 것이 아님
- 과거의 로그이기 때문에 몇가지 값이 자동으로 다른 값으로 변경 저장된것이 있음(예 - 대화종료 등)
- 본인의 부족함(????)

##  2. 외국인, 국내 유저 수

정말 단순하게 이름에 알파벳이 들어가면 외국인이라고 카운트했다.
놀랍게도 어느정도 들어맞은 것 같다.(이용하는 사람이 한글 또는 영어로 이름을 입력하기 때문에...)

```
import re
internationalUser = raw.filter(raw.user_nm.rlike("[^\W_]")) # 영어 매치

internationalUserCount = internationalUser.count()
wholeUserCount = raw.count()
domesticUserCount = wholeUserCount - internationalUserCount


internationalUserRate = internationalUserCount/wholeUserCount
print("internationalUserRate", round(internationalUserRate*100, 3), "%")

print("internationalUserCount", internationalUserCount)
print("domesticUserCount", domesticUserCount) 
```

결과는 아래와 같았다. 보통 1-2% 내외로 외국인 유저들이 존재했다.

```
internationalUserRate 2.624 %
internationalUserCount 312
domesticUserCount 11579
```

## 3. 시간별 상담 수

몇시에 제일 사람들이 붐볐을지에 대해 알아보도록 했다.
원 목표는 시간별 그래프를 만들기였는데 제플린에서 아직 dataframe으로 그래프를 만들게 못하는 것 같다.

```
hourlyRaw =raw.select("hour").rdd

hourlyRawCounts = hourlyRaw.map(lambda hour : (hour, 1)).reduceByKey(lambda a, b : a+b)
hourlySorted = hourlyRawCounts.map(lambda x : (x[1], x[0])).sortByKey()
results = hourlySorted.collect()

for x in results :
    count = x[0]
    hour = x[1]
    
    if(hour) :
        print(hour, ": ", count, "회")
```

결과는 아래와 같다.  오전 9-11시에 제일 바빴고 중간타임인 13-16시에도 1000건이 넘어가는 상담이 있었다.

```
Row(hour='04') :  4 회
Row(hour='05') :  18 회
Row(hour='02') :  35 회
Row(hour='07') :  39 회
Row(hour='03') :  43 회
Row(hour='00') :  72 회
Row(hour='06') :  76 회
Row(hour='01') :  80 회
Row(hour='21') :  150 회
Row(hour='08') :  171 회
Row(hour='19') :  195 회
Row(hour='20') :  215 회
Row(hour='23') :  238 회
Row(hour='22') :  265 회
Row(hour='18') :  339 회
Row(hour='17') :  581 회
Row(hour='12') :  889 회
Row(hour='14') :  1056 회
Row(hour='13') :  1071 회
Row(hour='16') :  1177 회
Row(hour='11') :  1194 회
Row(hour='10') :  1267 회
Row(hour='15') :  1317 회
Row(hour='09') :  1399 회
```

## 4. 여행 장소 분석 (국내, 해외 등...)


from konlpy.tag import Twitter
import sys
import operator
from collections import Counter

from pyspark.sql.types import StringType

#disp_nm 으로 여행자들의 비율을 알아보자!
placeRaw = raw.select("disp_nm")

## 국내 vs 해외 여행 비율을 알아보자!

국내여행과 해외여행을 가는 사람들의 비율과 행선지는 어떻게 되는지 알아보자.

```
from konlpy.tag import Twitter
import sys
import operator
from collections import Counter

from pyspark.sql.types import StringType

#disp_nm 으로 여행자들의 비율을 알아보자!
placeRaw = raw.select("disp_nm")

## 국내 vs 해외 여행 비율을 알아보자!
collectedRaw = placeRaw.rdd.map(lambda x : (x,1)).reduceByKey(operator.add).map(lambda x : ["국내" if ("국내" in x[0].disp_nm) else "해외", x])



placeCount = collectedRaw.countByKey()
print(placeCount)
```

placeRaw.rdd.map(lambda x : (x,1)).reduceByKey(operator.add).map(lambda x : ["국내" if ("국내" in x[0].disp_nm) else "해외", x]) 가 복잡해 보일 수도 있지만, 키별로 합산하고 그 중 국내여행은 국내라고 표시, 해외여행은 해외라고 표시했다.

countByKey()연산시 defaultdict가 리턴되며, 결과값은 `defaultdict(<class 'int'>, {'해외': 49, '국내': 5})` 와 같았다. 대강 국내가 해외의 1/10 정도 된다.

그 결과를 키별로 정리해 어느 해외에 제일 많이 사람들이 몰리는지 분석해보도록 하겠다.

```
## 해외여행별 국가 비율을 알아보자! 

from konlpy.tag import Kkma
from konlpy.utils import pprint
from konlpy.tag import Twitter


def get_tags(text, ntags=100) :
    
    kkma = Kkma()
    twitter = Twitter() 
    return_list = []
    
    nouns = twitter.nouns(text)
    count = Counter(nouns)
    
    for n, c in count.most_common(ntags):
        temp = {'tag': n, 'count': c}
        return_list.append(temp)
    return return_list


# 1. 해외인거만 고른다

internationalRaw = placeRaw.rdd.map(lambda x : (x,1)).filter(lambda x : ~("국내" in x[0].disp_nm)) # 람다식에서 !가 not이 아니고 ~가 not 임

# 2. 디스플레이 네임을 분석해서 국가 이름을 빼온다

placeKeys = internationalRaw.keys().keys().collect() # disp_nm만 빼온다.
text = "".join(placeKeys)

tags = get_tags(text) # 명사별 개수를 구한다.

for tag in tags:
    noun = tag['tag']
    count = tag['count']
    print('{} {}\n'.format(noun, count))
```

Twitter 모듈을 사용해서 명사별로 여행지를 분석해 본 결과, 아래와 같은 결과를 확인할 수 있었다.

```
해외 7029
항공 6976
공통 2862
동남아 1512
일본 1389
국내 813
해외여행 674
유럽 633
임박 625
미주 523
호텔 510
실시간 479
중국 357
숙박 328
대양주 300
...(생략)
```

이번에는 이 중 해외, 항공, 공통 ... 등을 빼고 실제 장소인 동남아, 일본, 유럽, 미주 등만 선택해보도록 하겠다. (how??)

