// 1-2+3-4+5.......+99-100


#includ <stdio.h>
main(){

  int i, j, sw;

  i = j = 0;
  sw = 0;

  do{

    i++;
    if (sw ==0) {
      j = j+i;
      sw = 1;
    }
    else {
      j = j-i;
      sw = 0;  
    }
      
  }while (i < 100);

  printf("%d", j);

}
