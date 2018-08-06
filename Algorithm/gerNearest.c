#include <stdio.h>
#define FIND 7

main(){

// 7에 가장 가까운 수 구하기

// 7을 빼본다!
// 한자리수에서 구한다면 우선 9에서 뺀걸 최소값(9가 최대값이니)으로 잡고 시작함

int i, j, k, L, m;

int a[10];

i = -1;

do{
  i++;
  scanf("%d", &a[i]);
}while(i<9);

j=9;
for ( k=0; k <=9 ; k++){
  if(a[k] >= FIND){
    L = a[k] - FIND;
  }else{
    L = FIND - a[k];
  }

  if (L<=j){
    j=L;
    m = a[k];
  }

  
}



}
