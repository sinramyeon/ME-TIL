#include <stdio.h>

//1!+2!+3!+4!....10!

main(){
  int i =1, k=1, j=1;

  do{
    i++; // 2
    j = j*i; // 1*2 = 2!
    k = k+j; // 1 + 2! ...(반복)
  }while(i<10)

  printf("%d", k);

}
