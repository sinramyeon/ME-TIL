  //write a method which will search a list for some given value
  
  func SequentialSearch(data []int, value int) bool {
  
  size := len(data)
  
  for i:=0; i<size; i++ {
    if value == data[i]{
      return true
    } 
  }
  return false
  }
