// bucket sort
// is the simplest and most efficient type of sorting
func BucketSort(data []int, lowerRange int, upperRange int) {
  rng := upperRange - lowerRange
  size := len(data)
  
  count := make([]int, rng)
  
  for i:=0; i<rng; i++ {
    count[i] = 0
  }
  for i:=0; i<size; i++ {
    count[data[i]-lowerRange]++
  }
  j:=0
  for i:=0; i<rng; i++ {
    for ; count[i] > 0 ; count[i] -- {
      data[j] = i + lowerRange
      j++
    }
  }
}
//먼소리임ㅂ대체?ㅆㅂ
