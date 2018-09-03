#include <stdio.h>

/*
12345
 678
  9
101112
1314151617
*/

main(){

    int x, m, L, i, j, k;
    scanf("%d", &x);
    int a[x][x]; // 배열생성
    for(i=0; i<x; i++){
        for(j=0; j<x; j++){
            a[i][j]=0;
        }
    }
    
    k = 0;
    L = x;
    m = x/2;
    
    for(i=0; i<=m; i++){
        L--;
        for (j=i; j<=L; j++) {
            // 1씩 줄어드는 L
            k++;
            a[i][j] = k;
        }
    }
    
    for (i=m+1; i<=x-1; i++){
     // 중간 이후론 늘어남
     L--;
     for(j=L; j<=i; j++){
        k++;
        a[i][j] = k;
     }
    }


}
