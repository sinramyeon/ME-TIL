# python datetime

---

또 하나 그리운 것은 datetime 모듈과 캘린더, 타임, 데이트유틸등 각종 시간날짜 모듈이다.

## datetime 모듈

```
from datetime import date
halloween = date(2017, 10, 31)

halloween.day //31
halloween.month //10
halloween.year //2017

halloween.isoformat() //2017-10-31
```

보통 이런식으로 날짜를 계산한다.

`dt = datetime.datetime.now()`로 지금 시각을 알아 볼 수 있다.

- year: 연도
- month: 월
- day: 일
- hour: 시
- minute: 분
- second: 초
- microsecond: 마이크로초(micro seconds, 백만분의 일초)

갖고 있는 메서드는 다음과 같다.

* weekday(): 요일 반환 (0:월, 1:화, 2:수, 3:목, 4:금, 5:토, 6:일)
* strftime(): 문자열 반환
* date(): 날짜 정보만 가지는 datetime.date 클래스 객체 반환
* time(): 시간 정보만 가지는 datetime.time 클래스 객체 반환

timedelta를 사용해 날짜 차를 구할 수 있다.

```
import datetime
 
now = datetime.datetime.now()
print(now)      
 
tomorrow = now + datetime.timedelta(days=1)
print(tomorrow) # 하루 후

edt = now - datetime.timedelta(hours=3) # 세시간 전
```

---

## time 모듈

우선 시간을 다루기 전 뭔시간이 뭔시간 인지부터 알아본다.

#### 타임스탬프(TimeStamp)

컴퓨터에서 시간을 측정하는 방법으로 1970년 1월1일 자정
(epoch)이후로 초 단위로 측정한 절대시간 입니다.

#### 협정세계시(UTC, Universal Time Coordinated)

1972부터 시행된 국제 표준시
(세슘 원자의 진동수에 의거한 초의 길이가 기준이 됩니다)

#### 그리니치 평균시(GMT, Greenwich Mean Time)

런던 그리니치 천문대의 자오선상에서의 평균 태양시 (1972년
부터 협정세계시를 사용하지만, 동일한 표현으로 널리 쓰임)

#### 지방표준시(LST, Local Standard Time)

UTC를 기준으로 경도 15도마다 1시간 차이가 발생하는 시간
(한국은 동경 135도를 기준으로 UTC보다 9시간 빠름)

####  일광절약 시간제(DST, Daylight Saving Time)

흔히 서머타임으로 알고 있는 것으로, 에너지 절약을 목적
으로 시간을 한 시간씩 앞당기거나 뒤로 미루는 제도.


[출처](http://devanix.tistory.com/297)

그러나 서로 다른 시간대를 변환하는 데 time모듈을 쓰면 시간대 사이 시간이 바뀌기도 한다.

이왕이면 datetime을 사용해야 하는게 맞다고 한다.

```
import time
now = time.localtime()


print "현재 년: %d" % (now.tm_year)
print "현재 월: %d" % (now.tm_mon)
print "현재 일: %d" % (now.tm_mday)

print

print "현재 시: %d" % (now.tm_hour)         # 24시간제
print "현재 분: %d" % (now.tm_min)
print "현재 초: %d" % (now.tm_sec)

print

print "오늘 요일: %d"      % (now.tm_wday)  # 월요일 = 0
print "올해 몇번째 날: %d" % (now.tm_yday)  # 1월 1일 = 1
print "서머타임 여부: %d"  % (now.tm_isdst) # 서머타임 없으면 0
```

