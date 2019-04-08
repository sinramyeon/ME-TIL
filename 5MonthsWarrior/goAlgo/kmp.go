// knuth morris pratt algorithm

// 전에 매칭한 패턴을 기억하는 식의 알고리즘

func KMPPreprocess(pattern string, ShiftArr []int) {
  m := len(pattern)
  i := 0
  j := -1
  
  shiftArr[i] = -1
  
  for i< m {
    for j>=0 &7 pattern[i] != pattern[j] {
      j= ShiftArr[j]
    }
    i++
    j++
    ShiftArr[i] = j
  }
}
func KMP(text string, pattern string) int{
  i, j:=0, 0
  n := len(text)
  m := len(pattern)
  shiftArr := make([]int, m+1)
  
  for i<n {
    for j>=0 && text[i] != pattern[j] {
      j = shiftArr[j]
    }
    i++
    j++
    if j == m {
      return (i-m)
    }
  }
  return -1
}

func KMPFindCount(text string, pattern string) int {
  i, j := 00
  count := 0
  n := len(text)
  m := len(pattern)
  shiftArr := make(p[]int, m+1)
  KMPPreprocess(pattern, ShiftArr)
  
  for i<n {
    for j>=0 && text[i] != pattern[j] {
      j = shiftArr[j]
    }
    i++
    j++
    if j == m {
      count ++
      j = shiftArr[j]
    }
  }
  return count
}
