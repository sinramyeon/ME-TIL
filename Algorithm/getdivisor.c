#include <stdio.h>

main(){

  // 약수 구하기
  // 나눠서 나머지가 0이 되게 하는 제수들
  // 몫과 나머지 quotient-and-remainder
  int b, quotient, remainder, i;

  int a[100]; // 약수 저장용 배열

  scanf("%d", &b);

  int c=0, d=-1;

  while(1){

    c++;
    if(c <= b){
      quotient = b/c;
      remainder = b- quotient*c;

      if(remainder == 0) {
        d++; // 0부터 시작하게 했음
        a[d] = c; // 약수 저장
      }
      
    }else{
      print("%d", b);
      for(i=0, i<=d, i++){
        printf("%d", a[i]); // 약수 출력
      }
      break;    
    }

  
  }

  
}
