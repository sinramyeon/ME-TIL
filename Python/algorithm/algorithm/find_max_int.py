#### 주어진 숫자 n개 중 최댓값 찾기

# 1.
# 숫자 여러 개를 리스트로 만들자

# 2.
# 최댓값을 찾자!

# 내 풀이

findsum = [17, 92, 18, 33, 58, 7, 33, 42]
max_int = 0;

for x in findsum :

    if x > max_int :
        max_int = x

print("max_int : "+str(max_int))


# 책 풀이

def find_max(a) :
    n = len(findsum)
    max_v = findsum[0]

    for i in range (1, n) :
        if a[i] > max_v :
            max_v = a[i]

    return max_v

#### 주어진 숫자 n개 중 최댓값의 인덱스 반환하기

# 내 풀이

def where_is_max_index(a) :

    max_int_is_me = 0;

    for x in a :
        if x > max_int_is_me :
            max_int_is_me = x

    return a.index(max_int_is_me)

# 책 풀이

def find_max2(a) :
    n = len(a)
    max_idx = 0
    for i in range(1, n) :
        if a[i] > a[max_idx] :
            max_idx = i
    return max_idx

#### 연습문제

# 숫자 n개를 리스트로 입력받아, 최솟값을 구하는 프로그램을 만들어 보세요.

# 내 풀이

def find_min(a) :

    length = len(a)

    im_min = a[length-1]

    for x in a :
        if x < im_min :
            im_min = x
    return im_min


# 책 풀이

def find_min2(a) :
    n = len(a)
    min_v = a[0]

    for i in range(1, n) :
        if a[i] < min_v :
            min_v = a[i]
    return min_v