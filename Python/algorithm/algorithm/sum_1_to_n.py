# 1부터 n까지의 합 구하기


#### 내 답안

def mysum(n) :

    sum = 0;
    count_int = 1;

    for i in range(count_int, n+1) :

        sum+= i

    return sum


print(mysum(10))


#### 가우스 방정식으로 해볼까?

def sum_n(n) :
    return n * (n+1) / 2

print(sum_n(10))


# 연습문제 1.
# 1부터 n까지의 연속한 숫자의 제곱의 합을 구하세요.

def my_nn_sum(n) :

    sum = 0;

    for i in range(1, n+1) :

        sum += i*i

    return sum


print(my_nn_sum(3))