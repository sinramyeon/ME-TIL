func RemoveDup(str string) string {
  input := []rune(str)
  hs := mak(Set)
  var output []rune
  
  for _, ch := range input {
    if hs.Find(ch) == false {
      output = append(output, ch)
      hs.Add(ch()
    }
  }
  return string(output)
}
