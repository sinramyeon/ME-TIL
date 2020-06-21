class Stack():
    def __init__(self):
        self.num_list = []
        self.operation = []
 
    def push(self, num):
        self.num_list.append(num)
        self.operation.append('+')
 
    def pop(self):
        if len(self.num_list) != 0:
            self.operation.append('-')
            return self.num_list.pop()
        else:
            return -1
 
    def size(self):
        return len(self.num_list)
 
    def empty(self):
        if len(self.num_list) == 0:
            return 1
        else:
            return 0
 
    def top(self):
        if len(self.num_list) != 0:
            return self.num_list[-1]
        else:
            return -1
 
stack = Stack()
count = 1 # stack starts from 1

isExist = True

for i in range(int(input())):
    num = int(input())
    
    while count <= num:
        stack.push(count)
        count += 1
        
    if stack.num_list[-1] == num:
        stack.pop()
        
    else:
        isExist = False
        break
 
if isExist == False:
    print('NO')
else:
    for k in stack.plus_minus:
        print(k)
