# 파이썬 함수 공부

## 일등객체 : 함수

**모든것은 객체**

파이썬에서 함수는 first-class-citizen 이다! 함수를 변수에 할당할 수 있고, 다른 함수에서 함수를 인자로 쓸 수 있고 함수에서 함수도 반환할 수 있다.

예제
`def run_something_with_args(func, arg1, args2 :
  func(arg1, arg2)`

### 함수 내 함수 정의 예제

>>> def outer(a,b) :
>>>> def inner(c, d ) :
>>>>>return c+d
	>	return inner(a,b)

>outer(4,7)
>11

내부함수는 루프나 코드 중복을 피하기 위해 함수 내 작업을 한번이상 수행할 시 유용히 이용이 가능하다.

### 클로저!

> 프로그래밍 언어에서의 클로저란 퍼스트클래스 함수를 지원하는 언어의 네임 바인딩 기술이다.
> 
클로저로 쓸 수 잇는 내부 함수
클로저는 다른 함수에 의해 동적으로 생성
바깥 함수로부터 생성된 변수값을 변경하고 저장할 수 있음

```
def kinghts(saying):
  def inner() :
    return 'We are say : %s' & saying
  return inner
```

inner()함수는 knights()함수가 전달받은 saying 을 이용한다. retun inner에서는 호출되지 않은 inner함수의 복사본을 반환한다. 이것이 외부 함수에 의대 동적으로 생성되고 그 함수의 변수값을 알고 있는 함수인 클로저이다.(먼소리임?)

예제 2

```
def outer_func() :
  messgae = 'Hi'
  def inner_func() :
    print message
  return inner_func()
```

1. outer func 실행
2. message 변수 할당
3. inner func 정의
4. message 변수를 참조해서 출력

결과
> Hi
> 
```
def outer_func() :
  messgae = 'Hi'
  def inner_func() :
    print message
  return inner_func
```

1. outer func 실행
2. message 변수 할당
3. inner func 정의
4. 함수 오브젝트를 리턴

결과
>

```
def outer_func():  #1
 message = 'Hi'  #3  
 def inner_func():  #4  
   print message #6 
 return inner_func #5

my_func = outer_func() #2 <-- 리턴값인 inner_func를 변수에 할당합니다.
print my_func
```

결과
><function inner_func at -x2438423>

my_func	에 inner_func가 할당된것을 확인

```
def outer_func():  #1
 message = 'Hi'  #3  
 def inner_func():  #4  
   print message #6 
 return inner_func #5

my_func = outer_func() #2 <-- 리턴값인 inner_func를 변수에 할당합니다.

my_func()
```

결과
>Hi

## 익명함수 lambda()

단일문으로 표현되는 익명함수 lambda()에 대해 알아본다. 람다는 많은 작은 함수를 정의히ㅏ고 이들을 호출해서 얻은 모든 결과값을 저장하는데에 유용하다.

```
stiars = ['thud', 'meow', 'thud', 'hiss']
def enliven(word) :
  return word.capitalize() + "!"
def edit_story(words, func) :
  for word in words :
    print(func(word))


edit_story(stairs, envlien)
Thud!
Meow!
Thud!
Hiss!

# 람다함수 적용
edit_story(stair, lambda word : word, capitalize() +"!")

```
