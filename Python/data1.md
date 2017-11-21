# Python with Data1

---

## FileIO


### 파일 열기

`fileobj = open(파일이름, mode)`

* mode?

모드 | 무슨 기능?
r    | 파일 읽기
w    | 파일 쓰기(있으면 덮어쓰고 없으면 생성)
x    | 파일을 쓰지만 파일이 없을때만 씀
a    | 파일 추가하기(있으면 끝부터 쓰기)

t(無)| 텍스트 타입
a    | 이진(바이너리)타입

Ex) test = open('hey.txt', 'wt')
hey.txt파일을 쓰기모드, 텍스트 타입으로 연다.

### 파일 쓰기

`write()`

```
fout = open('test.txt', 'wt')
fout.write("test")
fout.close()
```

#### 텍스트파일 읽기

`read()`

```
fin = open('test', 'rt')
poem = fin.read()
fin.close()
```

한번에 얼마만큼 읽을지 read(3) 이런식으로 변수로 넘기면 된다.

`readline()`은 한 줄씩 라인 단위로 읽을 수 있다.

```
poem = ''
fin = open('test', 'rt')
while True :
    line = fin.readline()
    if not line :
        break
    poem += line

fin.close()
```

이터레이터로 텍스트 파일을 읽는편이 더 파이써닉하다.

```
for line in fin :
    poem += line
fin.close()
```

한번에 모든 라인을 읽으려면 `readlines()` 를 호출한다.
대신 한 라인으로 된 문자열 리스트를 반환해 준다.

### 자동 파일 닫기

콘텍스트 매니저에서 한 줄 실행 후 자동으로 파일을 닫아 준다.

```
with open ('test', 'wt') as four :
    fout.write(poem)
```

---

## CSV

CSV 파일은 어떤 데이터를 ","를 통해 나누어 놓은 것이다.


연도	제조사	모델	설명	가격
1997	Ford	E350	ac, abs, moon	3000.00
1999	Chevy	Venture "Extended Edition"		4900.00
1999	Chevy	Venture "Extended Edition, Very Large"		5000.00
1996	Jeep	Grand Cherokee	MUST SELL!
air, moon roof, loaded	4799.00

위의 데이터 표는 다음과 같이 CSV 형식으로 표현할 수 있다:

연도,제조서,모델,설명,가격
1997,Ford,E350,"ac, abs, moon",3000.00
1999,Chevy,"Venture ""Extended Edition""","",4900.00
1999,Chevy,"Venture ""Extended Edition, Very Large""",,5000.00
1996,Jeep,Grand Cherokee,"MUST SELL!
air, moon roof, loaded",4799.00

파이썬에선 이미 [모듈](https://docs.python.org/3/library/csv.html)이 있어 가져다 행복하게 사용할 수 있다.



```
import csv
villains = [
['doctor', 'who'],
['test', 'test']
...
]

#콘텍스트 매니저로 자동으로 열고 닫을 것
with open('villains', 'wt') as fout :
    csvout = csv.writer(four)
    csvout.writerows(viallains)
```

은 csv 파일을 생성하고, 읽으려면 다음과 같이 한다.

```
with open('villains', 'rt') as fin :
    cin = csv.reader(fin)
    villains = [row for row in cin]
```

행별로 입력해서 배열을 얻을 수 있다.
만약 dict로 갖고오고 싶다면, 아래와 같이 행한다.

```
with open('villains', 'rt') as fin :
    cin = csv.DictReader(fin)
    villains = [row for row in cin]
```

마찬가지로 생각대로 DictWriter()도 있으니 필요대로 활용한다.

---

## XML

xml은 웹에서 쓰는 자료의 보편 표준 양식이다.
아래와 같은 양식을 띄고 있다.

```
<? xml version="1.0"?>
<컴퓨터언어>
<C언어> C       </C언어>
<C언어> C++    </C언어>
<C언어> C#     </C언어>
<JAVA> java     </JAVA>
<JAVA> android </JAVA>
</컴퓨터언어>
```

가장 잘 알려진 마크업 형식이고, 사용자 지정 태그 안에 값들을 넣는다.
태그 안에 태그를 중첩 가능하다. 보통 데이터 피드나 메시지 전송, RSS등에 쓰인다.

파이썬에서는 `ElementTree` 모듈을 받아쓰거나 표준으로는 `xml.dom`이나 `xml.sax`가 있다.

ElementTree 예를 들어보자.

```
from xml.etree.ElementTree import Element, dump

note = Element("note")
to = Element("to")
to.text = "Tove"

note.append(to)
dump(note)
```

를 실행하면 `<note><to>Tove</to></note>` 와 같이 된다.

파징은 더 간편하다.

```
from xml.etree.ElementTree import parse
tree = parse("note.xml")
note = tree.getroot()
```

이러헥 파징한 후, 어트리뷰트 값은 .get()이나 .keys(), .items()를 사용한다.

.find()와 .getiterator(), .getchildren()등을 사용하며 자세한 내용은 [여기](http://effbot.org/zone/element.htm) 에서 확인하면 좋다.

---

## JSON

파이썬 json 모듈 `json` 하나만 있으면 된다.
[참고](https://www.slideshare.net/dahlmoon/json-20160301)

```
import json
 
# 테스트용 Python Dictionary
customer = {
    'id': 152352,
    'name': '강진수',
    'history': [
        {'date': '2015-03-11', 'item': 'iPhone'},
        {'date': '2016-02-23', 'item': 'Monitor'},
    ]
}
 
# JSON 인코딩
jsonString = json.dumps(customer)
 
# 문자열 출력
print(jsonString)
print(type(jsonString))   # class str

# JSON 인덴트 넣기
jsonString = json.dumps(customer, indent=4)
print(jsonString)

# JSON 디코딩하기

dict = json.loads(customer)
 
# Dictionary 데이타 체크
print(dict['name'])
for h in dict['history']:
    print(h['date'], h['item'])
```

---

## YAML

 e-mail 양식에서 개념을 얻어 만들어진 '사람이 쉽게 읽을 수 있는' 데이터 직렬화 양식. 아래와 같이 생겼다.

 ```
 --- # Favorite movies, block format
- Casablanca
- Spellbound
- Notorious
--- # Shopping list, inline format
[milk, bread, eggs]
```

아직 표준 라이브러리가 없어서 [yaml](https://github.com/yaml) 을 설치해야 한다고 한다.

```
import yaml
with open("어쩌구.yaml", 'rt') as fin :
    text = fin.read()

data = yaml.load(text)
data['키들']
```

dict를 쓰는 것처럼 쓰면 된다.

---
