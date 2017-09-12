# 병합 정렬

def merge_sort(a) :
    n = len(a)
    # 종료조건 : 정렬할 리스트의 자료 개수가 한 개 이하라면 정렬할 필요가 없음
    if n <= 1 :
        return a
    # 그룹을 나누어 각각 병합 정렬을 호출
    mid = n//2 # n을 2로 나눈 몫
    # 중간을 기준으로 두 그룹을 나눔
    g1 = merge_sort(a[:mid]) # 첫 그룹
    g2 = merge_sort(a[mid:]) # 둘째 그룹

    # 두 그룹 병합
    result = []
    while g1 and g2 : # 두 그룹에 모두 자료가 남아 있는동안 반복
        if g1[0] <g2[0] : # 두 그룹 맨 앞 자료 값 비교
            # 작은걸 빼낸다
            result.append(g1.pop(0))
        else :
            result.append(g2.pop(0))

        # 아직 남은 자료들을 결과에 추가
        # g1과 g2중 빈 것은 whlie을 통과할거고
    while g1 :
        result.append(g1.pop(0))
    while g2 :
        result.append(g2.pop(0))
    return result

def merge_sort2(a) :
    # 자료 순서를 직접 바꿈
    n = len(a)
    if n<=1 :
        return
    mid = n//2
    g1=a[:mid]
    g2=a[mid:]
    merge_sort2(g1)
    merge_sort2(g2) # 재귀호출로 여기서 정렬부터 먼저 시킴
    i1=0
    i2=0
    ia=0

    # 병합 정렬
    while i1<len(g1) and i2<len(g2) :
        if g1[i1] < g2[i2] :
            a[ia] = g1[i1]
            i1+=1
            ia+=1
        else :
            a[ia] = g2[i2]
            i2+=1
            ia+=1
    # 아직 남은 자료를 추가

    while i1<len(g1) :
        a[ia] = g1[i1]
        i1+=1
        ia+=1
    while i2<len(g2) :
        a[ia]=g2[i2]
        i2 +=1
        ia +=1

d = [6,4,2,54,3,2,5,6,423,34,23,645]
print(merge_sort(d))