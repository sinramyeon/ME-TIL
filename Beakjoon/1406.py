from sys import stdin

# editer
# linked list?

# cursor at first / last / middle random
# cursor can be length+1 cases
# L left D right B delete left P add

stk1 = list(stdin.readline().strip())
stk2 = []

n = int(input())

for line in stdin:
    
    if line[0] == 'L': #left
        if stk1 :
            stk2.append(stk1.pop())
        else:
            continue
    elif line[0] == 'D': #right
        if stk2 :
            stk1.append(stk2.pop())
        else:
            continue
    elif line[0] == 'B':
        if stk1:
            stk1.pop()
        else:
            continue
    elif line[0] == 'P':
        stk1.append(line[2])

print(''.join(stk1 + list(reversed(stk2))))

# 근데 잘 이해안감 왜 abcd 안됨?
