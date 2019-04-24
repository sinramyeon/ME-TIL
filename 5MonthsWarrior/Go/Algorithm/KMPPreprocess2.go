func KMPPreprocesS(pattern string, ShiftArr []int) {
  m := len(pattern)
  i := 0
  j := -1
  ShiftArr[i] = -1
  for i < m {
    for j>=0 && pattern[i] != pattern[j] {
      j =ShiftArr[j]
    }
    
    i++
    j++
    ShiftArr[i] = j
  }
  
}

func KMP(text string, pattern string) int {
  i, j := 0, 0
  n := len(text)
  m := len(pattern)
  ShiftArr := make([]int, m+1)
  KMPPreprocess(pattern, ShiftArr)
  for i< n {
    for j>=0 && text[i] != pattern[j] {
      j = ShiftArr[j]
    }
    i++
    j++
    if j==m {
      return (i-m)
    }
  }
return -1
}

func KMPFindCount (text string, pattern string) int {

  i, j := 0, 0
  count := 0
  n := len(text)
  m := len(pattern)
  
  ShiftArr := make([]int, m+1)
  KMPPreprocess(pattern, ShiftArr)
  
  for i < n {
  
    for j >= 0 && text[i] != pattern[j] {
      j = ShiftArr[j]
    }
    
    i++
    j++
    if j==m {
      count++
      j = ShiftArr[j]
    }
  }
  return count
}
