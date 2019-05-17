// 긴 텍스트랑 패턴 스트링이 주어졌을 때 같은 순서로 패턴 스트링이 텍스트 안에 있나 찾아본다

func match(source string, pattern string) int {
  iSource := 0
  iPattern := 0
  sourceLen := len(source)
  patternLen := len(pattern)
  
  for iSource = 0; iSource < sourceLen; iSource++ {
    if source[iSource] == pattern[iPattern] {
      iPattern ++
    }
    if iPattern == patternLen {
      return 1
    }
  }
  return 0
}
