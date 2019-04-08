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

func (t *TST) insertUtil(curr *Node, word string, wordIndex int) *Node{

  if curr == nil {
    curr = new(Node)
    curr.data = word[wordIndex]
  }
  
  if word[wordIndex] < curr.data{
    curr.left = t.insertUtil(curr.left, word, wordIndex)
  } else if word[wordIndex] > curr.data{
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

  if curr == nil{
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
    fmt.Println(word)
    
    if ret {
      fmt.Println("Found")
    }
   
   return ret
   }
   
   
   
   
   
