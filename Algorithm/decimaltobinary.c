#include <stdio.h>

main(){

  int b, bb, c, quotient, remainder, i;
  int a[10];

  scanf("%d", &b);

  bb=b;
  c = -1; // c++시 0부터 배열 시작 용

  do{
    c++;
    quotient = b/2;
    remainder = b - quotient*2;
    a[c] = remainder;

    b = quotient;  
  } while ( quotient != 0); // 2로 계속 나눠

  printf("%d", bb); // 원래 수

  for (i=c; i>=0; i--){
    printf("%d", a[i]);
  }


}
