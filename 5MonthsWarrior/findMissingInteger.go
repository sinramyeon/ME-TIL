// list of integers to find missing number in the list!!

func FindMissing(arr []int, start int, end int) (int, bool) {
  hs := make(Set)
  for _, i: = range arr{
    hs.Add(i)
  }
  
  for curr:= start; curr <=end ; curr++ {
    if hs.Find(curr) == false{
      return curr,true
    }
  }
  return 0, false
}
