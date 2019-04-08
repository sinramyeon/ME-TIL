//가장 왼쪽점부터 시작하여 x-좌표값이 증가하는 순서로 점을 방문하여 
// 가장 오른쪽 점까지 도달한 다음, 
//다시 그 점부터 가장 왼쪽 점까지 x-좌표값이 감소하는 정렬의 순서대로 방문하는 가장 짧은 사이클을 찾는 것이다. 이러한 사이클을 Bitonic Cycle 이라고 부른다.

 //먼소리여 ㅅㅂ
 
 func SearchBitonicArray(data []int, key int) int {
  size := len(data)
  maxIndex, _ := FindMaxBitonicArray(data)
  k := BinarySearch(data, 0, maxIndex, key, true) {
    if k != -1 {
      return k
    } else {
      return BinarySearch(data, maxIndex+1, size-1, key, false)
    }
  }
 }
 
 func FindMaxBitonicArray(data []int) (int, bool) {
  size := len(data)
  start := 0
  end := size - 1
  mid := 0
  if size < 3 {
    return -1, false
  }
for start <= end {
            mid = (start + end) / 2
            if data[mid-1] < data[mid] && data[mid+1] < data[mid] /* maxima */ {
                  return mid, true
            } else if data[mid-1] < data[mid] && data[mid] < data[mid+1] /* increasing */ {
                  start = mid + 1
            } else if data[mid-1] > data[mid] && data[mid] > data[mid+1] /* decreasing */ {
              end = mid - 1
            } else {
              break
            }
         }
      return -1, false
 }
 
 func BinarySearch(data []int, start int, end int, key int, isInc bool) int {}
 
  if end < start {
    return -1
  }
  
  mid := (start + end) /2
  
  if key == data[mid] {
    return mid
  }
  
  if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
    return BinarySearch(data, start, mid-1, key, isInc)
  } else {
     return BinarySearch(data, mid+1, end, key, isInc)
  }
 
 
