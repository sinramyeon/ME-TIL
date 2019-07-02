# 최소공배수 구하기

def calc_LCM(n) :
  a = 2
  b = 3
  c = 5
  # 2, 3, 5 최소공배수 구하기
  
  i = 2
  
  while i< n :
    if (i%a==0) and(i%b ==0) and (i%c ==0) :
      break
    i =i +1
    
  print(i)
  
  # 3가지로 다 나눠지는 작은 수.
