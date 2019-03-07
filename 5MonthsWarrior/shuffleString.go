//ABCDE12345를 A1B2C3D4E5로 바꾸자(이런거하지말자..좀...

func shuffle(arr string) string {
  ar = []rune(arr)
  n := len(ar)/2
  count:= 0
  k := 1
  var temp rune
  
  for i:=1; i<n; i=i+2 {
    k = i
    temp = ar[i]
    for true {
      k = (2*k) % (2*n-1)
      temp, ar[k] = ar[k], temp
      count++
      if i==k {
        break
      }
    }
    if count == (2*n-2) {
      break
    }
  }
return string(ar)
}
