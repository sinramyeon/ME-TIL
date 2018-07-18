# 그래프
# 모든 연결된 친구 찾기

def print_all_friends(g, start) : # g 그래프 start 자신
  qu = []
  done = set() # 중복 방지용
  
  qu.append(start) # 자신을 큐에 넣고 시작
  done.add(start)
  
  while qu :
    p = qu.pop(0)
    print(p)
    for x in g[p] : # 먼저 꺼낸 애 친구 중
      if x not in done : # 아직 안추가된 사람
        qu.append(x)
        done.add(x)


# 그래프 2
# 모든 연결된 친구 찾기, 촌수 찾기

def print_all_friends(g, start) : # g 그래프 start 자신
  qu = []
  done = set() # 중복 방지용
  
  qu.append(start, 0) # 자신과 친밀도 0을 큐에 넣고 시작
  done.add(start)
  
  while qu :
    (p, d) = qu.pop(0) # 큐에서 사람 이름, 친밀도 꺼냈음
    
    print(p, d)
    for x in g[p] : # 먼저 꺼낸 애 친구 중
      if x not in done : # 아직 안추가된 사람
        qu.append(x, d+1) #친밀도 
        done.add(x)
        
