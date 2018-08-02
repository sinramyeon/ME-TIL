# include <stdio.h>

// 소수의 개수를 구해보자.
// 소수의 배수는 소수가 아니라는 원리를 이용한다.

// 소수가 아닌 자리에 0을 기억시키자

main(){
  int k, i, j, m;
  int a[99];

  k = 1;
  do{
    k++;
    a[k-2] = k; // 0부터 시작하자
  
  }while(k<100)

  i = -1, j = 0;

  while(1){
    i++; // 1는 0부터 시작
    if(i>98) {
      printf("%d", j);
      break;
    }

    if(a[i] == 0) {
      continue; // 소수가 아닐시 컨티뉴
    }

    j++;
    m = i;

    while (1) {
        m+= a[i];
        
        if(m>99) {
          break;
        }

        a[m] = 0;
    }

  
  }
  



}
