#include <stdio.h>

main(){

    int i, j, k;
    int a[5][5] = {0};
    
    /*
    1
    32
    654
    ...
    */
    
    k=0;
    for(i=0; i<=4; i++){
    
        for(j=i, j>=0; j--){
            // 반대로 입력
            k++;
            a[i][j] = k;
        }
    }    
    
}
