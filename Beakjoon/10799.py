### 스택으로만들어봐...
inputs = input()
sumval = 0
left = 0

for i in range(len(inputs)):
    if inputs[i] == '(':
        left += 1
    elif inputs[i] == ')':
        left -= 1 # 괄호수세기
        if inputs[i-1] == '(': #레이저
            sumval += left
        else :
            sumval +=1
