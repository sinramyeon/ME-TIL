# decorator
[참고자료](http://schoolofweb.net/blog/posts/%ED%8C%8C%EC%9D%B4%EC%8D%AC-%EB%8D%B0%EC%BD%94%EB%A0%88%EC%9D%B4%ED%84%B0-decorator/)

## 함수를 취해서 또 다른 함수를 반환

```
def document_it(func) :
    def new_func(*args, **kwargs) : 
    #*args 파라미터를 몇 개 받을지 모르겠다! 튜플 형태로 전달됨.
    #**kwargs 파라미터 명도 같이 보내자. 딕셔너리 형태로 전달됨.
        
        print("Running function : ", func.__name__)
        print("Positional arguments : ", args)
        print("Keyword argumetns : ", kwargs)

        printresult = func(*args, **kwargs)

        print(printresult)

        return result

    return new_func

```

위와 같은 데커레이터 코드가 있다고 하면, 아래와 같이 씁니다.
```
def add_ints(a,b) :
    return a+b

add_ints(3,5)
8

cooler_add_int = document_it(add_ints) # 데커레이터를 수동으로 할당했다.
cooler_add_int(3,5)

Running function : add+ints
Positional arguments : (3,5)
keyword arguments : {}
8
```

데커레이터를 자동으로 할당할 때는

```
@document_it
def add_ints(a,b) :
    retun a+b
```

위와 같이 @를 써 보세요.

예제
```
# -*- coding: utf-8 -*-
def decorator_function(original_function):
    def wrapper_function():
        print '{} 함수가 호출되기전 입니다.'.format(original_function.__name__)
        return original_function()
    return wrapper_function


@decorator_function  #1
def display_1():
    print 'display_1 함수가 실행됐습니다.'


@decorator_function  #2
def display_2():
    print 'display_2 함수가 실행됐습니다.'

# display_1 = decorator_function(display_1)  #3
# display_2 = decorator_function(display_2)  #4

display_1()
print
display_2()
```
```
$ python decorator.py
display_1 함수가 호출되기전 입니다.
display_1 함수가 실행됐습니다.

display_2 함수가 호출되기전 입니다.
display_2 함수가 실행됐습니다.
```