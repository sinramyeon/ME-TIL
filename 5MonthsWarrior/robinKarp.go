// Robin-Karp 알고리즘
// 브루트포스랑 비슷함
// 덩어리로 떼어서 비교함

func RobinKarp(text string, pattern string) int {
  n := len(text)
  m := len(pattern)
  prime := 101
  powm := 1
  TextHash := 0
  i, j := 0,0
  if m == 0 || m > n {
    return -1
  }
  
  for i=0 ; i<m-1; i++ {
    powm = (powm << 1) % prime
  }
  
  for i=0 ; i<m; i++ {
    PatternHash = (PatternHash << 1) + (int)(pattern[i]))
    TextHash = ((TextHash <<1) + (int)(text[i]))%prime
  }
  
  for i=0; i <= (n-m); i++ {
    if TextHash == PatternHash{
      for j=0; j<m; j++ {
        if text[i+j] != pattern[j] {
          break
        }
      }
      
      if j == m {
        return i
      }
    }
    
    TextHash = ((TextHash - (int)(text[i]*powm) << 1) + (int)(text[i+m])) % prime
    if TextHash < 0 {
      TextHash = TextHash+prime
    }
  }
  return -1
}
