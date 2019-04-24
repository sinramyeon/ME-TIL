func IsAnagram(str1, str2 string) bool {
  size1 := len(str1)
  size2 := len(str2)
  
  if size1 != size2{
    return false
  }
  cm := make(Counter)
  for _, ch := range str1 {
    cm.Add(ch)
  }
  
  for _, ch := range str2 {
    cm.Remove(ch)
  }
  
  return (len(cm) == 0 )

}
