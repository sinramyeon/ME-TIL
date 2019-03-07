// x의 n승 찾기

func pow(x int, n int) int {
  var value int
  if n == 0 {
    return 1
  } else if n%2 == 0 {
    value = pow(x, n/2)
    return (value*value)
  } else {
  
    value = pow(x, n/2)
    return (x*value*value)
  }
}
