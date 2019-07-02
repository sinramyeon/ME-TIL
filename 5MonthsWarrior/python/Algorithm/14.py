# 재귀호출로 최대공약수 구하기
def gcd(p, q) :
  if q == 0 :
    return p
  return gcd(q, p%q)
  
# 나눈걸 계속 리턴...
# 공통된 약수중 가장 큰 약수
