//regex matching 알고리즘

func matchExpUtil(exp string, str string, i int,j int) bool {
  if i == len(exp) && j== len(str) {
    return true
  }
  if (i == len(exp) && j!=len(str)) || (i!= len(exp) && j == len(str)){
    return false
  }
  
  if exp[i] == '?' || exp[i] == str[j] {
    return matchExpUtil(exp, str, i+1, j+1)
  }
  
  if exp[i] == '*' {
    return matchExpUtil(exp, str, i+1, j) || matchExpUtil(exp, str, i, j+1) ||  matchExpUtil(exp, str, i+1, j+1)
  }
  return false
} 

func matchExpUtil(exp string, str string) bool {
  return matchExpUtil(exp, str, 0, 0)
}
