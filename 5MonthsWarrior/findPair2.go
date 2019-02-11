func findPair2(data []int, value int) bool {
  size := len(data)
  first := 0
  second := size - 1
  ret := false
  
  sort.Insts(data)
  for first < second {
    curr := data[first] + data[second]
    
    if curr == value {
      fmt.Println(data[first], data[second])
      ret = true
    } if curr < value {
      first ++
    } else {
      second --
    }
  }
  return ret
}
