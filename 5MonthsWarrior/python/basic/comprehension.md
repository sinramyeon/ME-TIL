# 컴프리헨션

### 리스트 컴프리헨션

> [표현식 for 항목 in 순회 가능한 객체]
`number_list = [num for num in range(1,6)]`

> [표현식 for 항목 in 순회 가능한 객체 if 조건]
> `a_list = [num for num in range(1,6) if num%2 == 1]`


### 딕셔너리 컴프리헨션

> {키 표현식 : 값 표현식 for 표현식 in 순회가능한 객체}
`word='letters'`
`letter_counts = {letter : word.count(letter) for letter in word}`

### 셋 컴프리헨션
> {표현식 for 표현식 in 순회가능한 객체}
리스트, 딕셔너리와 비슷한 모습

### 제너레이터 컴프리헨션

튜플은 컴프리헨션이 없음
> number_things = (num for num in range(1,6)) 은 제너레이터 컴프리헨션이고 class generator를 반환함
