// -(1/2)+2/3-3/4+4/5-5/6 ... -99/100

#include <stdio.h>

main(){
  float i, j = 0;

  do{
    i++;
    if( (int) (i/2) == i/2){ //짝수

      j = j+ i / (i+1);
    
    }else{

      j = j - i/(i+1);
    
    }
  
  }while(i<100)

  printf("%f", j);

}
