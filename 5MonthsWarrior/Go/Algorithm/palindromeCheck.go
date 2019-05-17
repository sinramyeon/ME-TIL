// palindrome 회문(回文: madam이나 nurses run처럼 앞에서부터 읽으나 뒤에서부터 읽으나 동일한 단어나 구)
// 로꾸꺼

func isPalindrome(str string) bool {

  i := 0
  j := len(str) - 1
  for i < j && str[i] == str[j] {
    i++
    j--
  }
  
  if i < j {
    return false
  }
  
  return true
}
