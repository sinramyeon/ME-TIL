# 앞에서읽어도 반대로읽어도 같은 단어 찾기
# 큐, 스택 특성을 이용하자.
# 큐는 FIFO 스택은 LIFO
# 즉 큐돌려도 스택돌려도 같으면 역삼역 같은 단어다.


def palindrome(s) :
  que = []
  stack = []
  
  for x in s :
  
    if s.isalpha() : # 문자가 알파벳이면
      que.append(x.lower())
      stack.append(x.lower())
    
  while que :
    if que.pop(0) != stack.pop() # 두개가 다를시 아님
      return False
  return True
