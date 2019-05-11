1. 노드, 링크

파이썬에서 연결 리스트를 사용하기 위해서는, 노드를 다음과 같이 클래스로 정의해서 사용한다.

```
class Node :
  def __init__ (self, data, next=None) :
    self.data = data
    self.next = next
```

2. 연결 리스트의 삽입과 삭제 알고리즘

연결리스트 초기화

```
class Node :
  def __init__ (self, data, next=None) :
    self.data = data
    self.next = next
    
    
def init_list() :
  global node_A
  node_A = Node("A")
  node_B = Node("B")
  node_C = Node("C")
  node_D = Node("D")
  
  node_A.next = node_B
  node_B.next = node_C
  node_C.next = node_D
  
 
def print_list() :
  global node_A
  node = node_A
  while node :
    print(node.data)
    node = node.next
  print
  
  
if __name__ == '__main__' :
  init_list()
  print_list()
```

연결리스트 삽입

```
class Node :
  def __init__ (self, data, next=None) :
    self.data = data
    self.next = next
    
    
def init_list() :
  global node_A
  node_A = Node("A")
  node_B = Node("B")
  node_C = Node("C")
  node_D = Node("D")
  
  node_A.next = node_B
  node_B.next = node_C
  node_C.next = node_D
  
 
def print_list() :
  global node_A
  node = node_A
  
```


연결 리스트의 삽입과 삭제

```
def delete_node(del_data) :
  global node_A
  pre_node = node_A
  next_node = pre_node.next
  
  
  if pre_node.data == del_data :
    node_A = next_node
    del pre_node
    print
  while next_node :
    if next_node.data == del_data :
      pre_node.next = next_node.next
      del next_node
      break
    pre_node = next_node
    next_node = next_node.next
    
def insert_node(data):
  global node_A
  new_node = Node(data)
  node_P = node_A
  node_T = node_A
  
  while node_T.data <= data :
    node_P = node_T
    node_T = node_T.next
  new_node.next = node_T
  node_P.next = new_node
```

삭제 알고리즘은 아래와 같다.

```
def delete_node(del_data) :
  global node_A
  pre_node = node_A
  next_node = pre_node.next
  
  if pre_node.data == del_data :
    node_A = next_node
    del pre_node
    return
    
  while next_node :
    if next_node.data == del_data :
      pre_node.next = next_node.next
      del next_node
      break
  
  pre_node = next_node
  next_node = next_node.next
```
