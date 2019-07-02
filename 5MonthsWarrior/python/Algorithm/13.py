# 반복문을 사용해서 최대 공약수 구하기
# 두 수를 받아서 최대 공약수를 찾아 출력

def gcd(p, q) :
   if p > q :
     a = p
     b = q
   else :
     a= q
     b = p
   whlie b > 0 :
     c = b
     b = a%b
     a = c
   return a


# ?
