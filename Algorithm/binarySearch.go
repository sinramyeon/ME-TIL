package main

import (
	"fmt"
)


const arraySize = 10

func BinarySearch(array [arraySize]int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]
		
		if number == midItem {
			return midIndex
		}
		
		if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}

func main() {
   array := [arraySize]int{1,2,3,4,5,6,7,8,9,10}

   fmt.Println(BinarySearch(array, 3))
}
