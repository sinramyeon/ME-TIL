func getMax(data []int) {
  size := len(data)
  max := data[0]
  
  count := 1
  maxCount :=1
  for i:=0 ; i < size; i++{
    count = 1
    for j:=i+1; j<size; j++{
      if data[i] == data[j] {
        count++
      }
    }
    if count > maxCount {
      max = data[i]
      maxCount = count
    }
  }
  return max
}
