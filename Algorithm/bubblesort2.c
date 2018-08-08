#include <stdio.h>

main(){
    // for 문을 이용한 버블정렬
    
    int n, i, j, k;
    int data[10];
    n = -1;
    
    do{
        n++;
        scanf("%d", &data[n]);
    }whie(n<9);
    
    
    for(i=1; i<=9; i++){
        for(j=0; j<=9-i; j++){
            if (dtata[j] > data[j+1]){
                k = data[j];
                data[j] = data[j+1];
                data[j+1] = k;
            }
        }
    } 

}
