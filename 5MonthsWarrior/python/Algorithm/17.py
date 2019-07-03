# 배열 채우기

a = [1]*100
def calc_array(n) :
  a[1] = 1
  a[2] = 1
  
  i = 3
  while i <= n :
    a[i] = a[i-1]+a[i-2]+1
    i = i+1
    
    # 이전 두개의 값을 더해서 현재 값으로 저장
