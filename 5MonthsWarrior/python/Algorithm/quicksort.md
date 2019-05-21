퀵 정렬

기준이 되는 자료를 뽑아 그 자료를 기준으로 왼/오른쪽을 나눠 자료와 비교해 순서를 정렬한다.
나눠진 왼/오른쪽에서 다시 기준이 되는 자료를 뽑아 반복한다.

```
import random
import time

compare_counter = 0
swap_counter = 0

def swap(x, i, j) :
  x[i], x[j] = x[j], x[i]
  
def pivotFirst(x, lmark, rmakr) :
  pivot_val = x[lmark]
  pivot_idx = lmark
  while lmark <= rmark :
    while lmark <= rmark and x[lmark] <= pivot_val :
      lmark += 1
    while lmark <=rmark and x[lamrk] >= pivot_val :
      rmark -= 1
    if lmark <= rmark :
      swap(x, lamrk, rmark)
      lmark += 1
      rmark -= 1
      
  swap(x, pivot_idx, rmark)
  return rmark
  

def quickSort(x, pivotMethod=pivotFirst) :
  def _qsort(x, first, last) :
    if first < last :
      splitpoint = pivotMethod(x, first, last)
      _qsort(x, first, splitpoint-1)
      _qsort(x, splitpoint+1, last)
      _qsort(x, 0, len(x)-1)
      
```

divde and conquer 분할 정복 알고리즘이라고도 한다.
