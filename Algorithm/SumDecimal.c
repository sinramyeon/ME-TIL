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
    if(k==j) {
        hap = hap+k;
    }

    if(k <a){
        k++;
    }
    else{
      printf("%d", hap);
      break;
    }
  }
}
}
