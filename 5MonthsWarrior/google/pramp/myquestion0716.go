package main

import "fmt"

func FindBusiestPeriod(data [][]int) int {
	// your code goes here

	count := 0
  maxCount := 0
	biggestTime := 0

  m := make(map[int]int)
  
	for _, arr := range data {

		timestamp := arr[0]
		visitors := arr[1]
		enterOrExit := arr[2]
    //fmt.Println("time : ", timestamp)
    
		if enterOrExit != 0 {
			//fmt.Println("enter : ", visitors)
			count = count + visitors
		} else {
			//fmt.Println("out : ", visitors)
			count = count - visitors
		}
    
    m[timestamp] = count
	}
  
  
  for key, val := range m{
    
    if maxCount < count {
      maxCount = val
    }   
    
    if val == maxCount {
      biggestTime = key
      return biggestTime
    }
  }

	return biggestTime
}

func main() {

	// 0 timestamp
	// count of visitors
	// entered 1 exit 0

	data := [][]int{
		{1487799425, 14, 1}, {1487799425, 4, 0}, {1487799425, 2, 0},
		{1487800378, 10, 1}, {1487801478, 18, 0},
		{1487801478, 18, 1},
		{1487901013, 1, 0},
		{1487901211, 7, 1},
		{1487901211, 7, 0},
	}

	fmt.Println(FindBusiestPeriod(data))

}
