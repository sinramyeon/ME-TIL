func findMissingNum2(data []int, dataRange int) int {
  size := len(data)
  xorSum := 0
  // get the XOR of all the numbers from 1 to dataRange
  // XOR 연산 이용하기
  // XOR 연산은 두 값의 각 자릿수를 비교해,값이 0으로 같으면 0, 값이 1로 같으면 0, 다르면 1을 계산한다.

  for i:=1; i< = dataRange; i++{
    xorSum ^= i
  }
  
  for i:=0; i<size; i++{
    xorSum ^= data[i]
  }
  
  return xorSum
}
