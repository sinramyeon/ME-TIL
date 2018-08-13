#include <stdio.h>

// 석차 구하기

// 우선 1로 넣고 다음 배열 점수들을 하나씩 봐서 나보다 높으면 석차를 1씩 더한다.
// 바로출력 버전!
main(){

int i, j, r;
int kuk[10], mat[10], hap[10]; // 국어 수학 총점

i=-1;

do{
  i++;
  scanf("%d %d", &kuk[i], &mat[i]);
  hap[i] = kuk[i]+mat[i];
}while(i<9);

for(i=0; i<=9; i++){
  if(hap[i] != 0){
    // 빵점은 출력안함
    r=1; // 우선 일등으로둠

    for(j=0; j<=9; j++){
      if(hap[i] < hap[j]) {
        r++; // 남이더크면 석차증가
      }
    }
  }

}


}
