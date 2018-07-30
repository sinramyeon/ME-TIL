#include <stdio.h>

//-1+2-4+7-11+16-22 ....

main(){

  int i, j, k, L;

  i = 0;
  j = 1;
  k = -1;
  L = 01;

  do{

    i++;
    j = j+i; // 1+1
    L = L*-1; // -1*-1 = 1
    k = k+ (j*L) // -1 + 2
  
  }while(i<19);

  printf("%d", k);

}
