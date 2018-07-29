#include <stdio.h>


// fibo 합계 1+1+2+3+5+8...

main(){

    int hah, cnt, c;
    int a = 1; b = 1;

    hap = 2;
    cnt = 2; // 1+1+2 는 이미 했다치기

    while(1) {
      c = a+b;
      hap = hap+c // 1+1+2

      cnt++;

      if(cnt < 20) {
        a=b;
        b=c;
      }
      else{
        printf("%d", hap);
        break;
      }
      
    }



}
