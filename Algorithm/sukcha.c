#include <stdio.h>

// 석차 구하기

// 우선 1로 넣고 다음 배열 점수들을 하나씩 봐서 나보다 높으면 석차를 1씩 더한다.

main(){

  int m, n, i, j;
  int jumsu[10], rank[10];


  m = -1;

  do{
    m++;
    scanf("%d", &jumsu[m]);  
  }while(m<9);

  n=9; // 배열 끝
  i=0; // 배열 시작

  while( i<= n ){

    rank[i]=1;
    i++; // 우선다 1등으로넣었다  
  }

  i = 0;
  while( i<= n ){
    j=0;
    while (j<=n){
      
      if(jumsu[i] < jumsu[i+1]){
          rank[i]++
      }
      j++;
    }

    i++;
  }

}
