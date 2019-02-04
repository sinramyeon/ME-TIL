//‘Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.’

func MacSubArraySum(data []int) int {
  size := len(data)
  maxSoFar := 0
  maxEndingHere := 0
  for i:= 0; i< size; i++ {
    maxEndingHere = maxEndingHere + data[i]
    if maxEndingHere < 0 {
      maxEndingHere = 0
    }
    
    if maxSoFar < maxEndingHere {
      maxSoFar = maxEndingHere
    }
  
  }

return maxSoFar
}
