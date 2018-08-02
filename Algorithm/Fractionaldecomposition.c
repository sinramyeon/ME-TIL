#include <stdio.h>
#include <math.h>

main(){

  // 소인수분해
  int b, c,d,e,quotient, remainder;
  int a[100]; // 약수 저장용

  scanf("%d", &b);

  c = -1;
  d=2;

  while(1){

    e = (int)sqrt(b); // 제곱근을 정수로 변환해서 저장

    while(1){
      if (d>e) {
        d=b;
        break; //d 자체가 소인수라면?
      }
      quotient = b/d;
      remainder = b - quotient*d;

      if(remainder ==0){
        break;
      }else{
        d++;
      }
    }
    c++;
    a[c] = d;
    if(b==d) {
    break;
    }

    
      b = quotient;
  }


  //시발 모르겠음
  
}
