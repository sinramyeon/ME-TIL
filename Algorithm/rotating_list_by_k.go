/*
‘Given a list, you need to rotate its elements K number of times.
For example, a list [10,20,30,40,50,60] rotate by 2 positions to [30,40,50,60,10,20]’
*/

func RotateArray(data []int, k int) {
  n := len(data)
  ReverseArray(data, 0, k-1)
  
  ReverseArray(data, k, n-1)
  
  ReverseArray(data, 0, n-1)
}

func ReverseArray(data []int, start int, end int) {
  i := start
  j := end
  
  for i<j {
    data[i], data[j] = data[j], data[i]
    i++
    j--
  }

}

/*

Ex)
[10,20,30,40,50,60], 2 positions -> [30,40,50,60,10,20]

RotateArray(array, 2)
k = 2, n = 6
1.

  ReverseArray(data, 0, k-1)
  ReverseArray(data, 0, 2-1)
  
  0, 1
  1 2 3 4 5 6 > 2 1 3 4 5 6
  1, 0
  
  
  ReverseArray(data, k, n-1)
  ReverseArray(data, 2, 6-1)
  
  2, 5
  2 1 3 4 5 6 > 2 1 6 4 5 3
  3, 4
  2 1 6 4 5 3 > 2 1 6 5 4 3
  4, 3
  
  
  ReverseArray(data, 0, n-1)
  ReverseArray(data, 0, 6-1)
  0, 5
  2 1 6 5 4 3 > 3 1 6 5 4 2
  1, 4
  3 1 6 5 4 2 > 3 4 6 5 1 2
  2, 3
  3 4 6 5 1 2 > 3 4 5 6 1 2
  3, 4
  
  > 30 40 50 60 10 20
*/
