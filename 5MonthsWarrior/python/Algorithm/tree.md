이진 트리

자식 노드를 2개 이하로 갖는 트리

```
class Node :
  def __init__(self, data) :
    self.data = data
    self.left = None
    self.right = None
    
def init_tree() :
  global root
  
  new_node = Node("A")
  root = new_node
  new_node = Node("B")
  root.left = new_node
  new_node = Node("C")
  root.right = new_node
  new_node_1 = node("D")
  new_node_2 = node("E")
  node = root.left
  node.left = new_node_1
  node.right = new_node_2
  
  new_node_1 = Node("F")
  new_node_2 = Node("G")
  node = root.right
  node.left = new_node_1
  node.right = new_node_2
```

트리구조의 규칙은 *노드는 오직 한 번만 방문* 이다.

---

1. 전위순회
가운데 > 왼쪽 > 오른쪽

```
def preorder_traverse(node) :
  if node == None : return
  print(node.data)
  preorder_traverse(node.left)
  preorder_traverse(node.right)
```

2. 중위순회
왼쪽 > 부모 > 오른쪽
```
def inorder_traverse(node) :
  if node == None : return
  inorder_traverse(node.left)
  print(node.data)
  inorder_traverse(node.right)
```

3. 후위 순회
왼쪽 자식 > 부모 > 오른쪽 자식
중위와 다르게 마지막으로 부모 노드를 방문하게 된다.

```
def postorder_traverse(node) :
  if node == None : return
  postorder_traverse(node.left)
  postorder_traverse(node.right)
  print(node.data)
```

4. 단계 순회

루트부터 단계대로 왼쪽부터 오른쪽으로

```
levelq = []

def levelorder_traverse(node) :
  global levelq
  levleq.append(node)
  
  while len(levelq) != 0 :
    visit_node = levleq.pop(0)
    print(visit_node.data)
    
    if visit_node.left != None :
      levelq.append(visit_node.left)
    if visit_node.right != None :
      levelq.append(visit_node.right)
```
