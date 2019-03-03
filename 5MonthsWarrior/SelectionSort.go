// 선택 정렬

func SelectionSort(arr []int) {
  size := len(arr)
  var i, j, max, temp int
  for i:=0; i< size-1; i++ {
    max = 0
    for j=1; j< size-1-i; j++ {
      if arr[j] > arr[max] {
        max = j
      }
    }
    // arr[0]부터 비교해서
    // 뭐가 더 크면 그거를 선택해서
    
    // 뒤에 있는 거랑 그 제일 큰걸 바꿔침
    arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
  }

}


// best wort average case!!!
// 프로토 타입 등에나 유효하지 가장 쓸모가 없음.
