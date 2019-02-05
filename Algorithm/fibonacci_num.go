// given n find the nth number in the fibonacci....

func fibonacci(n int) int {

  if n <=1 {
    return n
  }
  
  retun fibonacci(n-1) + fibbonacci(n-2)

}
