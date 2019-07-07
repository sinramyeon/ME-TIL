# 재귀 팩토리얼 출력

def fact(n) :
  if n==0 :
    return 1
  else :
    return n*fact(n-1)
