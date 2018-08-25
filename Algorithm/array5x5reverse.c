#define LENGTH 5
#include <stdio.h>

main(){
    int i, j, k;
    int a[5][5];
    
    k=0;
    
    for(i=0; i<= LENGTH-1; i++){
        for(j=0; j<=LENGTH -1; j++){
            k++;
            a[j][i] = k; 
            
            /*
            1 6
            2 7
            3 8
            4 9
            5 10 (요런모양 배열임)
            */
        }
    }
    
    for(int x=0; x<=LENGTH-1; x++){
        for(int y=0; y<=LENGTH-1; y++){
            printf("%3d", a[x][y]);
        }
    }

}
