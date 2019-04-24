 func postfixEvalueate(expn string) int {
 
  stk := new(Stack)
  str := strings.Split(expn, "")
  
  for _, tkn : range str {
    value, err := strconv.Atoi(tkn)
    if err == nil {
      stk.Push(value)
    } else {
      num1 := stk.Pop().(int)
      num2 := stk.Pop().(int)
      
      switch tkn {
        case "+" :
          stk.Push(num1+num2)
        case "-" :
          stk.Push(num1-num2)
        case "*" :
          stk.Push(num1*num2)
        case "/" :
          stk.Push(num1/num2)
      }
    }
  }
 return stk.Pop().(int)
 }
