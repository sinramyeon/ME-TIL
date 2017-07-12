# 선택 정렬
## 자료를 크기 순서대로 맞춰 일렬로 나열하기

def find_min_idx(a) :
    n = len(a)
    min_idx = 0

    for i in range(1, n) :
        if a[i] < a[min_idx] :
            min_idx = i
    return min_idx


def sel_sort(a) :
    result = []

    while a : #리스트내에 값이 남아있을 때 동안
        min_idx = find_min_idx(a)
        value = a.pop(min_idx)
        result.append(value)

    return result


d = [1,3,4,5,6,76,4,3,2,5,23,54,2,1,24,45,523]

print(sel_sort(d))


#### 좀 더 효율적으로!

def sel_sort_efficient(a) :
    n = len(a)
    for i in range(0, n-1) : #0부터 n-2까지 반복

        min_idx = i
        for j in range(i+1, n) :
            if a[j] < a[min_idx] :
                min_idx = j

        #찾은 최솟값 i 위치로 변경

        a[i], a[min_idx] = a[min_idx], a[i]

sel_sort_efficient(d)