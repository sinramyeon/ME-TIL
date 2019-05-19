거품 정렬

순차적으로 바로 옆에 있는 데이터와 비교해서 옆의 데이터가 크면 자신과 위치를 변경한다.

1회전을 수행하고 나면 가장 큰 자료가 맨 뒤로 이동하므로 2회전에서는 맨 끝 자료는 정렬에서 제외



```
import random
import time

compare_counter = 0
swap_counter = 0


def bubble_sort(random_list) :
  for start_index in range(len(random_list) - 1 ) :
    for index in range(1, len(random_list) - start_index) :
      if random_list[index-1] > random_list[index] :
        temp = random_list[index-1]
        ranodm_list[index-1] = random_list[index]
        random_list[index] = temp
        
        
if __name__ == '__main__' :
  list = []
  input_n = input("정렬할 데이터 수 : ")
  for i in range(int(input_n)) :
    list.append(random.randint(1, int(intput_n))
  print(list)
```

