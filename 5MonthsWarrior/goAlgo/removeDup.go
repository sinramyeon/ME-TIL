// integer list에서 중복 제거하기

func removeDuplicates(data []int) int {
  j := 0
  size := len(data)
  
  if size == 0 {
    return 0
  }
  
  sort.Ints(data)
  for i:=1; i<size; i++ {
    if data[i] != data[j] {
      j++
      data[j] = data[i]
    }
  }
  return j+1
}
