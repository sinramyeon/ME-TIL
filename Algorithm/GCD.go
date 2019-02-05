//최대공약수 찾기

func GCD(m int, n int) int {

  if m < n {
    return GCD(n, m)
  }
  
  if m%n == 0 {
    return n
  }
  
  return GCD(n, m%n)

}
