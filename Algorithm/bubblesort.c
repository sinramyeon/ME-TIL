#include <stdio.h>

main(){

  // 버블정렬
  // 첫번째랑 두번째를 두번째랑 세번째를... 식으로 비교해서 교환해서 정렬.
  // 마지막까지 1회전 돌면 가장 큰 자료가 맨 뒤니까 맨 끝 자료는 제외하고 다시 반복
  
  int n, i, j, k;
  int data[10];
  n = -1;
  
  do{
    n++;
    scanf("%d", &data[n]);
  }while(n<9);
  
  
  i = -1;
  
  do{
    i++;
    j = -1;
    do{
      j++;
      if (data[j] > dtata[j+1]) {
        k = data[j];
        data[j] = data[j+1];
        data[j+1] = k;
      }
    }while(j<8-i); // 맨 끝 자료는 제외하려는 몸부림
  }while(i<8);

}
