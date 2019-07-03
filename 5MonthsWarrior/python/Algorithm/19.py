# 반복분을 사용해서 조합 구하기

n개중 r개를 뽑는 조합 구하는법?
n!/r!(n-r)!

def combi(n, r) :
  i = 1
  p = 1
  while i <= r :
    p = p*(n-i+1)
    i = i+1
  return p #의미를 모르겠음
