E = 1 # 15배수
S = 2 # 28배수
M = 3 # 19배수
# 5266

result = 0
year = 1

while True :
    if (year - E) % 15 == 0 and (year - S) % 28 == 0 and (year - M) % 19 == 0 :
        result = year
        break
    year += 1

print(result)
