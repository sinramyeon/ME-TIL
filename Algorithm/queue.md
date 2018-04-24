# 큐(Queue)

자료구조의 기본인 큐를 알아보자.
큐는 스택과 마찬가지로 삭제의 위치와 방법이 제한적인 유한적 순서 리스트이다.
스택과 달리 먼저 삽입된 것부터 순서대로 삭제되는 **FIFO(First In First Out)** 구조이다.
참고로 스택은 나중에 집어 넣은 데이터가 먼저 나오는 방식이다.

이러한 큐 구조는 컴퓨터 과학 전반에 자주 쓰이는 자료구조이다. 가장 대표적인 예로 ‘버퍼(Buffer)’를 들 수 있는데, 데이터가 마구 입력이 되면, 이를 처리하지 못할 수가 있있으므로 이를 해결 하기 위해, 순서대로 입력된 데이터를 보관해두는 곳을 바로 버퍼라고 한다.

큐는 put(insert)와 get(delete)을 이용하여 구현된다. put는 큐에 자료를 넣는 것을, get은 큐에서 자료를 꺼내는 것을 의미한다. front(head)와 rear(tail)는 데이터의 위치를 가리킨다. front는 데이터를 get(가져오기)할 수 있는 위치를, rear은 데이터를 put(넣기)할 수 있는 위치를 의미한다. 또, 큐가 꽉 차서 더 이상 자료를 넣을 수 없는 경우(put 할 수 없는 경우)를 오버플로우(Overflow), 큐가 비어 있어 자료를 꺼낼 수 없는 경우(get 할 수 없는 경우)를 언더플로우(Underflow)라고 한다.


## 선형 큐(ArrayQueue)

![](http://3.bp.blogspot.com/-4DqngffLAMk/VOlx-CuarjI/AAAAAAAAAms/tEKrYsEY0PM/s1600/%EA%B7%B8%EB%A6%BC2.PNG)

이러한 모양으로 이루어져 있는데, 데이터를 제거하면 아래와 같이 된다.

![](http://3.bp.blogspot.com/-QwfV3ahT6gc/VOlx-JkKyWI/AAAAAAAAAmo/2iqbT16NgGo/s1600/%EA%B7%B8%EB%A6%BC3.PNG)

이러한 제거연산(Dequeue)을 계속 수행했다고 가정한다. 그렇다면 가장 첫번째 있는 요소를 빼내고 나면, 그 뒤에 있는 원소들을 하나씩 하나씩 앞으로 옮겨줘야 합니다. 이게 규모가 작을때는 크게 상관이 없겠으나 연산이 많아지면 아주 귀찮고 오래 걸리게 된다.

![](http://1.bp.blogspot.com/-cfP7gbz7HGo/VOlyzOyx7II/AAAAAAAAAm4/3vH6PfyfbJE/s1600/%EA%B7%B8%EB%A6%BC4.PNG)

이 문제를 해결하기 위해, 전단과 후단의 위치를 기억해서 전단에서 데이터를 뺀 후 후단에 입력하는 방식도 있다.

![](http://2.bp.blogspot.com/-UqnxVg4okB4/VOly9CxSerI/AAAAAAAAAnA/1VYzezVfxkg/s1600/%EA%B7%B8%EB%A6%BC5.PNG)

하지만 이 구조 또한 단점을 지니는데, 전단의 데이터를 앞으로 옮겨주지 않아 후단에 데이터가 들어올 곳이 없어져 버린다. 즉 제거 연산을 수행할 수록 큐의 가용 용량이 줄어든다.

[참고](http://logonluv.blogspot.kr/2015/02/datastructure-queue.html)

## 환형 큐(CircularQueue)

원형 큐는 이와같은 단점을 보완하는 구조로 큐의 맨 끝과 맨 처음이 연결된 형태의 구조이기 때문에 이미 꺼내어 사용한 큐의 앞부분에 다시 데이터를 저장하여 재사용이 가능한 형태의 큐이다. 데이터가 순환하며 저장되기 때문에 끝이라는 개념도 없다.

![](http://cfile27.uf.tistory.com/image/2102763B5602C2D42B5384)

포화 상태에서는 한 칸의 여유를 둔다. 더이상 값을 삽입할 수 없는 상태.

![](http://cfile22.uf.tistory.com/image/235C374B5616967229FCE9)

다시 모든 값을 제거하면 시작과 끝이 같아진다.


---

### python queue

1. list를 사용

```
class Queue(list):
    # enqueue == > insert 관용적인 이름
    enqueue = list.append

    # dequeue == > delete
    def dequeue(self):
        return self.pop(0) # 넣은것부터 꺼내기

    def is_empty(self):
        if not self:
            return True
        else:
            return False

    def peek(self):
        return self[0]

if __name__ == '__main__':
    q = Queue()
    q.enqueue(1)
    q.enqueue(2)
    q.enqueue(3)
    q.enqueue(4)
    q.enqueue(5)

    while not q.is_empty():
        print(q.dequeue(), end= ' ') # 1 2 3 4 5
 ```

2. Node를 사용

```
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class Queue:
    def __init__(self):
        self.head = None
        self.tail = None

    def is_empty(self):
        if not self.head:
            return True

        return False

    def enqueue(self, data):
        new_node = Node(data)

        if self.is_empty():
            self.head = new_node
            self.tail = new_node
            return

        self.tail.next = new_node
        self.tail = new_node

    def dequeue(self):
        if self.is_empty():
            return None

        ret_data = self.head.data
        self.head = self.head.next
        return ret_data

    def peek(self):
        if self.is_empty():
            return None

        return self.head.data
```

---

### golang queue

```
package QueueLinkedList

type Node struct {
	data int
	next *Node
}

type Queue struct {
	rear *Node
}

func (list *Queue) Enqueue(i int) {
	data := &Node{
		data: i
	}
	
	if list.rear != nil {
		data.next = list.rear
	}
	list.rear = data
}

func (list *Queue) Dequeue() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	if list.rear.next == nil {
		i := list.rear.data
		list.rear = nil
		return i, true
	}
	current := list.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

func (list *Queue) Peek() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	return list.rear.data, true
}

func (list *Queue) Get() []int {
	var items []int
	current := list.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Queue) IsEmpty() bool {
	return list.rear == nil
}

func (list *Queue) Empty() {
	list.rear = nil
}
```


