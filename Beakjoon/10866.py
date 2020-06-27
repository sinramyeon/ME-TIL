class deque():
    def __init__(self):
        self.deque = []
        self.size = 0
    def push_front(self, n):
        self.deque.insert(0, n)
        self.size += 1
    def push_back(self, n):
        self.deque.append(n)
        self.size += 1
    def pop_front(self,):
        if self.size == 0:
            return -1
        else:
            self.size -= 1
            return self.deque.pop(0)
    def pop_back(self):
        if self.size == 0:
            return -1
        else:
            self.size =- 1
            return self.deque.pop(-1)
    def my_size(self):
        return self.size
    def empty(self):
        if self.size == 0:
            return 1
        else:
            return 0
    def front(self):
        if self.size == 0:
            return -1
        else:
            return self.deque[0]
    def back(self):
        if self.size ==0:
            return -1
        else:
            return self.deque[-1]
