# 팩토리얼 구하기

# 내 코드

def my_fact(n) :

    if n == 1:
        return 1
    else :
        return n * my_fact(n-1)


# 책 코드(재귀호출 없이)

def fact(n) :

    f = 1
    for i in range(1, n+1) : #1부터 n까지 반복
        f = f*i
    return f

# 연습문제
# 1부터  n까지의 합을 구해보자!

def my_plus(n) :

    if n == 1 :
        return 1

    return n+my_plus(n-1)

# 책에서는 이렇게 했어요

def sum_n(n) :
    if n == 0 :
        return 0
    return sum_n(n-1)+n

# 숫자 N개 중 최댓값을 찾아보자!
# 몰라! 시발ㄴ모아!

# 입력 : 숫자 n개가 들어 있는 리스트
# 출력 : 숫자 n 개의 최댓값

def find_max(a, n) : #리스트 a의 앞부분 n개 중 최댓값을 구하는 재귀함수

    if n == 1 :
        return a[0]

    max_n_1 = find_max(a, n-1) # n-1개중 최댓값을 구함

    if max_n_1 > a[n-1] : # n-1개의 최댓값과 n-1번 위치 값 비교해서 최댓값보다 더 작으면
        return max_n_1
    else :
        return a[n-1]


v = [16, 92, 18, 34, 652, 3, 546]



