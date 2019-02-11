// 해시 테이블을 사용해서 찾기

func FindPair3(data []int, value int) bool {
   s:= make(Set)
   size := len(data)
   ret := false
   
   for i:=0; i<size; i++ {
    if s.Find(value-data[i]){
      fmt.Println(data[i], value-data[i])
      ret = true
    } else {
      s.Add(data[i])
    }
   }
   return ret
}
