#include <stdio.h>

main(){
    int j, L, h, m;
    int data[2][10] = {
        {1,2,3,4,5,6,7,8,9,10},
        {100,66,25,88,90,65,34,56,98,13}
    };
    
    scanf("d", &j);
    L= 0; h=9;
    
    while(1){
        if(L<=h){
            m = (L+h) / 2;
            if(j==data[0][m]) {
                printf("%d %d", j, data[1][m]);
                break;
            }
            
            if(j<data[0][m]){
                h = m-1;
            }else{
                l = m+1;
            }
            
        }else{
            printf("%d Not found", j);
            break;
        }
    }
    
    
}
