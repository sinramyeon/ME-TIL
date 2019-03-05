// 트리는 퍼포먼스가 좋으나 공간 효율성이 좋지 않다
// 그래서 각 노드마다 여러 노드에 대한 레퍼런스를 가지고 있는 TST를 사용한다 ...(무슨소리임)

//TST 라고 하나봄

type Node struct {
  data byte
  isLastChar bool
  left, equal, right *Node
}

type TST struct {
  root *Node
}

func (t *TST) Insert(word string) {
  t.root = t.insertUtil(t.root, word, 0)
}

func (t *TST) insertUtil(curr *Node, word string, wordIndex int) *Node {
      if curr == nil {
            curr = new(Node)
            curr.data = word[wordIndex]
      }
      if word[wordIndex] < curr.data {
            curr.left = t.insertUtil(curr.left, word, wordIndex)
      } else if word[wordIndex] > curr.data {
            curr.right = t.insertUtil(curr.right, word, wordIndex)
      } else {
            if wordIndex < len(word)-1 {
                  curr.equal = t.insertUtil(curr.equal, word, wordIndex+1)
            } else {
                  curr.isLastChar = true
            }
      }
      return curr
}
 
func (t *TST) findUtil(curr *Node, word string, wordIndex int) bool {
      if curr == nil {
            return false
      }
      if word[wordIndex] < curr.data {
            return t.findUtil(curr.left, word, wordIndex)
      } else if word[wordIndex] > curr.data {
            return t.findUtil(curr.right, word, wordIndex)
      } else {
      if wordIndex == len(word)-1 {
        return curr.isLastChar
      }
      return t.findUtil(curr.equal, word, wordIndex+1)
      }
    }
    
func (t *TST) Find(word string) bool {
  ret := t.findUtil(t.root, word, 0)
  if ret {
    fmt.Println("Found")
  } dlse {
    fmt.Println("NO")
  }
  return ret
}
