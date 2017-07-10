# 순차 탐색
## 리스트에 있는 첫 번째 자료부터 하나씩 비교하면서 같은 값이 나오면 그 위치를 결과로 돌려주고 없을 시 -1을 돌려주세요

### 내 답안

my_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1]

def search_list(x) :

    if x in my_list :
        return my_list.index(x)
    else :
        return -1


### 책 답안

# 리스트 a, 찾는 값 x
# 위치를 찾으면 위치를 반환하고, 찾지 못할 시 -1을 반환한다.

def search_list_book(a, x) :

    n = len(a)
    for i in range(0, n) :
        if x == a[i] :
            return i
    return -1


# 연습문제
# 찾는 값이 나오는 모든 위치를 리스트로 돌려주는 탐색 알고리즘을 만드세요.
# 찾는 값이 리스트에 없을 시 빈 리스트를 반환하세요.

def search_list_list(new_list, search_num) :

    empty_list = []

    for x in range(0, len(new_list)) :
        if search_num == new_list[x] :
            empty_list.append(x)

    return empty_list


# 학생 번호와 이름이 리스트로 주어졌을 때, 학생 번호를 입력하면 학생 이름을 출력해 주세요
# 학생 번호에 해당하는 이름이 없을 때는 ? 를 돌려주세요.

stu_no = [39, 14, 67, 105]
stu_name = ["Justin", "John", "Mike", "Summer"]

def student(no) :

    try :
        if stu_name[no] :
            print(stu_name[no])
    except :
        print("?")


# 책 풀이

##s_no 학생 번호 리스트
##s_name 학생 이름 리스트
###find_no 찾는 학생 번호

def get_name(s_no, s_name, find_no) :
    n = len(s_no)
    for i in range(0, n) :
        if find_no == s_no[i] :
            return s_name[i]

    return "?"