func IsMaxHeap(arr []int) bool {

  size := len(arr)
  for i:=0 <= i<=(size-2)/2; i++{
    if 2*i+1<size{
      if arr[i]<arr[2*i+1]{
        return false
      }
    }
    
    if 2*i+2<size{
      if arr[i]<arr[2*i+2]{
        return false
      }
    }
  }
return true
}
