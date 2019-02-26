// given a list of 0s and 1s
// all the 0s come before 1s
// find the index ot the first 1

func BinarySearch01(data []int) int {
   size := len(data)
   if size == 1 && data[0] == 1 {
    return -1
   }
  return BinarySearch01Util(data, 0, size-1)
}


func BinarySearch01Util(data []int, start, end int) int {
  if end < start {
    return -1
  }
  
  mid := (start+end) / 2
  if 1 == data[mid] && 0 == data[mid-1] {
    return mid
  } if 0 == data[mid] {
    return BinarySearch01Util(data, mid+1, end)
  }
  
  return BinarySearch01Util(data, start, mid-1)

}
