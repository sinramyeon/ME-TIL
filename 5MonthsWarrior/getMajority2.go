// sort 한 후 과반수 찾기

func getMajority2(data []int) (int, bool) {
  size := len(data)
  majIndex := size/2
  
  sort.Ints(data)
  candidate := data[majIndex] // 정렬 후 가운데 걸 뽑음
  
  count :=0
  for i:=0; i<size; i++{
    if data[i] == candidate {
      count++
    }
  }
  if count > size/2{
    return data[majIndex], true
  }
  
  return 0, false
}
