N = 4
input = [3, 5, 2, 7]

stack = []
result = [-1 for _ in range(N)]
stack.append(0)

print(input, stack, result)
print("-----------------")

i = 1
while stack and i < N:
    
    while stack and input[stack[-1]] < input[i]:
        
        result[stack[-1]] = input[i] # 뒤가더크면
        print("result 값추가", result)
        
        stack.pop() # -1 대신 값넣고, 스택에서 꺼내고
        print("stack.pop :", stack)
        
    stack.append(i) # 안크면 스택에 값추가하고
    print("stack.append :", stack)
    i += 1

for i in range(N):
    print(result[i], end = " ")

# print(result[-1])
