# 두번 이상 나온 이름 찾기
# 딕셔너리 활용

def find_same_name(a) :
  name_dict= []
  for name in a :
    if name in name_dict :
      name_dict[name] += 1
    else :
      name_dict[name] = 1
      
  result = set() # 만든 딕셔너리 중 두번이상 나온걸 추가
  
  for name in name_dict :
    if name_dict[name] >=2 :
      result.add(name)
      
      
  return result
