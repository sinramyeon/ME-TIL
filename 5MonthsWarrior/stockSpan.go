// given a list of daily stock price in a list A[i]
// find the span of the stocks for each day
// a span of stock is the max num of days for which the price ot stock was lower than that day

func StockSpanRange(arr []int) []int {
  n := len(arr)
  SR := make([]int, n)
  SR[0] = 1
  for i:=1; i<len(arR); i++ {
    SR[i] = 1
    for j := i-1 ; (j>=0) && (arr[i] >= arr[j]); j-- {
      SR[i]++
    }
  }
  return SR
}


func StockSpanRange2(arr []int) []int {
  stk := new(Stack)
  n := len(arr)
  SR := make([]int, n)
  stk.Push(0)
  SR[0] = 1
  
  for i:=1; i<len(arr); i++ {
    for!stk.IsEmpty() && arr[stk.Top().(int)] <= arr[i] {
      stk.Pop()
    }
    
    if stk.IsEmpty() {
      SR[i] = i+1
    } else {
      SR[i] = i-stk.Top().(int)
    }
    stk.Push(i)
  }
  return SR
}
