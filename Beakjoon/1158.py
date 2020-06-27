def josepus(n, k):
    # make array
    num = 1
    array = []
    while num <= n :
        array.append(num)
        num += 1
    
    result = []
    
    # pick k
    pointer = k-1
    print("start: ", array)
    
    for i in range(n): # 여기서 in array 로 했다가 5 1 4 안나와서 한참 개고생함 ㅡㅡ
        
        if len(array) > pointer :
            result.append(array.pop(pointer))
            pointer += k-1
            
        elif len(array) <= pointer :
            pointer = pointer % len(array)
            result.append(array.pop(pointer))
            pointer += k-1
    print("result : ", result)
