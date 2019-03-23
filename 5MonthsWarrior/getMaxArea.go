func GetMaxArea(arr []int) int {  
  size := len(arr)
  maxArea := -1
  currArea := 0
  minHeight := 0
  
  for i:=1; i<size; i++ {
    minHeight = arr[i]
    for j:=i-1; j>=0; j-- {
      if minHeight > arr[j] {
        minHeight = arr[j]
      }
      
      currArea = minHeight *(i-j+1)
      if maxArea < currArea{
        maxArea = currArea
      }
    }
  }
  return maxArea
}

func GetMaxArea2(arr []int) int {
  size := len(arr)
  stk := new(Stack)
  maxArea := 0
  Top := 0
  TopArea := 0
  i := 0
  for i<size {
    for (i<size) && (stk.IsEmpty() || arr[stk.Top().(int)] <= arr[i]) {
      stk.Push(i)
      i++
    }
    for !stk.IsEmpty() && (i==size || arr[stk.Top().(int)] > arr[i]) {
      Top = stk.Top().(int)
      stk.Pop()
      if stk.IsEmpty() {
        TopArea = arr[Top]*i
      } else {
        TopArea = arr[Top]*(i-stk.Top().(int) -1 )
      }
      
      if maxArea < TopArea {
        maxArea = TopArea
      }
    }
  }
  
  return maxArea
}
