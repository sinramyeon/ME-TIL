def factorial(n):
    
    if n == 1:      # n이 1일 때
        return 1    # 1을 반환하고 재귀호출을 끝냄
    return n * factorial(n - 1) # 나머지는 계속 재귀쓰

def get_zero_cnt(fact) :
    num_list = list(map(int, str(num)))
    num_list.reverse()
    
    cnt = 0
        
    for n in num_list :
        if n == 0:
            cnt += 1
        else :
            return cnt

get_zero_cnt(factorial(10))
