#include <stdio.h>

main(){

    int i, j, k;
       int a[5][5] = {0}; // 모든 요소 0으로 초기화
       // 아래있는 직각삼각형을 배열에 그려보자
       
       k = 0;
       for (i=0; i<=4; i++){
            for(j=0; j<=i; j++){
                k++;
                a[i][j] = k;
            }
       }
       
       for(int x =0; x<=4; x++){
            for(int y=0; t<=4; y++){
                printf("%3d", a[x][y]);
                printf("\n");
            }
       }
}
