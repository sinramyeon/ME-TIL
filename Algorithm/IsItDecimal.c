# include <stdio.h>

main(){
  int a, i, j;

  scanf("%d", &a); // 10진수 입력

  i = a-1; // 그거보다 작은 수부터 나눠보기
  j = 2;

  while(1) {
    if(j <= i) {    
      if (a % j == 0) { 
        printf("소수 아님");
        break;
      }else{
        j++; // j값 누적해서 2부터 나누기
      }
   }else{
      printf("소수");
      break;
    }
  }
}
