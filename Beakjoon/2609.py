x = 24
y = 18

# 최대공약수
def gcd(x,y):
    mod = x % y
    while mod >0:
        x = y
        y = mod
        mod = x % y
    return y    

# 최소공배수
def lcm(x, y):
    return x * y // gcd(x,y)

print(gcd(a, b))
print(lcm(a, b))
