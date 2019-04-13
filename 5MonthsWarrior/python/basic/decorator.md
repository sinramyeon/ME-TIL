# 데코레이터

하나의 함수를 취해 또 다른 함수를 반환

 > 퍼스트 클래스 함수와 클로저부터 이해하기...

데코레이터 > 장식자 구문

```
def outer_function(msg) :
  def inner_function() :
    print msg
  return inner_function

hi = outer_function('Hi')

hi_func()
Hi
```

클로저가 위랑 비슷하면 데코레이터의 예제는 아래를 생각하면 좋다.

```
def decorator_function(original_function):  #1 
 def wrapper_function():  #5 
   return original_function() #7  
 return wrapper_function #6  

def display():  #2  
  print  'display 함수가 실행됐습니다.'  #8 


decorated_display = decorator_function(display) #3 

decorated_display() #4
```


함수 래퍼라고 생각하면 편하다. display라는 함수를 wrapper와 decorator 안에 감싸고 있다.

```
# -*- coding: utf-8 -*-  
def decorator_function(original_function): 
 def wrapper_function():
   print  '{} 함수가 호출되기전 입니다.'.format(original_function.__name__)
   return original_function()
 return wrapper_function 


def display_1():  
  print  'display_1 함수가 실행됐습니다.'  

def display_2(): 
 print  'display_2 함수가 실행됐습니다.' 

display_1 = decorator_function(display_1) #1 


display_2 = decorator_function(display_2) #2 

display_1()
print
display_2()
```

이렇게 함수를 코드수정 없이 래퍼함수를 이용하여 여러가지 기능을 추가하기 위해 사용한다.
그렇지만 보통 데코레이터는 이렇게 풀어쓰지 않고 @를 붙여 사용한다.

```

def decorator_function(original_function):
  def wrapper_function(*args, **kwargs):  #1 
   print  '{} 함수가 호출되기전 입니다.'.format(original_function.__name__) 
   return original_function(*args, **kwargs) #2  
return wrapper_function


@decorator_function
def display() :
  print  'display 함수가 실행됐습니다.'

@decorator_function
def display_info(name, age) : 
 print  'display_info({}, {}) 함수가 실행됐습니다.'.format(name, age) d

display()
print
display_info('hero', 28)

display 함수가 호출되기전 입니다.
display 함수가 실행됐습니다.

display_info 함수가 호출되기전 입니다.
display_info(John, 25) 함수가 실행됐습니다.
```

---

도저히 먼소린지 모르겠지만 대략

### 일급 객체
 - 함수 내에 함수를 정의할 수 있음
 - 함수를 인자로 전달할 수 있음
### 클로저
 - 내부 함수가 외부 함수 인자를 기억하고 있음
 - *args, **kwargs에서 args는 인자 순서가 있지만 kwargs는 없음
 ### 데코레이터
-  함수를 취해 또 다른 함수를 반환
