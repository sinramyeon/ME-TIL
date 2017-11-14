# python collections

python 내 많은 모듈 중 collections모듈을 살펴보자.

[참고](https://www.slideshare.net/dahlmoon/collections-20160313)
[공식문서](https://docs.python.org/3/library/collections.html#module-collections)

주요 요소 | 설명 | released
----------|------|-----------
namedtuple() | tuple타입 subclass를 만들어주는 함수 | 2.6
OrderedDict | 순서가 있는 dict | 2.7
Counter | hash가능한 객체를 카운트하는 dict | 2.7
defaultdict | dict subclass that calls a factory function to supply missing values | 2.5
deque | list-like container with fast apends and pops on eiter end | 2.4

---

## namedtuple

> Named tuples assign meaning to each position in a tuple and allow for more readable, self-documenting code. They can be used wherever regular tuples are used, and they add the ability to access fields by name instead of position index.

튜플 : 인덱스 기준, 네임드튜플 : 키 기준
`import collections`로 임포트해서 쓸 수 있음 `collections.namedtuple`

1. tuple 의 경우 (숫자 인덱스가 있음)

```
bob = ('Bob', 30, 'male')
print ('Representation:', bob)

jane = ('Jane', 29, 'female')
print ('\nField by index:', jane[0])

print ('\nFields by index:')
for p in [ bob, jane ]:
    print ('%s is a %d year old %s' % p)
```

2. namedtuple의 경우 (키가 있음)

```
import collections

Person = collections.namedtuple('Person', 'name age gender')

print ('Type of Person:', type(Person))

bob = Person(name='Bob', age=30, gender='male')
print ('\nRepresentation:', bob)

jane = Person(name='Jane', age=29, gender='female')
print ('\nField by name:', jane.name)

print ('\nFields by index:')
for p in [ bob, jane ]:
    print ('%s is a %d year old %s' % p)
```


키, 밸류형태로 값을 전송해야하며 아닐시 ValueError가 발생합니다.

```
import collections

try:
    collections.namedtuple('Person', 'name class age gender')
except ValueError, err:
    print (err)

try:
    collections.namedtuple('Person', 'name age gender age')
except ValueError, err:
    print (err)
```


```
$ python collections_namedtuple_bad_fields.py

Type names and field names cannot be a keyword: 'class'
Encountered duplicate field name: 'age'
```

rename 옵션을 True로 줄 시 `rename=True` 원래 있던 중복 필드명 등을 이름을 바꿔버림.

```
import collections

with_class = collections.namedtuple('Person', 'name class age gender', rename=True)
print with_class._fields

two_ages = collections.namedtuple('Person', 'name age gender age', rename=True)
print two_ages._fields
```

dict에서도 namedtuple을 만들 수 있으나 두 개는 별개다.

```
parts = {'bill','orange', 'tail':'long'}
duck = Duck(**parts)
duck2
>> Duck(bil='orange', tail='long')
```

네임드 튜플은 불변하므로 필드를 바꿔 또 다른 튜플을 반환하는 식으로 운용한다.

`duck3 = duck2._replace(tail="magnificent', bill='crushing')`

---

## OrderedDict

> Return an instance of a dict subclass, supporting the usual dict methods. An OrderedDict is a dict that remembers the order that keys were first inserted. If a new entry overwrites an existing entry, the original insertion position is left unchanged. Deleting an entry and reinserting it will move it to the end.

dict인데 순서가 있는 dict. linked list로 내부에 구성되어 각 순서를 유지한다.

dict와 비교해보자.

```
import collections

print('Regular dictionary:')
d = {}
d['a'] = 'A'
d['b'] = 'B'
d['c'] = 'C'

for k, v in d.items():
    print(k, v)

print('\nOrderedDict:')
d = collections.OrderedDict()
d['a'] = 'A'
d['b'] = 'B'
d['c'] = 'C'

for k, v in d.items():
    print(k, v)
```

```
$ python3 collections_ordereddict_iter.py

Regular dictionary:
c C
b B
a A

OrderedDict:
a A
b B
c C
```

순서가 달라지면 동일하지 않은 객체가 된다.

popitem()메서드나 move_to_end(key, last=True) 메서드가 있다.

```
>>> d = OrderedDict.fromkeys('abcde')
>>> d.move_to_end('b')
>>> ''.join(d.keys())
'acdeb'
>>> d.move_to_end('b', last=False)
>>> ''.join(d.keys())
'bacde'
```
---

## Counter

> A Counter is a dict subclass for counting hashable objects. It is an unordered collection where elements are stored as dictionary keys and their counts are stored as dictionary values. Counts are allowed to be any integer value including zero or negative counts. The Counter class is similar to bags or multisets in other languages.

개수를 세는데 특화된 클래스.
most_common()을 통해 가장 자주 나타난 내역을 집계할 수도 있다.

```
import collections
 
myList = ['a', 'b', 'c', 'c', 'a', 'a']
myCounter = collections.Counter(myList)
print('myCounter:', myCounter)
# myCounter: Counter({'a': 3, 'c': 2, 'b': 1})
 
print("myCounter['a']:", myCounter['a'])
# myCounter['a']: 3
 
yourList = ['a', 'd', 'c', 'a', 'b']
yourCounter = collections.Counter(yourList)
print('yourCounter:', yourCounter)
# yourCounter: Counter({'a': 2, 'd': 1, 'b': 1, 'c': 1})
 
ourCounter = myCounter + yourCounter
print('ourCounter:', ourCounter)
# ourCounter: Counter({'a': 5, 'c': 3, 'b': 2, 'd': 1})
 
print('ourCounter.most_common(3):', ourCounter.most_common(3))
# ourCounter.most_common(3): [('a', 5), ('c', 3), ('b', 2)]

```

없는 내역을 집계할시 KeyError대신 0을 반환한다.

```
>>> c = Counter(['eggs', 'ham'])
>>> c['bacon']                              # count of a missing element is zero
0
```

---

## defaultdict

> Returns a new dictionary-like object. defaultdict is a subclass of the built-in dict class. It overrides one method and adds one writable instance variable. The remaining functionality is the same as for the dict class and is not documented here.

dict스러운 객체를 반환해 준다. dict인것 같지만 키값이 없어도 에러를 출력하지 않고, dict에 기본값을 정의해 준다.

defaultdict 없이는 키가 존재하는지 확인해주고 기본값을 정해주어야만 한다.

```
# Load the OSX built-in dictionary
with open('/usr/share/dict/words', 'r') as f:
    words = f.read().splitlines()


d = {}
for word in words:
    firstletter = word[0].lower()
    if firstletter not in d:
        d[firstletter] = 0
    d[firstletter] += 1
```

대신 defaultdict를 사용하면 아래와 같이 적용할 수 있다.

```
from collections import defaultdict
d = defaultdict(int)
```

dict.setdefault()보다 빠르다.

* dict.setdefault()란?

> If key is in the dictionary, return its value. If not, insert key with a value of default and return default. default defaults to None.


---

## deque

> Returns a new deque object initialized left-to-right (using append()) with data from iterable. If iterable is not specified, the new deque is empty.

Deques ar

자료구조에서 배운 큐(queue)가 양방향으로 작동할 때 이를 deque라고 한다.

```
# -*- coding:utf-8 -*-
import collections
 
d = collections.deque([10, 20, 30, 40, 50])
print('Deque: ', d)
# 'Deque: ', deque([10, 20, 30, 40, 50])
 
# 오른쪽에 추가
d.append(60)
print('Deque: ', d)
# 'Deque: ', deque([10, 20, 30, 40, 50, 60])
 
# 왼쪽에 추가
d.appendleft(0)
print('Deque: ', d)
# 'Deque: ', deque([0, 10, 20, 30, 40, 50, 60])
 
# 입력값을 순환하면서 오른쪽에 추가(append)
d.extend([70, 80])
print('Deque: ', d)
# 'Deque: ', deque([0, 10, 20, 30, 40, 50, 60, 70, 80])
 
# 입력값을 순환하면서 왼쪽에 추가(appendleft)
d.extendleft([-10, -20, -30])
print('Deque: ', d)
# 'Deque: ', deque([-30, -20, -10, 0, 10, 20, 30, 40, 50, 60, 70, 80])
 
# 값 삭제
d.remove(0)
print('Deque: ', d)
# 'Deque: ', deque([-30, -20, -10, 10, 20, 30, 40, 50, 60, 70, 80])
 
# 오른쪽의 끝값 가져오면서 deque에서 제거
maxValue = d.pop()
print('maxValue:', maxValue)
# maxValue: 80
print('Deque: ', d)
# Deque:  deque([-30, -20, -10, 10, 20, 30, 40, 50, 60, 70])
 
# 왼쪽의 끝값 가져오면서 deque에서 제거
minValue = d.popleft()
print('minValue:', minValue)
# minValue: -30
print('Deque: ', d)
# Deque:  deque([-20, -10, 10, 20, 30, 40, 50, 60, 70])
 
# 값 회전(rotating)
d = collections.deque(range(5))
print('Deque: ', d)
# Deque:  deque([0, 1, 2, 3, 4])
 
d.rotate(1)
print('Deque.rotate(1): ', d)
# Deque.rotate(1):  deque([4, 0, 1, 2, 3])
 
d.rotate(1)
print('Deque.rotate(1): ', d)
# Deque.rotate(1):  deque([3, 4, 0, 1, 2])
 
d.rotate(-1)
print('Deque.rotate(-1): ', d)
# Deque.rotate(-1):  deque([4, 0, 1, 2, 3])
 
d.rotate(-1)
print('Deque.rotate(-1): ', d)
# Deque.rotate(-1):  deque([0, 1, 2, 3, 4])

```

rotate()함수로 하나씩 앞으로 밀거나 당길 수 있는 것을 확인 가능하다.

활용하면 아래와 같다.

```
def tail(filename, n=10):
    'Return the last n lines of a file'
    with open(filename) as f:
        return deque(f, n)
```

---