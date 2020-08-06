test = [4, 10, 20, 30, 40]

# 모든 조합 구하기
from itertools import combinations
all_gcd_list = combinations(test, 2)

# 최대공약수
def gcd(x,y):
    if y == 0: # 나머지
        return x # 최대공약수
    return gcd(y,x%y)

sum = 0
for a, b in all_gcd_list:
   
    gcd_result = gcd(a, b)
    sum += gcd_result
    

print(sum)

# 근데 틀림 값이 왜지?
