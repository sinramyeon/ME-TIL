def f(n) :
  if n > 1:
    print(n)
    f(n-1)
    
# n부터 1까지 역순으로 출력

def print_to_n(n) :
  print(n, end= ' ')
  if n > 1 :
    print_to_n(n-1)
