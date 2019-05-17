//reverse order of words in a string sentence

func reverseStringRange(a []rune, lower int, upper int) {
  for lower < upper {
    a[lower], a[upper] = a[upper], a[lower]
    lower++
    upper--
  }
}

func reverseWords(str string) string {
  length := len(str)
  upper := -1
  lower := 0
  arr := []rune(Str)
  
  for i:=0 ; i <length; i++ {
    if arr[i] == '' {
      reverseStringRange(arr, lower, upper)
      lower = i+1
      upper = i
    } else {
      upper++
    }
    reverseStringRange(arr, lower, upper)
    
    reverseStringRange(arr, 0, length-1)
    
    return string(arr)
  }
}
