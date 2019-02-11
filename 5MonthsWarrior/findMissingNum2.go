func findMissingNum2(data []int, dataRange int) int {
  size := len(data)
  xorSum := 0
  // get the XOR of all the numbers from 1 to dataRange
  // XOR 연산 이용하기
  
  for i:=1; i<=dataRange; i++{
    xorSum ^= i
  }
  
  for i:=0; i<size; i++{
    xorSum ^= data[i]
  }
  
  return xorSum
}
