# 소수 판별기
# 23571113
# 자기보다 작은수까지해서 나누어 떨어지는지 보기

def check_prime(n) :
  i = 2
  while i < n :
    if n % i == 0 :
      break # 두 수가 동일
    i = i+1
  
  if i == n :
    # 자기자신으로밖에 안나눠지면...
    print("소수")
  else :
    print("합성수")
  
