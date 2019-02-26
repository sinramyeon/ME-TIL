// 처음 반복되는 거 찾기

func FirstRepeated(data []int) int {
  size := len(data)
  for i:=0; i<size; i++ {
    for j:=i+1; j<size; j++ {
      if data[i] == data[j] {
        return data[i]
      }
    }
    return 0
  }
}
