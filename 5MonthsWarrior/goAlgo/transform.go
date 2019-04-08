// a1 a2 a3 a4 b1 b2를 a1 b1 a2 b2 ..로 바꾸기

func transformArrayAB1(str string) string {
  data := []rune(str)
  size := len(data)
  
  N := size/2
  for i:=1; i<N; i++ {
    for j:=0; j<i; j++ {
      data[N-i+2*j], data[N-i+2*j+1] = data[N-i+2*j+1], data[N-i+2*j]
    }
  }
  
  return string(data)
}
