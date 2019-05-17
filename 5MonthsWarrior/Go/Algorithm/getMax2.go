// 소팅 부터 하기

func getMax2(data []int) int {
  size := len(data)
  max := data[0]
  maxCount := 1
  curr := data[0]
  currCount := 1
  sort.Ints(data)
  
  for i:=1; i<size; i++ {
    if data[i] == data[i-1] {
       currCount ++
    } else {
      currCount = 1
      curr = data[i]
    }
    
    if currCount > maxCount {
      maxCount = curCount
      max - curr
    }
  }
  return max
}
