func (t *Tree) SumAllBT() int {
      return sumAllBT(t.root)
}
 
func sumAllBT(curr *Node) int {
      var sum, leftSum, rightSum int
      if curr == nil {
            return 0
      }
      rightSum = sumAllBT(curr.right)
      leftSum = sumAllBT(curr.left)
      sum = rightSum + leftSum + curr.value
      return sum
}
