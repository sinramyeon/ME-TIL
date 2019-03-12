func (list *List) RemoveLoop() {
  loopPoint := list.LoopPointDetect()
  if loopPoint == nil {
    return
  }
  
  firstPtr := list.head
  if loopPoint == list.head {
    for firstPtr = firstPtr.next{

  firstPrt.next = nil
  return
}

secondPtr := loopPoint
for firstPtr.next != secondPTr.next {
  firstPtr = firstPtr.next
  secondPtr =secondPtr.next
}

secondPtr.next = nil
}

func (list *List) LooopPointDetect() *Node {
  slowPtr := list.head
  fastPtr := list.head
  
  for fastPtr.next != nil && fastPtr.next.next != nil {
    slowPtr = slowPtr.next
    fastPtr = fastPtr.next.next
    if slowPtr == fastPtr {
      return slowPtr
    }
  }
  
  return nil
}



