#include <stdio.h>

main(){


// 배수의 개수와 합 구하기

int cnt, hap, quotient, remainder, i;
cnt =0, hap = 0;


for(i=1; i<=100; i++){
  quotient = i/5;
  remainder = i-quotient*5;

  if(remainder == 0){
    cnt++;
    hap = hap+i;
  }
}


}
