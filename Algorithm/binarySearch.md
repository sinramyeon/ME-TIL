# 바이너리 서치, 이진탐색

### 바이너리 서치란?

>   BigO :  `O(log N)` 라고도 한다. 오름차순으로 정렬된 자료를 반으로(1/2)나눠 탐색하는 방법이다.

1. 자료 준비
2. 자료의 가운데 지점 찾기
3. 찾고싶은 값이랑 그 가운데 값 비교하기
4. 가운데 값이 저 크면 그 전의것들 보고, 더 작으면 그 뒤에것들 보기
5. 반복!

### 바이너리 서치 풀이

#### 파이썬으로 풀기

```
def binary_search(target, data):
    data.sort() # 오름차순으로 데이터 정렬
    start = 0
    end = len(data) - 1

    while start <= end:
        mid = (start + end) // 2 # 자료 반띵

        if data[mid] == target: # 중간값과 비교하기
            return mid
        elif data[mid] < target:
            start = mid + 1
        else:
            end = mid -1

    return None
```

#### 고랭으로 풀기

```
const arraySize = 10

func BinarySearch(array [arraySize]int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	
	for minIndex <= maxIndex { 
		midIndex := int((maxIndex + minIndex) / 2) // 자료 반띵
		midItem := array[midIndex]
		
		if number == midItem { // 비교
			return midIndex
		}
		
		if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
```

