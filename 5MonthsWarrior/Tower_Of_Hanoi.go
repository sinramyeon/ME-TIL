// 하노이의 탑~

func ToHUtil(num int, from string, to string, temp string){
  if num < 1 {
    return
  }
  
  TOHUtil(num-1, from, temp, to)
  fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
  TOHUtil(num-1, temp, to, from)
}

func TowersOfHanoi(num int) {
  ToHUtil(num, "A", "C", "B")
}
