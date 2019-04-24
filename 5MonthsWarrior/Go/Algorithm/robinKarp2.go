func RobinKarp(text string, pattern string) int {

  n:=len(text)
  m:=len(pattern)
  prime := 101
  pown := 1
  TextHash := 0
  PatternHash := 0
  j, j := 0, 0
  if m ==0 || m>n {
    return -1
  }
  
  for i=0; i <m-1; i++ {
    PatternHash = ((PatternHash<<1)+(int)(pattern[i])) % prime
    TextHash = ((TextHash <<1) + (int)(text[i]) % prime
  }
  
  for i=0; i<=(n-m); i++ {
    if TextHash == PatternHash {
      for j=0; i<m; j++ {
        if tex[i+j] != pattern[j] {
          break
        }
        if j == m {
          return i
        }
      }
              TextHash = (((TextHash - (int)(text[i])*powm) << 1) + (int)(text[i+m])) % prime
            if TextHash < 0 {
                  TextHash = (TextHash + prime)
            }
      }
      return -1
}
