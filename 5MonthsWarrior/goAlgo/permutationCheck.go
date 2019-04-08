// 각 두 글자가 서로의 순열 조합인지 체크하기

func isPermutation(s1 string, s2 string) bool {
  count := make(map[byte]int)
  length := len(s1)
  if len(s2) != length {
    return false // 두 길이가 다르면 순열 조합이 아님
  }
  
  for i:=0 ; i <length; i++ {
   ch := s1[i]
   count[ch]++
   ch = s2[i]
   count[ch]--
  }
  
  for i:=0 ; i<length; i++ {
    ch := s1[i]
    if count[ch] != 0 {
      return false // 순열 조합이 아님
    }
    
    return true
  }
}
