func (t *Tree) FloorBST(val int) int {
  curr := t.root
  floor := math.MaxInt32
  
  for curr != nil {
    if curr.value == val {
      floor = curr.value
      break
    } else if curr.value > val {
      curr = curr.left
    } else {
      floor = curr.value
      curr = curr.right
    }
  }
  
  return floor
}

func (t *Tree) CeilBST(val int) int {
  curr := t.root
  ceil := math.MinInt32
  
  for curr != nil {
    if curr.value == val {
      ceil = curr.value
      break;
    } else if curr.value > val {
      ceil = curr.value
      curr = curr.left
    } else {
      curr = curr.right
    }
  }
  return ceil
}
