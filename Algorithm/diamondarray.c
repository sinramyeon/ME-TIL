#include <stdio.h>

main(){

    int i, j;
    int a[5][5] = {0};
    int k = 0;
    int s=2, e=2;
    // 0 1 2의 중간 2라는거임
    
    for (i=0; i<=4; i++){
        for (j=s; j<=e; j++){
            k++;
            a[i][j]=k;
        }
        
        if (i>=2) {
            // 행 중간위치일시
            s++;
            e--; // 다이아몬드꼴 12345니까 이제 줄임
        }else{
            s--;
            e++;
        }
    
    }
    
    
}
