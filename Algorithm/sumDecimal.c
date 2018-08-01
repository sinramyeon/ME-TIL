#include <stdio.h>


main(){

//소수의 합 구하기
// 2부터 증가시켜서 소수이면 합에 더하기

int a, hap, k, j;

scanf("%d", &a);

hap=0;
k=2;

while(1){

  j=2;
  while(k%j!=0) {
    j++;
    if(k==j) { // 3을 3으로 나눠 나머지 0이니까 소수
        hap = hap+k;
    }

    if(k < a){ // 그 수보다 작을 때까지 더함 * 3일경우 1, 2, 3 으로 나눠본다.
        k++;
    }
    else{
      printf("%d", hap);
      break;
    }
  }
}
}
