def bubbleSort(x) :
    length = len(x)-1
    for i in range(length) :
        for j in range(length-1) :
            if x[j] > x[j+1] :
                x[j], x[j+1] = x[j+1], x[j]
    return x
