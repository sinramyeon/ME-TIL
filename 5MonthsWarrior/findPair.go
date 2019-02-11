// 제공 된 숫자 중 두 개를 더하면 특정 값이 되는 페어를 찾자!

func findPair(data []int, value int) bool{
  size := len(data)
  ret := false
  
  for i :=0 ; i <size; i++{
    for j:=i+1; j<size; j++{
      if (data[i] + data[j]) == value {
        fmt.Println(data[i], data[j])
        ret = true
      }
    }
  }
  return ret
}
