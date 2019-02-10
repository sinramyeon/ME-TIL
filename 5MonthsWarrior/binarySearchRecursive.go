// 이진탐색 재귀호출

func binarySearchRecursive(data []int, low int, high int, value int) bool {
  if low > high {
    return false
  }
  
  mid := low+(high-low)/2
  
  if data[mid] == value {
    return true
  }else if data[mid] < value{
    return binarySearchRecursive(data, mid+1, high, value)
  }else {
    return binarySearchRecursive(data, low, mid-1, value)
  }

}
