#include <stdio.h>

main(){

// 최대, 최소값을 제외한 평점평균

int m, min, max, hap, avg, i;
int a[7];

m = -1;

  do{
    m++;
    scanf("%d", &a[m]);
  }while(m<6);

  min = a[0];
  max = a[0];
  hap = a[0];
  i=0;
  while (i<6){
    i++;
    hap = hap+a[i];
    if(a[i] < min) {
      min= a[i];
    }
    if(a[i] > max){
      max = a[i];
      }    
  }

  hap = hap-min-max;
  avg = hap/5;


}
