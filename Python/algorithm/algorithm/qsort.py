#퀵정렬

def quick_sort(a) :
	n = len(a)

# 남은게 없으면 정렬할 필요 없음
	if n<=1 :
		return a

	pivot = a[-1] # 리스트의 마지막 값을 기준값으로 정했음
	g1 = []
	g2 = []

	# g1에는 그거보다 작은거, g2에는 그거보다 큰걸 담겠음.

	for i in range(0, n-1) : 
		if a[i] < pivot :
			g1.append(a[i])
		if a[i] > pivot :
			g2.append(a[i])

# 중간값 하나를 임의로 정하고 그 앞에거 뒤에건 지알아서 정렬하라고 해서 그 세가지를 더함
	return quick_sort(g1) + [pivot] + quick_sort(g2)



# 리스트를 새로 안 만들고 하는 일반적인 퀵 정렬

def quick_sort_sub(a, start, end) :

	# 리스트 넣고 어디서부터 어디까지 정렬할지 하고 재귀호출로 처리

	if end-start <=0 :
		return # 정렬대상이 하나면 정렬 할필요가 없음

	pivot = a[end] #리스트 맨 마지막값을 우선 정렬 값으로 했음
	i = start
	for j in range (start, end) :
		if a[j] <= pivot : #기준 값보다 작으면
			a[i], a[j] = a[j], a[i] #자리 바꿨음
			i += 1
			# end-1 까지 돈 후
	a[i], a[end] = a[end], a[i]

	quick_sort_sub(a, start, i-1) # 기준값보다 작은 그룹 정렬

	quick_sort_sub(a, i-1, end) # 기준값보다 큰 그룹 정렬 

def quick_sort2(a) :
	quick_sort_sub(a, 0, len(a)-1) # 리스트 처음부터 재귀함수를 호출