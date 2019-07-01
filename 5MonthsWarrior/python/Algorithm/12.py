# 약수의 합 출력하기

def solve(n) :

  ans = 0
  i = 1
  
  while i<= n :
    if n%i == 0 :
      ans = ans+i
    i = i +1
  
  
  return ans
