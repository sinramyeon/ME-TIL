#include <stdio.h>

main(){

  int j, L, h, m;
  int data[10] = {8, 15, 35, 55, 60, 61, 70, 80, 92, 99};

  scanf("%d", &j);
  L = 0; // 배열 0부터 시작
  h =9; // 1~10까지

  while(1){

    if(L<=h){
      m = (L+h) / 2;
      if(j==data[m]){
        printf("%d %d", j, m+1);
      }
      if (j<data[m]){
        h = m-1;
      }else{
        h = m+1;
      }
    }else{
      printf("%d NOT FOUND", j);
      break;
    }
  
  }

}
