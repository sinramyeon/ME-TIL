stack = []
result = ''
string = '<open>tag<close>'

for char in string:
    if char == '<':
        while not len(stack) == 0:
            result += stack.pop() #단어뒤집기~
            print("stack not empty : ", result)
        result += char
        print("stack empty :", result)
        tag = True # < 시작부터 더하다
    elif char =='>':
        result += char
        print("char > ", result)
        tag = False # > 나오기까지 더함
    elif tag:
        result += char 
        print("tag True", result)
    elif char == ' ':
        while not stack.empty():
            result += stack.pop()
        result += char
        # 빈거일땐 걍 뒤집고 넣어
    else:
        stack.append(char) # 스택에 더했음
        print("else ", stack);
while not len(stack) == 0:
    result += stack.pop()
print(result)
