    #리스트에서 특정숫자 위치 찾기
    #입력 리스트 a, 찾는값 x
    # 출력 찾으면 값 위치, 못찾으면 -1 반환
    
def binary_search(a, x) :
    start = 0 #처음부터 시작
    end = len(a) -1 #마지막까지
    
    while start <= end :
        mid = (start+end)  // 2 # 중간부터 찾음
        if x == a[mid] :
            return mid
        elif x > a[mid] :
            start = mid+1
        else x < a[mid] :
            start = mid-1
            
    return -1 # 못찾으면
    
def binary_search_advanced(a, x) :
    start = 0
    end = len(a) -1
    
    if start == end : # 빈 배열이라면 굳이 찾지 말자
        return
    
    while start <= end :
        mid = (start+end) // 2
        if x == a[mid] :
            return mid
        elif x > a[mid] :
            binary_search_advanced(a[mid:], x) #재귀호출
        else x < a[mid] :
            binary_search_advanced(a[:mid], x) #재귀호출
    return -1
