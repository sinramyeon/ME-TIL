func printAnagram(a string) {
  n := len(a)
  printAnagramUtil([]rune(a), n, n)
}

func printAnagramUtil(a []rune, max int, n int) {
  if max == 1 {
    fmt.Println(string(a))
  }
  
  for i := -1; i <max-1; i++ {
    if i!= -1 {
      a[i], a[max-1] = a[max-1], a[i]
    }
    printAgagramUtil(a, max-1, n)
    if i!= -1 {
      a[i], a[max-1] = a[max-1], a[i]
    }
  }
}
