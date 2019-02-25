// 정렬된 데서 중간값을 찾아라?!

func getMedian(data []int) int {
  size := len(data)
  sort.Ints(data)
  return data[size/2]
}
