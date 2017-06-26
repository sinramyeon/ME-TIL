# 동명이인 찾기
# n명의 사람 이름 중에서 같은 이름을 찾아 집합으로 돌려주자!


name = ["Tom", "Brandon", "Alice", "Rihanna", "Tom", "Jerry", "Mike"]

# 내 풀이

def my_name(a) :

    myset = set()

    for x in a :

        if a.count(x) >= 2 :
            print(x)
            myset.add(x)

    return myset


# 다른 풀이

def find_name(a) :

    return set([x for x in a if a.count(x) >= 2])

# 책 풀이

def find_same_name(a) :
    n = len(a)

    result = set()

    for i in range (0, n-1) : #0부터 n-2까지 반복
        for j in range(i+1, n) : #i+1부터 n까지 반복
            if a[i] == a[j] :
                result.add(a[i])

    return result


# n명 중 두 명을 뽑아 짝을 짓는다고 할 때, 짝을 지을 수 있는 모든 조합을 출력하세요
# 예를들어 tom, jerry, mike 일시 tom-jerry, jerry-mike, tom-mike 로 출력하세요.

# 내 답. . .이 책 답하고 같았음

def my_frnd(a) :
    n = len(a)

    result = set()

    for i in range(0, n - 1):  # 0부터 n-2까지 반복
        for j in range(i + 1, n):  # i+1부터 n까지 반복
            if a[i] != a[j]:
                print(a[i]+"-"+a[j])


my_frnd(name)