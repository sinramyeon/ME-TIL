// 중복 찾기

func printRepeating(data []int) {
  size := len(data)
  for i =0 ; i <size ; i++ {
    for j := i+1; j < size; j++ {
      if data[i] = data[j] {
        fmt.Println(data[i])
      }
    }
  }
}
