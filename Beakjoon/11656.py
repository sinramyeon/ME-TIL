test = 'baekjoon'
test_arr = list(test)
result = []

while len(test_arr) > 0:
        result.append(''.join(test_arr)) # 더하고
        test_arr.pop(0) # 1단어씩뺌

print(result)
