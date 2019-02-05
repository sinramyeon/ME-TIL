func Factorial(i int) int {
    if i <=1 {
        return 1
    }
    
    return i*Factorial(i-1)
}
