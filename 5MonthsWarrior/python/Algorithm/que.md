FIFO(First in Fisrt Out)

```
def put(item) :
  queue.append(item)
  
def get() :
  return queue.pop()
  
if __name__ == '__main__' :
  queue = []
  put(1)
  put(2)
  
  while queue :
    pop()
```
