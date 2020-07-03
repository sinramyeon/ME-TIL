num_arr = [1, 1, 2, 3, 4, 2, 1]
size = len(num_arr)


freq_arr = [0 for _ in range(size)]
for index in num_arr:
    
    freq_arr[index] += 1 # 1에 1추가 1에 1 더추가...이런식
    ans_arr = [-1 for _ in range(len(num_arr))]
    stack = []
    
print(freq_arr)
    
for index, content in enumerate(num_arr):
    while stack and freq_arr[num_arr[stack[-1]]] < freq_arr[content]:
        print(stack[-1], num_arr[stack[-1]], content)
        print(freq_arr[num_arr[stack[-1]]], freq_arr[content])
        ans_arr[stack.pop()] = content
        
    stack.append(index) # 우선더함
    
    
print(ans_arr)
