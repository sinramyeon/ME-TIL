// sorted or increasing order / decreasing order

func linearSearchSorted(data []int, value int) bool {
  size := len(data)
  for i:=0 ; i < size; i++{
    if value == data[i] {
      return true
    } else if value < data[i] {
      return false
    }
  }
  return false
}
