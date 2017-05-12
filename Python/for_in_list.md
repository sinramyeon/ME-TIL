#### list 안에 for 문을 넣어 보자.

##### list comprehension

`[표현식 for 항목 in 반복가능 객체 if 조건]`

---

예제)


```
a = [1,2,3,4]
result = []
for num in a :
    result.append(num*3)

print(result)
[3,6,9,12]
```

리스트 내포를 이용해서 바꿔보자.

```
result = [num * 3 for num in a]

print(result)
[3,6,9,12]
```

만약 짝수에만 3을 곱하여 담고 싶다면 다음과 같이 "if 조건"을 사용할 수 있다.

```
result = [num * 3 for num in a if num % 2 == 0]
print(result)
[6, 12]
```

구구단을 인쇄해보자.

```
gugu_result = [x*y for x in range(2,10) for y in range(1,10)]
print(result)
[2, 4, 6, 8, 10, 12, 14, 16, 18, 3, 6, 9, 12, 15, 18, 21, 24, 27, 4, 8, 12, 16,
20, 24, 28, 32, 36, 5, 10, 15, 20, 25, 30, 35, 40, 45, 6, 12, 18, 24, 30, 36, 42
, 48, 54, 7, 14, 21, 28, 35, 42, 49, 56, 63, 8, 16, 24, 32, 40, 48, 56, 64, 72,
9, 18, 27, 36, 45, 54, 63, 72, 81]
```