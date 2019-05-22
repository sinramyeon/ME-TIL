# 123 을 321 로 역순 출력하라
# 자릿수 계산
def solve(n) :
  if n == 0 :
    return 0
  print(n%10, end=' ')
  solve(n//10) // 1, 10, 100, 1000... 의 자리 순서대로 출력하는 방식
