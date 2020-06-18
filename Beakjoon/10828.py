class Stack:
    def __init__(self):
        self.len = 0
        self.stack = []
     
    def push(self, x):
        self.stack.append(x)
        self.len += 1
    
    def pop(self, x):
        if self.len.len() == 0:
            print("-1")
        result = self.stack[self.len-1]
        print(result)
        del self.stack[self.len-1]
        self.len -= 1
    
    def size(self):
        print(self.len)
        
    def empty(self):
        if self.len == 0 :
            print("1")
        else:
            print("0")
  
    def top(self):
        if self.len == 0 :
            print("-1")
        else:
            print(self.list[-1])
 
stack = Stack()
