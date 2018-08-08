#include <stdio.h>

main(){
 // 선택 정렬
 // 첫 번째자료를 두 번째 자료부터 마지막까지 비교해서 가장 작은걸 첫번째에 놓자
 // 이후 반복함(두 번째 자료부터...)

 int m,i,j,k,x;
 int data[10];
 m = -1;

 do{
  m++;
  scanf("%d", &data[m]) 
 }while (m<9) // 배열 입력받기

i = -1;
do{
  i++;
  j=i;
  do{
    j++;
    if(data[i]>data[j]){
      k = data[i]; // 교환
      data[i] = data[j];
      data[j] = k;
    }
  }while(j<9);
}while(i<8);

for(x=0; x<10; x++){
  printf("%d", data[x]);
}

}
