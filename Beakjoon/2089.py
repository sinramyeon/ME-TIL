#-2 진법

N = -13
res = ''
while N: 
    if N%(-2):  # -2로나눠
        res = '1' + res
        N = N//-2 + 1 
    else: # 아닌데
        res = '0' + res
        N //= -2
        
print(res)
