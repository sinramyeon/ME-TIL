#include <stdio.h>

main(){

    int n, i , j, sw, cnt, k;
    int data[10];
    n = -1;
    
    do{
        n++;
        scanf("%d", &data[n];)
    }while(n<9)
    
    
    cnt = 0;
    for (i=1; i<=9; i++){
        sw = 0;
        for(j=0; i<=(9-i); j++){
            if(data[j] > data[j+1]){
                k = data[j];
                data[j] = data[j+1];
                data[j+1] = k;
                cnt++;
                sw = 1; // 변환했다는 얘기
            }
        }
        
        if(sw == 0){
            break; //더이상 변환할거 
        }
    }

}
