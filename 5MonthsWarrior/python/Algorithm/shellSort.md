셸 정렬 알고리즘

셸 정렬 알고리즘은 정렬할 데이터를 일정한 구간을 쪼개서 그 구간내 정렬을 한 후 구간을 합쳐서 정렬한다.
일정 그룹을 쪼개서 정렬을 한 후 그룹 내의 정렬이 완료된 후 전체 정렬을 하기 때문에 비교횟수나 데이터의 이동 횟수가 줄어든다.

---


```
import random
import time

def shell_sort(random_list) :
  h = 1
  while h < len(random_list) :
    h = h*3+1
  h = h//3
  
  while h >  0 :
    for i in range(h) :
      start_index = i+h
      
      while start_index < len(random_list) :
        temp = random_list[start_index]
        
        insert_index = start_index
        
        while insert_index > h-1 and random_list[insert_index - h] > temp :
          random_list[insert_index] = random_list[insert_index-h]
          insert_index = insert_index - h
          
        random_list[insert_index] = temp
        start_index = start_index + h
      h = h// 3
      
      
if __name__ == '__main__' :
  list = []
  input_n = input("정렬할 데이터의 수 : ")
  
  for i in range(int(input_n)) :
    list.append(random.randint(1, int(input_n))
  
  print("정렬 전")
  print(list)
  
  start_time = time.time()
  shell_sort(list)
  
  running_time = time.time() - start_time
  print("정렬 후")
  print(list)
  
  print("데이터의 크키 : {}".format(int(input_n)))
  print("실행 시간 : {}".format(running_time))
```
