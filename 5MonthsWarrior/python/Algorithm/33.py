def print_alphabet(end) :
 # 재귀호출 사용
 if end < 'A' or end > 'Z' :
   return -1
 c = ord('A')
 
 while <= ord(end) :
   print(c)
   c = c+1
 
 
 print()
 next = ord(end-1)
 return f(chr(next))
