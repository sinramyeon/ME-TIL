// counting

func getMax3(data []int, dataRange int) int {
  max := data[0]
  maxCount := 1
  size := len(data)
  count := make([]int, dataRange)
  
  for i:=0; i<size; i++{
    count[data[i]]++
    if count[data[i]] > maxCount {
      maxCount = count[data[i]]
      max = data[i]
    }
  }
  return max
}
