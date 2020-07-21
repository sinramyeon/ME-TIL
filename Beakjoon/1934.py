# 최소공배수 찾기

def LCM(X, Y):
    if X<Y:
        max = Y
    else:
        max = X
        
        # 최소~ 최대
    for i in range(max,(X*Y)+1):
        if i%X==0 and i%Y==0:
            res = i
            break
    return res
