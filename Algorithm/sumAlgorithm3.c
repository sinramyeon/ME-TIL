// (-1) * 2 * (-3) ...*100

#include <stdio.h>

main(){

  int i;
  double j;

  i=0;
  j= 1;

  do{

  i++;
  if (i%2 ==0 ){
    j = j*i;
  }else{
    j = j*i*(-1);
  }
  
  }while(i<100)

printf("%11.4e", j);
}
