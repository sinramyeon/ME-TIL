// 가장 느리고 만들기쉽당...
// 하나씩 쌍을 지어 비교해서 바꾸는 방식

func BubbleSort(arr []int, comp func(int, int) bool) {
  size := len(arr)
  for i:=0; i < (size-1); i++ {
    for j:=0; j<size-i-1; j++ {
      if comp(arr[j], arr[j+1]) {
        arr[j+1], arr[j] = arr[j], arr[j+1]
      }
    }
  }
}

// 이미 거의 정렬된 데이터에 유효함
// 예를들어 데이터 한쌍만 정렬이 안됐다거나
