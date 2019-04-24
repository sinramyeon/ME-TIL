// ‘Stacks can be used to check a program for balanced symbols (such as {}, (), [])’

func IsBalancedParenthesis(expn string) bool {
  stk := new(Stack)
  
  for _, ch := range expn {
    switch ch {
      case '{', '[', '(' :
        stk.Push(ch)
      case '}':
        val := stk.Pop()
        if val != '{' {
          return false
        }
        
      case ']' :
        val := stk.Pop()
        if val != '[' {
          return false
        }
        
      case ')' :
        val := stk.Pop()
        if val! = '(' {
          return false
        }
    }
  }
  
  return stk.IsEmpty()
}
