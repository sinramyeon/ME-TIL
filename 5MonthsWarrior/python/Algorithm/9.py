def check_prime(n) :
  i = 2
  while i < n :
    if n%i == 0 :
      break
    i = i+1
  if i == n :
    print("소수")
  else :
    print("합성수")



if __name__ == '__main__" :
  i = 2
  while i<= 20 :
    check_prime(i)
    i = i +1
