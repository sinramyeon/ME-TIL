# 리스트와 튜플

> Tupe : 불변
> List > 변경 가능

---

### 리스트

```
empty_list = []
empty_list2 = list()
```

> list()는 타 데이터 타입을 리스트 타입으로 변환
> ex) tuple_a = ('a', 'b')
> list(tuple_a)

> split()는 문자열을 나눠 리스트로 반환
> ex) birthday = '22/12/1992'
> birthday_split('/')

> join()는 문자열으로 리스트를 변환
> ex) birthday = ['1992', '12', '22']
> '-',.join(birthday)
- list.append()
	- ex) list.append(2)
- list.extend()
	- ex) list.extend(otherlist)
- list.insert()
	- ex) list.insert(3, 'insert in index 3')
- list.del()
	- ex) list.del(3, 'del in index 3')
- list.remove()
	- ex) list.remove('delete this')
- list.count()
	- ex) list.count('Bob')
- list.pop() 이나 list.index(), list.sort(), list.len(), list.copy() 등

- :in 으로 존재여부 확인
> if 'Bob' in lists


---

### 튜플

> empty_tuple = ()
> one_marx = ('Groucho')
> tuple(some_list)

### 왜 튜플을 쓰는지?

- 튜플은 더 적은 공간을 사용한다
- 불변의 목록을 원할 때
- 튜플을 딕셔너리 키로 응용이 가능
- 네임드 튜플 때문에
- 함수 인자들이 튜플로 전달되기 때문에
