#[WIP]

n=10
def ans(n, cnt):
    while n > 1 :
        if n%3 == 0:
            print("3나누기", n, n/3, cnt+1)
            ans(n/3, cnt+1)
        if n%2 == 0:
            print("2나누기", n, n/2, cnt+1)
            ans(n/2, cnt+1)
        else:
            print("1빼기", n, n-1, cnt+1)
            ans(n-1, cnt+1)
    
    print("결과", cnt)
    return cnt
ans(10, 0)

X= 10

count = 0
result = [X]

def calculation(x):
    lst = []
    for i in x:
        lst.append(i-1)
        if i%3 ==0:
            lst.append(i/3)
        if i%2 ==0:
            lst.append(i/2)
    lst = set(lst)
    lst = list(lst)
    return lst

while True:
    if X == 1:
        break
    temp = result
    result = []
    result = calculation(temp)
    count += 1
    if min(result) == 1:
        break
        
print(count)
