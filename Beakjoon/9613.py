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

##########

# 그후 내가 고친거(타임아웃)

import sys
from itertools import combinations

input = sys.stdin.readline

# 최대공약수
def gcd(x,y):
    if y == 0: # 나머지
        return x # 최대공약수
    return gcd(y,x%y)

t = int(input())
for i in range(t):
    num_list = list(map(int, input().split()))
    num_list = num_list[1:] # 개수빼고

    all_gcd_list = combinations(num_list, 2)
    sum = 0
    for a, b in all_gcd_list:

        gcd_result = gcd(a, b)
        sum += gcd_result


print(sum)

#################3

# 남의답 (뭔차인지?)

import sys
import itertools
input = sys.stdin.readline

def gcd(a, b):
    return gcd(b, a % b) if b else a

t = int(input())
for i in range(t):
    num_list = list(map(int, input().split()))
    num_list = num_list[1:]
    ans = 0
    for a, b in itertools.combinations(num_list, 2):
        ans += gcd(a, b)
    print(ans)
