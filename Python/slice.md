# Python Slice

python에서 slice 기법을 잘 활용해 보자. 아주 편리하다. IndexOf()와 len()의 지옥에서 탈출할 수 있다.

slice(시작점, 중단점, 범위)는 정말 편하다. 빌트인 펑션 중 하나고, class이다. `help(slice)` 명령어를 입력하면 사용법을 확인할 수도 있다.(파이썬에서 갑자기 뭔가의 사용법이 기억이 안나면 help()함수가 은근 유용하다)


```
> help(slice)

class slice(object)
 |  slice(stop)
 |  slice(start, stop[, step])
 |
 |  Create a slice object.  This is used for extended slicing (e.g. a[0:10:2]).
```

사용 예제는 아래와 같다.

```
>>> s = slice(100, 200, 300)
>>> s.start
100
>>> s.stop
200
>>> s.step
300
```

특히 문자열 등을 다룰 때 아주 편하게 사용할 수 있다.

```
word = "인터파크메롱"
part = word[:4]
print(part)

> 인터파크
```

홀수, 짝수번 추출에도 유용히 쓸 수 있다.

```
values = "AaBbCcDdEe"

evens = values[::2]
print(evens)

> ABCDE
```

역추출에도 편하다.

```
a[-1]    # 마지막 아이템
a[-2:]   # 마지막 두개 아이템
a[:-2]   # 마지막 두개만 빼고
a[::-1]  # 거꾸로 복사
```

문법을 정리해보면 다음과 같다.

```
a[start:end] # start부터 end-1 까지
a[start:]    # start부분부터
a[:end]      # end-1까지만
a[:]         # [다 복사](http://docs.python.org/2/library/copy.html)
a[start : end : step] # step씩 띄어서 start부터 end까지
```

```
 +---+---+---+---+---+---+
 | P | y | t | h | o | n |
 +---+---+---+---+---+---+
 0   1   2   3   4   5   6
-6  -5  -4  -3  -2  -1
```

---

## __getitem__()

Python에서, 특정 object에 [] 연산자를 사용하면 내부적으론 __getitem__이라는 메소드가 실행된다. 예를 들어서 arr[1] 이라고 했다면 arr.__getitem__(1) 이 작동한다.