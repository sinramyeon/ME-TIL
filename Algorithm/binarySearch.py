data = [1,2,3,4,5,6,7,8,9,10]


def binary_search(target, data):
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

print(binary_search(3, data))
