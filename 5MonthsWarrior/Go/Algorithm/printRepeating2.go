func printRepeating2(data []int) {
  size := len(data)
  sort.Ints(data) // Sort(data, size)
  
  for i:=1; i<size; i++ {
    if data[i] == data[i-1] {
      fmt.Println(data[i]
    }
  }
}
