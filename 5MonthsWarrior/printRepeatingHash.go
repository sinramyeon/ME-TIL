func PrintRepeating(arr []int) {
  hs := make(Set)
  for _, val := range arr {
    if hs.Find(val) {
      fmt.Println(val)
    } else {
      hs.Add(val)
    }
  }
}

func PrintFirstRepeating(arr []int) {
  size := len(arr)
  hs := make(Counter)
  
  for i:=0;i<size;i++{
    hs.Remove(arr[i])
    if hs.Find(arr[i]) {
      fmt.Println(arr[i])
      return
    }
  }
}
