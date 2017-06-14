### python 함수 인자값 정리

---

[참고](https://godoftyping.wordpress.com/2017/05/24/python-%ED%95%A8%EC%88%98-%EC%9D%B8%EC%9E%90%EA%B0%92-%EC%A0%95%EB%A6%AC/)

#### 디폴트 인자

```
def myFunc(a, b=20, c=30) :

#디폴트 인자 뒤에 일반 인자가 위치할 수 없음.

    total = a+b+c
    return total
    
myResult = myFunc(10)
print (myResult)

# 60
```

#### 가변형 인자

```
def yourFunc(*args) :
    total = 0
    for i in args :
        total += i
    return total

yourResult = yourFunc(1, 2, 3)
print(yourResult)

#6
```

#### varargs(키워드 인자)

함수에 인자값 앞에 \*\*를 붙이면 키워드 인자가 됨.
이런 함수를 호출할때는 변수명=값 형태로 호출. 함수에서 값 출력시 사전형으로 출력됨

```
def hisFunc(**args) :
    args['z' = 99
    return args
    
hisResult = hisFunc(a=10, b=20)
print (hisResult)

# {'z'=99, 'b'=20, 'a'=10}
```

#### 복합형 인자

인자 나열 순서 : 일반 > 디폴트 > 가변형 > 키워드

```
def ourFunc(a, b, c=30, *varargs, **kwargs):
    print('a:', a)
    print('b:', b)
    print('c:', c)
    print('varargs:', varargs)
    print('kwargs:', kwargs)
 
ourFunc(10, 20, 'a', 'b', 'c', x=100, y=200)
# a: 10
# b: 20
# c: 'a'
# varargs: ('b', 'c')
# kwargs: {'x': 100, 'y': 200}
```