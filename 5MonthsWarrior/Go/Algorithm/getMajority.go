// 자주 나타나는(2회 이상) 숫자 찾기
// 없으면 0

func getMajority(data []int) (int, bool) {
  size := len(data)
  max := 0
  count := 0
  maxCount := 0
  
  for i:=0 ; i<size; i++ {
    for j:=i+1; j<size; j++{
      if data[i] == data[j] {
        count ++
      }
    }
    if count > maxCount {
      max = data[i]
      maxCount = count
    }
  }
  
  if maxCount > size/2 {
    return max, true
  }
  
  return 0, false
}
