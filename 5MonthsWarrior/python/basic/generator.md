# 제너레이터

파이썬의 시퀀스를 생성하는 객체 ex) range()

> def generator(n) :
>   i = 0
>    while i < n :
>      yield i
>      i += 1

일반 함수와의 차이는 **yield** 구문이다.

> 실행 중 yield를 만나면 해당 함수는 그 상태로 정지되며 반환값을 next()를 호출한데로 전달한다.

```
def generator(n) :
  i = 0
  while i < n :
    yield i
    i+=1

for x in generator(5) :
  print x

0
1
2
3
4
```

list, set, dic는 모두 iterable 하기 때문에 for문에서 유용하게 쓸 수 있다. 그렇지만, 모든 값이 메모리에 있기 때문에 큰 값을 다룰 때는 별로 좋지 않다.

이럴때 제너레이터를 사용하면, yield를 통해 그때그떄 필요한 값만 받아서 쓰기때문에 모든 값을 메모리에 들고 있을 필요가 없다.

![http://nvie.com/img/relationships.png](http://nvie.com/img/relationships.png)

