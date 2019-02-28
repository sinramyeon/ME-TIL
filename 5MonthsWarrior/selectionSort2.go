// 삽입정렬을 이해해보자...
// 삽입정렬을 앞쪽부터 시작하는 방식이라고
func selectionSort2(arr []int) {

  size := len(arr)
  var i, j, min, temp int
  
  for i:=0; i<size-1; i++ {
    min = i
    for j=i+1; j<size; j++ {
      if arr[j] < arr[min] {
        min = j
      }
    }
    // 작은걸 앞으로 
    arr[i], arr[min] = arr[min], arr[i]
  }
}
