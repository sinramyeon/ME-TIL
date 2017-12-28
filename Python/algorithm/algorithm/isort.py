# 삽입 정렬

def find_ins_idx(r, v) :
	# r 리스트 안에서 V가 들어갈 위치?

	for i in range(0, len(r)) :
		if v < r[i] :
			# v보다 몇번째 자리에 있는 값이 더 크면
			return i # v는 그 앞에놔
		# 아니면 v가 제일 크므로

		return len(r) # 맨끝에 두겠음



def ins_sort(a) :
	result = []
	while a : # 리스트 안에 값이 남아있을때까지 반복
		value = a.pop(0) # 기존 리스트에서 맨 앞값을 하나 꺼내서
		ins_idx = find_ins_idx(result, value) # 그 값을 어디 들어갈지? 위치 찾기
		result.insert(ins_idx, value) # 그 찾은 자리에 넣음

	return result


# 안에서 직접 위치를 바꿈

def ins_sort(a) :
	n = len(a)
	for i in range (1, n) :
		key = a[i]

		j = i-1 # i 바로 앞

		while j >= 0 and a[j] > key : # key값이 a[j]값보다, 그러니까 앞에있는데보다 작으면
			a[j+1] =a[j] #걔는 한칸 뒤로보냄
			j -= 1
		a[j+1] = key # 찾은 위치에 key를 삽입



"""

[2,4,5,1,3] 이 있을때

n = 5

1부터 4까지

key = a[1] 즉 2
j = 0

a[1]이 a[0]
j = -1

a[0] = 2 > 그대로 둠

key 4
j = 0

2 > 4니까 2, 4

key 5

j = 1
4 > 5가 아니니까
a[2] = 5
2, 4, 5

key 1
j=2

5 > 1이니까
a[3] = a[2] 즉 2 4 5 1
반복해서 2 4 1 5 > 2 1 4 5 > 1 2 4 5... 반복

"""

