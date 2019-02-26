// 순열 체크를 해시로 한다고...
// 진짜 이렇게 순열 체크가 됨??

func checkPermutation2(data1 []int, data2 []int) bool {
  size1 := len(data1)
  size2 := len(data2)
  h := make(map[int]int)
  
  if size1 != size2 {
    return false
  }
  
  for i:=0; i<size1; i++ {
    h[data1[i]]++
    h[data2[i]]--
    
  }
  
  // 해시를 즐리고 늘여서
  for i:=0; i<size1; i++ {
    if h[data1[i]] !=0 {
      return false
    }
  }
return true
}
