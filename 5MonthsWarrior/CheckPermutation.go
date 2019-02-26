// check if two lists are permutation of each other
// permutation 이 뭐임
//  순열, 치환

func CheckPermutation(data1 []int, data2 []int) bool {
  size1 := len(data1)
  size2 := len(data2)
  
  if size1 != size2 {
    return false
  }
  
  sort.Ints(data1)
  sort.Ints(data2)
  
  for i:=0; i<size1; i++ {
    if data1[i] != data2[i] {
      return false
    }
  }
  return true
}
