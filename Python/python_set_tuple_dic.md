## PYTHON 기본 자료형

---

#### 리스트

##### 값의 나열, 여러 종류를 담아낼 수 있다.
##### []

###### append(), insert(), +연산자
###### del, remove, sort, reverse, index, pop, count, extend

```
>>> a = [ ]
>>> b = [1, 2, 3]
>>> c = ['Life', 'is', 'too', 'short']
>>> d = [1, 2, 'Life', 'is']
>>> e = [1, 2, ['Life', 'is']]
```

인덱싱

```
>>> a = [1, 2, 3, ['a', 'b', 'c']]

>>> a[0]
1
>>> a[-1]
['a', 'b', 'c']
>>> a[3]
['a', 'b', 'c']
```

슬라이싱

```
>>> a = [1, 2, 3, 4, 5]
>>> a[0:2]
[1, 2]
```

---

#### 튜플

##### 읽기 전용, 속도가 빠르다.
##### ()

```
>>> nlist = [1,2,3]
>>> ntuple = (4,5)
```
---

#### 딕셔너리

##### 키와 밸류로 구성
##### {a : b}

```
>>> dic = {'name':'pey', 'phone':'0119993323', 'birth': '1118'}
```

Key 리스트 만들기

```
>>> a = {'name': 'pey', 'phone': '0119993323', 'birth': '1118'}
>>> a.keys()
dict_keys(['name', 'phone', 'birth'])
```

> python 2.7까지는 keys()함수에서 리스트를 리턴했으나, 3부터는 dict_keys라는 객체를 리턴해 줌

Values 리스트 만들기

```
>>> a.values()
dict_values(['pey', '0119993323', '1118'])
```

Key, Value 쌍 얻기

```
>>> a.items()
dict_items([('name', 'pey'), ('phone', '0119993323'), ('birth', '1118')])
```

Key 로 밸류 얻기 : .get("key")
해당 Key의 존재 여부 조사 : "key" in dict (boolean 반환)

---

#### 세트(집합)

##### {}
##### 중복을 허용하지 않음. 순서가 없음

```
>>> num1 = {1,2,3}
>>> num2 = {3,4,5}
>>> print (num1, num2)
>>> {1,2,3} {3,4,5}
>>> num1.union(num2) #합집합
>>> {1,2,3,4,5}
>>> num1.intersection(num2) #교집합
>>> {3}
>>> num1 - num2 #차집합
>>> {1,2}
```

Unordered

```
>>> s2 = set("Hello")
>>> s2
{'e', 'l', 'o', 'H'}
```

값 1개 추가 .add()
값 여러개 추가 .update([])
특정 값 제거 .remove()