fucn QuickSort(arr []int, comp func(int, int) bool) {

  size := len(arr)
  quickSortUtil(arr, 0, size-1, comp)
}

func quickSortUtil(arr []int, lower int, upper int, comp func(int, int) bool) {

  if upper <= lower {
    return
  }
  
  pivot := arr[lower]
  start := lower
  stop := upper
  
  for lower < upper {
    for comp(arr[lower], pivot) == false && lower < upper {
      lower ++
      for comp(arr[upper], pivot) && lower <= upper {
        upper--
      }
      if lower < upper {
        swap(arr, upper, lower) // 근데 대체 comp는 뭐하는거임
      }
    }
    
    swap(arr, upper, start)
    quickSortUtil(arr, start, upper-1, comp)
    quickSortUtil(arr, upper+1, stop, comp)
   
  }
}

func swap(arr []int, first int, second int) {
  arr[first], arr[second] = arr[second], arr[first]
}

// 언제쓰는가
// stable sort 가 필요할때
// 최악을 피한 평균 퍼포먼스를 찾을때
// 데이터가 랜덤일때
