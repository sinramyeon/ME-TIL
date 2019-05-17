# 파이썬 함수 인자

### 위치인자

```
def menu(wine, entree, dessert) :
  return {'wine' : wine, 'entree' : entree, 'dessert': dessert}
```

각 위치 순서대로 함수 인자를 넣는 일반적 방식

### 키워드 인자

```
menu(entree='beef', desset='bagel', wine='bordeaux')
```

위치 혼동을 피하기 위해 입자 매개변수 이름을 지정한다.

### 기본 매개변수값을 지정

```
def menu(wine, entree, dessert='pudding') :
  return{'wine' : wine, 'entree' : entree, 'dessert' : dessert} 
```

값이 **None**이면 기본값을 리턴한다.


### *args

위치인자 변수들을 튜플로 묶는다.

> def print_args(*args) :
>   print(args)
>   print_args(3,2,1, 'wait', '..')
>   3, 2, 1, wait, ... #가 출력

### **kwargs

위치인자 변수들을 딕셔너리로 묶는다.

> def print_kwargs(**kwargs) :
>   print kwargs
>   print_kwargs(wine='merlot', entree='mutton', dessert='macaroon')
> {'dessert' : 'macaroon', 'wine' : 'merlot', 'entree' : 'mutton'}


