# 알파벳 하나씩 줄여 출력
# abcd...z
# abcd...y

#이렇게


def f() :
  a = ord('A') # 아스키코드로 변환
  whiel a <= ord('Z') :
    b = ord('A')
    while b <= (ord('Z')-(a-65)) :
      print(b)
      b=b+1
    print()
    a = a+1
