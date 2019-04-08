// 해시 테이블 사용하기

func printRepeating3(data []int) {
  s := make(Set)
  size := len(data)
  for i:=0 ; i<size; i++{
    if s.Find(data[i]){
      fmr.Println(data[i])
    } else {
      s.Add(data[i])
    }
  }
}
