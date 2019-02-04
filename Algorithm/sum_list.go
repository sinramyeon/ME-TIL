//write the method that will return the sum of all the elements of the integer list.
//given list as an input argument

func SumArray(data []int) int {

size = len(data)
total := 0
for index:=0; index < size; index++{
  total = total + data[index]
}
return total
}
