// given a sorted list
// fine the number of occurrences

func findKeyCound(data []int, key int) int {
  count := 0
  size := len(data)
  
  
  for i:=0; i<size; i++ {
    if data[i] == key {
      count ++
    }
  }
  
  return count
}
