#include <stdio.h>

main(){
    int i,j, k;
    int a[5][5];
    
    k=0;
    for(i=0; i<=4; i++){
    
        for(j=0; j<=4; j++){
            k++;
            a[i][j] = k;
        } 
        
    }
    
    for(int x=0; x<=4; x++){
        for(int y=0; y<=4; y++){
            printf("%3d", a[x][y]);
        }
    }

}
