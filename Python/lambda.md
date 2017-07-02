# 람 다 함 수

## 람다식, 람다 표현식, 익명 함수

```
def f(x,y) :
    return x+y

f = lambda x,y : x+y
```

`lambda 인자1, 인자2 .... : 표현식`

### 람다는 왜 써?

익명 함수 > 즉 일회용으로 쓰고 버리고 하고 싶을 때, 그냥 함수를 쓰고 싶은데 이름을 안정해주고 inline으로 사용.


예제)
```
>>> def pos(x):
…..    return x>0
>>> list(filter(pos, range(-5,6)))
[1,2,3,4,5]
```

를 이렇게 줄일 수 있다.

`list(filter(lambda x : x>0, range(-5, 6)))`

