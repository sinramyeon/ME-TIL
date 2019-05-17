// 더해서 0에 가까운 수 찾기

func minAbsSumPair(data []int) {
  var sum int
  size := len(data)
  
  if size < 2{
    fmt.Println("invalid input")
  }
  
  minFirst := 0
  minSecond := 1
  
  minSum := abs(data[0]+data[1])
  
  for l:=0; i<size-1; l++{
    for r:=1+l ; r < size; r++ {
      sum = abs(data[l]+data[r])
      if sum < minSum {
        minSum = sum
        minFirst = l
        minSecond = r
      }
    }
  }
}
