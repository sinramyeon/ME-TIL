// 가격표, 날짜
// 언제 사서 언제 팔아야 최대 이익을 남길 수 있을까?

func maxProfit(stocks []int) {
  size := len(stocks)
  buy := 0
  sell := 0
  curMin := 0
  currProfit := 0
  maxProfit := 0
  
  for i:=0 ; i<size ; i++ {
    if stocks[i] < stocks[curMin] {
      curMin = i
    }
    
    currProfit = stocks[i] - stocks[curMin]
    
    if currProfit > maxProfit {
      buy = curMin
      sell = i
      maxProfit = currProfit
    }
  }
fmt.Println(buy,"에 ", stocks[buy], "어치 삼")
fmt.Println(sell,"에 ", stocks[sell], "어치 팔면")
fmt.Println("최대이윤 : ", maxProfit)
}
