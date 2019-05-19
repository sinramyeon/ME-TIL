#공배수 계산하기


def check_common(n) :
  i = 1
  while i <=n :
    if (i %3 == 0) and (i%5 == 0) :
      print("%d"%(i))
    i = i+1


# 3이랑 5의 배수 구하기
