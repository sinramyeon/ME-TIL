선택 정렬

데이터 끝까지 훑어서 가장 작은 값을 찾아서 첫번째랑 바꾸고...를 반복

```
def selected_sort(random_list) :
  for sel in range(len(random_list)-1) :
    min = random_list[sel]
    minindex = sel
    
    for step in range(sel+1, len(random_list)) :
      if min > random_list[step] :
        min = random_list[step]
        minindex = step
    
    
    random_list[minindex] = random_list[sel]
    random_list[sel] = min



if __name__ == '__main__' :
  list = []
  for i in range(10) :
    list.append(random.randint(1,10))
  print(list)
  
  selected_sort(list)
  print(list)
    
```

