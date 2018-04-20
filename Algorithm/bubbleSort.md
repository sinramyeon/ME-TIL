# 버블 정렬(거품 정렬)

## 버블 정렬이란

![](http://ejklike.github.io/assets/20170301/bubblesort.gif)

이웃한 두 값을 비교하여 정렬할 때, 큰 값이 오론쪽으로 이동하는 과정이 반복되면서 비교했던 모든 값들의 최댓값이 맨 오른쪽으로 옮겨지는 정렬 방법.

귀여운 꿀벌그림을 보다시피 조금 느리다.

만약 데이터 정렬이 잘 되어 있다면 조금 빠르게 멈추므로 데이터 정렬 여부를 파악하기 위한 알고리즘으로 사용될 수 있다.

[유튜브](https://youtu.be/lyZQPjUT5B4) 에 거품 정렬을 춤으로 비교한 영상도 있다. 너무 귀엽다.

### 파이썬으로 구현한 버블 정렬

```
def bubbleSort(x):
	length = len(x)-1
	for i in range(length): #이 배열 끝까지
		for j in range(length-i): # 이 배열 마지막 열 전까지
			if x[j] > x[j+1]: # 옆에거보다 더 크면
				x[j], x[j+1] = x[j+1], x[j] # 마지막거랑 자리를 옮긴다.
	return x
```

좀 더 성능을 향상할 수 있다면, 정렬이 다 끝났다면 더는 반복하지 않게 좀 손볼 수 있다.

```
def optimized_bubble_sort(L):
    swapped = True
    while swapped:
        swapped = False # for 에서 정렬 다 될시 False로 멈춤
        for i in range(len(L)-1):
            if L[i] > L[i+1]:
                temp = L[i+1]
                L[i+1] = L[i]
                L[i] = temp
                swapped = True
```


### 고로 구현한 버블 정렬

```
package BubbleSort

  func BubbleSort(array []int) {

    swapCount := 1

      for swapCount > 0 {

        swapCount = 0

          for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {

            if array[itemIndex] > array[itemIndex+1] {

               array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]

               swapCount += 1

             }

           }

      }

}
```
