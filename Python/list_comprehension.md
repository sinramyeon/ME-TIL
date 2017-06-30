# 파이썬 컴프리헨션

`[row_element for row in matrix for row_element in row]`


## list comprehensions
`[표현식 for 항목 in 순회 가능한 객체 if 조건]`
```
multiples = [i for i in range(30) if i % 3 is 0]
print(multiples)
# output : [0, 3, 6, 9, 12, 15, 18, 21, 24, 27]
```

## dict comprehensions
`[키_표현식 : 값_표현식 for 표현식 in 순회 가능한 객체]`

```
mcase = {'a':10, 'b':34, 'A':7, 'Z':3}

mcase_frequency = { k.lower() : mcase.get(k.lower(), 0) + mcase.get(k.upper(), 0) for k in mcase.keys() }

# mcase_frequency == {'a': 17, 'z': 3, 'b': 34}
```

## set comprehensions
`[표현식 for 표현식 in 순회 가능한 객체]`

```
squared = {x**2 for x in [1, 1, 2]}
print(squared)

# Output: {1, 4}
```

### 예제들

```
a_list = [number for number in range(1,6) if number % 2 ==1]
[1,3,5]
```

```
rows = range(1,4)
cols = range(1,3)

cells = [(row, col) for row in fows for col in cols]

for cell in cells  :
    print(cell)

(1,1)
(1,2)
(2,1)
(2,2)
(3,1)
(3,2)
```

```
word = 'letters'
letter_counts = {letter : word.count(letter) for letter in set(word)}

{'t':2, 'l':1, 'e':2, 'r':1, 's':1}
```

# 튜플은 컴프리헨션이 없다!

대신 제너레이터 객체를 반환한다
```
number_thing = (number for number in range(1,6))

type(number_thing)
>>> <class 'generator'>
```