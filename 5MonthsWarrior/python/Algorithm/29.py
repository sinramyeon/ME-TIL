# 하 노 이

def hanoi(n, a, b, c) :
  if n > 0 :
    hanoi(n-1, a, c, b)
    print("%d번 원반을 %c에서 %c로 옮김"%(n, a, b))
    hanoi(n-1, c, b, a)
