// 1부터 n까지 숫자가 주어졌는데 그 중 하나가 빠진 n-1개의 리스트가 있다
// 빠진 수를 찾아라

func findMissingNum(data []int) (int, bool) {
  var found int
  size := len(data)
  
  for i:=1; i<=size; i++{
    found = 0
    for j:=0; j<size; j++{
      if data[j] == i {
        found = 1
        break
      }
    }
    if found == 0 {
      return i, true
    }
  }
  return 0, false
}
