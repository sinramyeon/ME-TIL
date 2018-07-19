#include <stdio.h>

main(){

  int i, j;
  i = 0, j = 0;
  do {
    i++;
    j = j+i;  
  }while (i < 100);
  
  printf("%d %d", i, j);
}
