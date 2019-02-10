// not ordered list search

func linearSearchUnsorted(data []int, value int) bool {
  size := len(data)
  for i:=0; i < size; i++{
    if value == data[i] {
      return true
    }
  }
return false
}
