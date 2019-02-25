// binary search 방식

func findKeyCount2(data []int, key int) int {
  size := len(data)
  firstIndex := findFirstIndex(data, 0, size-1, key)
  lastIndex := fintLastIndex(data, 0, size-1, key)
  
  return (lastIndex - firstIndex + 1)

}

func findFirstIndex(data []int, start int, end int, key int) int {
  if end < start {
    return -1
  }
  
  mid := (start+end) /2
  if key == data[mid] && (mid == start || data[mid-1]! = key) {
    return mid
  }
  
  if key <= data[mid] {
    return findFirstIndex(data, start, mid-1, key)
  }
  
  return findFirstIndex(data, mid+1, end, key)
}

fund findLastIndex(data []int, start int, end int, key int) int {
  if end < start {
    return -1
  }
  
  mid := (start+end) /2
  if key == data[mid] && (mid == end || data[mid-1]! = key) {
    return mid
  }
  
  if key <= data[mid] {
    return findLastIndex(data, start, mid-1, key)
  }
  
  return findLastIndex(data, mid+1, end, key)

}
