func printRepeating4(data []int, intrange int) {
  size := len(data)
  count := make([]int, intrange)
  
  for i:=0 ; i<size; i++ {
    count[i] = 0
  }
  
  for i:=0; i< size; i++{
    if count[data[i]] == 1 {
      fmt.Println(data[i])
    } else {
      count[data[i]]++
    }
  }
}
