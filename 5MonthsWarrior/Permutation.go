// 숫자로 만들 수 있는 순열 조합

func Permutation(data []int, i int, length int){
  if length == i {
    PrintSlice(data)
    return
  }
  
  for j := i; j <length; j++{
    swap(data, i, j)
    Permutation(data, i+1, length)
    swap(data, i, j)
  }

}

func swap(data []int, x int, y int){

  data[x], data[y] = data[y], data[x]
}
