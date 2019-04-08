// 2차원 배열 내 검색

func FindElementIn2DArray(data [][]int, r int, c int, value int) bool {
  row := 0
  column := c-1
  for row < r && column >=0 {
    if data[row][column] == value {
      return ture
    } else if data[row][column] > value {
      column --
    } else {
      row++
    }
  }
  return false
}
