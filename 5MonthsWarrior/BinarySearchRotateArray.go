// given a sorted list s of n integer
// s is rotated an unknown number of times
// find an element in the list
// 뭔소리야?

func BinarySearchRotateArray(data []int, key int) int {
  size := len(data)
  return BinarySearchRotateArrayUtil(data, 0, size-1, key)
}

func BinarySearchRotateArrayUtil(data []int, start int, end int, key int) int {
  if end < start {
    return -1
  }
  
  mid := (start+end) / 2
  if key == data[mid] {
    return mid
  }
  
  if data[mid] > data[start] {
    if data[start] <= key && key < data[mid] {
      return BinarySearchRotateArrayUtil(data, start, mid-1, key)
    }
    return BinarySearchRotateArrayUtil(data, mid+1, end, key)
  } else {
  
    if data[mid] < key && key <= data[end] {
      return BinarySearchRotateArrayUtil(data, mid+1, end, key)
    }
    return BinarySearchRotateArrayUtil(data, start, mid-1, key)
  }

}
