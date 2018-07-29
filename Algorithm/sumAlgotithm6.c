#include <stdio.h>

//-1+2-4+7-11+16-22 ....

main(){

  int i, j, k L;

  i = 0;
  j = 1;
  k = -1;
  L = 01;

  do{

    i++;
    j = j+i;
    L = L*-1; // -1*-1 = 1
    k = k+ (j*L)
  
  }while(i<19);

  printf("%d", k);

}
