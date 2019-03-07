// 리턴값

// 0 첫, 두번째가 같다
// 음수 처음게 뒤에것보다 짧다
// 양수 두번째거보다 처음게 길다

func strcmp(a string, b string) int {
  index := 0
  lend1 := len(a)
  len2 := len(b)
  minlen := len1
  if len1 > len2 {
    minlen = len2
  }
  
  for index < minlen && a[index] == b[index] {
    return 0
  }else if len 1 == index {
    return -1
  }else {
    return 1
  }
  
  return (int)(a[index) - (int)(b[index])
}
