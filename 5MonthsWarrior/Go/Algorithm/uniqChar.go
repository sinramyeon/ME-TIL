// 중복없는 글자가 있으면 1을 리턴하는 펑션

func isUniqueChar(str string) bool {
  mp := make(map[bye]int)
  size := len(str)
  for i:=0; i<size; i++ {
    c := str[i]
    if mp[c] != 0 {
      // 중복이 있었을 때
      return false
    }
    mp[c] = 1
  }
  return true
}
