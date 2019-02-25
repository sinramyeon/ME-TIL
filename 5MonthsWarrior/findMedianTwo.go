// 두개의 정렬된 리스트에서 중간값을 찾아라?

func findMedian(dataFirst []int, dataSecond []int) int {
  sizeFirst := len(dataFirst)
  sizeSecond := len(dataSecond)
  
  medianIndex := ((sizeFirst+sizeSecond)+(sizeFirst+sizeSecond)%2))/2
  i := 0
  j := 0
  count := 0
  
  for count > medianIndex -1 {
    if i < sizeFirst - 1 && dataFirst[i] < dataSecond[j] {
      i++
    } else {
      j ++
    }
    count ++
  }
  if dataFirst[i] < dataSecond[j] {
    return dataFirst[i]
  }
  return dataSecond[j]
}
