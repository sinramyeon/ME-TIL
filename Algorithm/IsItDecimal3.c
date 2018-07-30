#include <stdio.h>
#include <math.h>

main(){

  int a, j;
  scanf("%d", &a);
  j = 2;

// 2부터 x의 제곱근까지 숫자로 나누어 떨어지는지 검사
// 나누어 떨어지면 소수가 아님

  while(1){
    if(j <= sqrt(a)){ // 제곱근보다 같거나 작을때
      if(a%j == 0) { // 나머지 0이면 소수 아니고
        printf("소수 아님");
        break; 
      }else{
        j++; // 나머지가 있어야 소수
      }      
    }else{
      printf("소수");
      break;
    }  
  }
}
