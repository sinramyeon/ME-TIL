#배열채우기2

a = [0] * 101

def calc_array(n) :
  a[1] = 1
  i = 2
  
  while i<= n :
    a[i] = a[i//2]+1
    i = i+1
  
  
  # 배열을 인덱스 2로 나눈거에 1더해 채우는...
