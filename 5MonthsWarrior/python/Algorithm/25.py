# 등수 구하기

max = 10

a = [75, 25, 35, 46, 76, 45,21, 1, 34, 98]

rank = [0]*max

i = 0
while a[i] != -1 :
  rank[i] = 1
  j = 0
  # 그 전 점수가 더 높으면
  while a[j] != -1 :
    if a[j] > a[i] :
    # 지금 점수 순위를 1번 더함
      rank[i] = rank[i]+1
    j = j+1
  i = i+1
