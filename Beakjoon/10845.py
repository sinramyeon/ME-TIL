class Queue:
    def __init__(self):
        self.queue = []
 
    def push(self, num):
        self.queue.append(num) # +
 
    def pop(self):
        if len(self.queue) != 0:
            return self.queue.pop(0) 
        else:
            return -1
 
    def size(self):
        return len(self.queue)
 
    def empty(self):
        if self.size() == 0:
            return 1 
        else:
            return 0
 
    def front(self):
        if self.size() != 0 :
            return self.queue[0]
        else:
            return -1
 
    def back(self):
        if self.size() != 0:
            return self.queue[-1]
        else :
            return -1

num = int(input())
queue = Queue()
for _ in range(num):
    tokens = input().split()
 
    if tokens[0] == "push":
        queue.push(tokens[1])
    elif tokens[0] == "pop":
        print(queue.pop())
    elif tokens[0] == "size":
        print(queue.size())
    elif tokens[0] == "empty":
        print(queue.empty())
    elif tokens[0] == "front":
        print(queue.front())
    elif tokens[0] == "back":
        print(queue.back())
