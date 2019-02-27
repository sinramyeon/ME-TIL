// 버블정렬보다 조금! 빠름
// 서브어레이 를 두고 함

func InsertionSort(arr []int, comp func(int, int) bool) {
 size := len(arr)
 var temp, i, j, int
 for i=1; i<size; i++ {
  temp = arr[i]
  for j=i; j>0 && comp(arr[j-1], temp); j-- {
    arr[j] = arr[j-1]
    // arr[1] = arr[0] 이란거지
  }
  arr[j] = temp
  // arr[0] = temp
 }
}
