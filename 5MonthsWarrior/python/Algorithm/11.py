# 에라토스테네스의 체 소수 고르기
# n의 배수를 모두 제거하는 방법
True = 1
False = 0

def calc_prime(n) :
  check = [True]*1000
  cnt = 0 
  i = 2
  k = 1
  
  while i < n :
    if check[i] = True :
      print(i)
      j = 1
      while j < n :
        check[j] = False
        k = k+1
        j = i*k # 배수를 다 빼는건지??
      k = 1
    i = i+1
    
    
    # 이해를 잘 모샇겠음...
    # 소수를 걸러내는?
