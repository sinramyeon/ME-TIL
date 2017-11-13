# python 의 getter setter?

---

캡슐화 문제가 있을 때 객체지향 언어인 java에서는 Getter와 Setter를 사용했다. 거진 모든 변수는 public인 파이썬에서는 파이써닉한 방법 @property를 사용한다.

## Property?

> @property 식으로 데코레이터로 쓰임.

데코레이터에 대한 참고는 아래와 같음.

### decorator
#### 함수를 취해서 또 다른 함수를 반환

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

그럼 property는 무엇인가? 이것은 파이썬이 제공하는 기본 함수인데, Instance 의 Attribute 를 할당/삭제할 때 사용한다.

파이썬의 클래스 접근제한자는 *_(protected), __(private)* 이 있고 가장 간단하게 __init__()함수를 떠올려보면 쉽다.

이를 naming convention이라고 하고 __가 붙은 애들은 바로 호출이 안된다.
mangling이라고도 한다.

단적인 예로 `main.__init__()` 은 호출이 불가능하다.

property 는 그럼 뭘까?
이것은 파이썬이 제공하는 기본 함수인데, Instance 의 Attribute 를 할당/삭제할 때 사용한다.

---

## Get/Set 예제

```
class Movie:
    def __init__(self, movie_name):
        self.__movie_name = movie_name
 
    @property
    def movie_name(self): # 이때 메서드 이름은 변수(속성)의 이름과 동일하게 하는 것이 좋습니다. 
        return self.__movie_name
 
    @movie_name.setter
    def movie_name(self, new_movie_name): # 이때 메서드 이름은 변수(속성)의 이름과 동일하게 하는 것이 좋습니다. 
        """ 영화를 변경하는 setter 메서드"""
        self.__movie_name = new_movie_name
        print("============ setter를 통해 영화를 변경합니다============")
        print('변경 후 영화이름 : {} '.format(self.movie_name))
```

같은 movie_name()이지만 장식자가 @property냐 @movie_name.setter냐에 따라 결과가 달라집니다.

```
movie = Movie('총알 탄 사나이')
print(movie.movie_name)
movie.movie_name = '히든 피겨스' # 객체의 속성값에 직접 접근하는듯이 사용하지만 
                               # 실제로는 메서드 호출을 통해 변수에 접근한다. 
print(movie_movie_name)
 
>>> 출력결과

> 총알 탄 사나이
> ============ setter를 통해 영화를 변경합니다============
> 변경 후 영화이름 : 히든 피겨스
> 히든 피겨스
```

@property 없이는 `movie = Movie('새영화'), movie = Movie("또다른영화")` 식으로 변경해야 한다.

get, set 함수를 사용할수도 있다.

```
class Test:

def __init__(self):
self.color = "red"

def set_color(self,clr):
self.color = clr

def get_color(self):
return self.color

if __name__ == '__main__':

t = Test()
t.set_color("blue")

print(t.get_color())

```

이걸 @property 로 적용시켜본다고 치면 아래와 같다.

```
class Test:

    def __init__(self):
        self.__color = "red"

    @property
    def color(self):
        return self.__color

    @color.setter
    def color(self,clr):
        self.__color = clr

if __name__ == '__main__':

    t = Test()
    t.color = "blue"

    print(t.color)

```

그렇다면 왜 @property를 쓰는가?


> 1. 변수를 변경 할 때 어떠한 제한을 두고 싶어서
> 2. get,set 함수를 만들지 않고 더 간단하게 접근하게 하기 위해서
> 3. 하위호환성에 도움이 됨. 

---

## 하위 호완성

```
class Duck() :
    def __init__(self, input_name) :
        self.hidden_name = input_name
    @property
    def name(self) :
        print("getter 안")
        return self.hidden_name

    @name.setter
    def name(self, input_name) :
        print("setter 안")
        self.hidden_name = input_name
```

`foul = Duck("오리")`로 선언 후 `foul.name` 등으로 접근하거나 `foul.name = "리오"`등으로 값을 변경할 수는 있지만 `get_name()`또는 `set_name()`메서드는 보이지 않는다.

```
class Circle() :
    def __init__(self, radius) :
        self.radius = radius
    @property
    def diameter(self) :
        reutrn 2 * self.radius
```

`c.diameter`로는 접근을 할 수 없는 read-only 속성이다.
직접 속성에 접근하지 않고 property를 사용해서 접근하자.

```
c = Circle(5)
c.radius
>> 5

c.diameter
>> 10
```

property를 통해 접근하면 속성 정의를 바꾸려고 해도 모든 호출자를 수정할 필요가 없어진다.

---

## 정리

@메서드이름.setter
@메서드이름.getter
@property


test = Test() 식으로 객체를 먼저 만들고
test.color 에 접근하면 된다.





