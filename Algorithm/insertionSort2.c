#include <stdio.h>

main(){

  // 삽입 정렬
  // if문을 이용해보자.

  int i, k, key;
  int a[20] = {9,23,5,6,10,13,19,17,1,20,15}

  i = 1;

  do{
    key = a[i];
    k = i=1;
    do{
      if(key < a[k]){
        a[k+1] = a[k];
        k--;
      }else{
        break;
      }    
    }(while k>=0); // 배열이 0부터 시작하므로

    a[k+1] = key;
    i++  
  }while(i<=len(a)); 


}
