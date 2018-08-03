#include<stdio.h>

void getPrime(int FirstInputNumber) // 소인수 구하기
{
		int i,PrimeNumber,InputNumber;
		int Prime[50]; // 소수 넣을 함수
		int X;
		PrimeNumber=0;
		InputNumber=FirstInputNumber;
	

		for(i=2;i<=InputNumber;i++) // 우선 소수는 2부터 시작하니 2부터 나눠
		{
			if(InputNumber%i==0)
			{
				Prime[PrimeNumber]=i; // 나눠지면 소인수
				PrimeNumber++;
				InputNumber=InputNumber/i; // 2로 나눈 나머지 몫으로 또 나누기 시작하자고
				i=1;
			}
		}
  X=1;

		i=0;
		
		while(FirstInputNumber!=X) // 1까지 나눠지기 전까지
		{
			printf("%d ",Prime[i]);
				X=X*Prime[i]; // 소수 출력
			i++;
		}
}

void main()
{
	int wanted;
	printf("소인수분해하고 싶은 수를 입력하시오\n");
	scanf("%d",&wanted);

	getPrime(wanted);
}
