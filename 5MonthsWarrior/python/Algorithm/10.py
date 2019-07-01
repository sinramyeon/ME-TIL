# 약수 구하기

def solve(n) :
  ans = 0
  
  i = 1
  while i <=n :
    if n%i == 0 :
      print("약수 : {}".format(i))
    i = i +1
 
 
 # 그 수로 나누어 떨어지면 약수~
