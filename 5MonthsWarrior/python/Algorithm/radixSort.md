기수 정렬 알고리즘

자릿수 정렬 방식이다. 우선 1의 자리수로 정렬한 후 10의자리수로 다시 정렬하는 식이다.
자릿수 정렬이므로 데이터가 이동하거나 비교하는 횟수가 거의 없다.

```
from math import log10
from random import randint

def get_digit(number, base, pos) :
  return (number//base**post) % base


def prefix_sum(array) :
  for i in range(1, len(array)) :
    array[i] = array[i] + array[i-1]
  return array



def radixsort(l, base=10) :
  passes = int(log10(max(1))+1)
  output = [0] * len(1)
  
  for pos in range(passes) :
    count [0] * base
    
    for i in l :
      digit = get_digit(i, base, pos)
      count[digit] += 1
    
    count = prefix_sum(count)
    
    
    for i in reversed(l):
      digit = get_digit(i, base, pos)
      count[digit] -= 1
      new_post = count[gidit]
      output[new_pos] = i
    
    
    l = list(output)
  return output
  
```

