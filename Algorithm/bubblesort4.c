#include <stdio.h>

main(){

  int n, i, shift, left, right, buf;
  int d[10] = {8,5,6,2,1,3,2,4,7,9,10}
  
  // 좌우로 번갈아 버블정렬 수행
  // 시작과 끝 위치를 기억할 변수
  // 가장큰게 맨 뒤로가면 2회전에서는 맨오른쪽거 제외하고 왼쪽방향으로 가는거지 3회전은 맨왼쪽거 제외하고 반복하는 식으로
  
  n = 9;
  left = 0; //맨왼쪽
  right = 9; //맨오른쪽
  
  while(left < right){
    for(i=left; i<=right-1; i++){
      if(d[i]>d[i+1]){
        buf = d[i];
        d[i] = d[i+1];
        d[i+1]=buf;
        shift = i; // 오른쪽 마지막 기억
      }
    }
    
    right = shift;
    for(i=right; i>= left+1; i--){
       if(d[i-1] > d[i]) // 왼쪽으로 돌고있음{
          buf = d[i-1];
          d[i-1] = d[i];
          d[i] = buf;
          shift = i;
       }
    }
  left=shift;
  }
  
  
}
