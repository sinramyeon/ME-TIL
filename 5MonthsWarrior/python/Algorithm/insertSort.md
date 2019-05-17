삽입 정렬은 순차적으로 정렬하면서 현재 값을 정렬된 값들과 비교해서 위치로 삽입하는 방식

```
import time

compare_counter = 0
swap_counter = 0


def insertion_sort(my_list) :
  global compare_counter, swap_counter
  
  my_list.insert(0, -1)
  
  for s_idx in range(2, len(my_list)) :
    temp = my_list[s_idx]
    
    ins_idx = s_idx
    compare_counter += 1
    
    while my_list[ins_idx-1] > temp :
      swap_counter += 1
      my_list[ins_idx] = my_list[ins_idx-1]
      ins_idx = ins_idx - 1
      
    my_list[ins_idx] = temp
  del my_list[0]
  
  
if __name__ == '__main__' :
  list = []
  input_n = input("정렬할 데이터 수 : ")
  for i in range (int(input_n)) :
    list.append(random.randint(1,int(input_n)))
```

정렬할 데이터의 수가 많을수록 삽입 정렬이 선택 정렬보다 빠르다.
정렬이 어느정도 되어있을 시 유용하다.
