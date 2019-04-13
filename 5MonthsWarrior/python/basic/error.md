# 파이썬 에러 만들기

사용자 정의 에러

```
class UppercaseException(Exceptiono) :
   pass

words = = ['eeenie', 'meenie', 'miny', 'MO']

for word in words :
  if word.issuper() :
    raise UppercaseException(word)
```

try/except 문

```
try: 
  ... 
except [발생 오류[as 오류 메시지 변수]]: 
  ...
```

물론 finally 문도 사용 가능함.



