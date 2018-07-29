#include <stdio.h>

// 1+2+4+7+11... 로 1씩 더해져서 증가되는 수열
// 1+1
// 2+2
// 4+3 ...
// 의 20번째 항

main(){

  int i,j, k;

  i=0;
  j=1;
  k=1;

  do{
      i++;
      j = j+i;
      k = j+k 
  } while(i < 19);

  printf("%d", k);

}
