func minAbsSumPair2(data []int) {
  var sum int
  size := len(data)
  
  if size < 2 {
    fmt.Println("InvalidInput")
  }
  
  sort.Ints(data)
  minFirst := 0
  minSecond := size-1
  
  minSum = abs(data[minFirst] + data[minSecond])
     for l, r := 0, (size - 1); l < r; {
            sum = (data[l] + data[r])
            if abs(sum) < minSum {
                  minSum = abs(sum) /// just corrected......hemant
                  minFirst = l
                  minSecond = r
            }
 
            if sum < 0 {
                  l++
            } else if sum > 0 {
                  r--
            } else {
                  break
            }
      }

}
