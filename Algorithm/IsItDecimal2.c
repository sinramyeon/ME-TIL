#include <stdio.h>

main(){

  int a, j;
  scanf("%d", &a);

  j=2;

  while(a%j != 0) // j로 나눠서 나머지가 있으면
    j++;
  if(a==j) //
    printf("소수");
  else
    printf("소수 아님"); 

}
