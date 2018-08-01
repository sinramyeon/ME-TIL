//1부터 입력한 수까지 소수의 합
#include <stdio.h>

int prime( int a );

int main( )
{
	int input;
	int i;
	int sum;

	printf( "0을 입력하면 종료\n\n" );

	while(1)
	{
		sum = 0;

		//입력
		printf("양의 정수를 입력 : ");
		scanf( "%d", &input );
		if( input==0 )		return 0;

		//입력한 수까지 반복하고, 소수가 리턴되면서 누적
		for( i=2 ; i<=input ; i++ )
		{
			sum+=prime( i );
		}

		printf( "%d까지의 소수의 합 : %d\n\n\n", input,  sum );
	}

	return 0;
}

//소수를 구하는 함수
int prime( int a )
{
	int j;

	for(j=2 ; j<=a ; j++) // 2부터 시작해서 나눠봄
	{
		if(a%j == 0)
		{
			if(a == j) return a; // 3이 3으로 나눠져서 나머지 0이면 소수
			if(a != j) return 0; // 4가 2로 나눠져서 나머지 0이면 소수가 아닌경우!
		}
	}
}
