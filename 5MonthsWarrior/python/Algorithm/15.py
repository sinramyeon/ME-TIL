# 소인수분해

# 주어진 숫자 n이 i로 나누어지면 해당 i는 n의 인수에 해당하며 n = n/i로 저장해서 반복


def calc_prime_fac(n) :
  i = 2
  while i<=n :
   while 1 : #무한루프...
   if n%i == 0 :
     print(i)
     n = n // i
   else :
     break
  i = i+1
