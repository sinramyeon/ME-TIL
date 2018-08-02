#include <stdio.h>

main(){

  // 최대공약수, 최소공배수 구하기
  // 큰 수를 작은 수로 나눠서 나머지가 0이면 작은수가 최대 공약수고
  // 두 수 곱을 최대공약수로 나누면 최소공배수임
  // 유클리드 호제법 이용

  // 최대공약수는 영어로 하면 Greatest Common Measure인데, 첫 글자를 따서 알파벳 G로, 최소공배수는 Least Common Multiple
  // 몫과 나머지 quotient-and-remainder

  int a, b, big, small, quotient, remainder, gcm, lcm;

  scanf("%d %d", &a, &b);

  if( a>=b ){
      big = a;
      small = b;
  }else{
      big = b;
      small= a;
  }

  while(1){
    quotient = big/small;
    remainder = big - quotient * small;
    if( remainder == 0 ){
        gcm = small;
        lcm = a*b / gcm;
        printf("%d %d", gcm, lcm);
        break;
    }

    big = small;
    small = remainder;
  }
}
