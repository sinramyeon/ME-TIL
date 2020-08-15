a = [20,
7,
23,
19,
10,
15,
25,
8,
13] # 아홉난쟁이

result = sum(a) # 우선 다합친 합을 구해놨음
a.sort()

for i in range(9) :
    for j in range(i+1, 9) : # 난쟁이 두놈을 차례로 뽑았음
        
        if result - a[i] - a[j] == 100: #  합에서 두놈뺀게 100이면
            
            for k in range(9):
                if k == i or k == j : # 두놈은무시하고
                    continue
                else: # 나머지애들을 확인
                    print(a[k]) 
