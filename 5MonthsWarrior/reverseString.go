func reverseString(a string) string {
  chars := []rune(a)
  reverseStringUtil(chars)
  return string(chars)
}

func reverseStringUtil(a []rune) {
  lower := 0
  uppder := len(a)-1
  
  for lower<upper {
    a[lower], a[upper] = a[upper], a[lower]
    lower++
    upper--
  }
}
