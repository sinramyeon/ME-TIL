func InfixToPrefix(expn string) string {
  expn = reverseString(expn)
  expn = replaceParanthesis(expn)
  expn = InfixToPostfix(expn)
  expn = reverseString(expn)
  return expn
}

func reverseString(in string) string {
  expn := []rune(in)
  lower := 0
  upper := len(expn) -1
  for lower < upper {
    expn[lower], expn[upper] = expn[upper], expn[lower]
    lower++
    upper--
  }
  
  return string(expn)
  
}

func replaceParanthesis(str string) string {
  a := []rune(str)
  lower := 0
  upper := lend(a) - 1
  for lower <= upper {
    if a[lower] == '(' {
      a[lower] = ')'
    } else if a[lower] == ')' {
      a[lower] = '('
    }
    lower++
  }
return string(a)
}
