// 퀵 선택 정렬 알고리즘

func QuickSelect(arr []int, key int) int {
  size := len(arr)
  QuickSelectUtil(arr, 0, size-1, key)
  return arr[key-1]
}

func QuickSelectUtil(arr []int, lower int, upper int, key int) {

  if upper <= lower {
    return
  }
  
  pivot := arr[lower]
  start := lower
  stop := upper
  
  for lower < upper {
    for arr[lower] <= pivot && lower < upper{
      lower++
    }
    for arr[upper] > pivot && lower <= upper {
      upper--
    }
    if lower < upper {
      swap(arr, upper, lower)
    }
  }
  swap(arr, upper, start) // upper is the pivot position...
  // what the heck if pivot... 중심점
  
  QuickSelectUtil(arr, start, upper-1, key)
  QuickSelectUtil(arr, upper+1, stop, key)

}

fun swap(arr []int, first int, second int) {
  arr[first], arr[second] = arr[second], arr[first]
}
