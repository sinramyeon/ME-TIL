#ifndef FRACTION_H_INCLUDED
#define FRACTION_H_INCLUDED

class Fraction {
	int numerator, denominator; // 분자, 분모
public :
	//생성자
	Fraction(int n = 1, int d = 1) : numerator(n), denominator(d) {}

	Fraction add(const Fraction &f)
	{
		int n = numerator * f.denominator + f.numerator*denominator;
		int d = denominator * f.denominator;
		return Fraction(n, d); // 더하기
	}

	Fraction sub(const Fraction &f) {
		int n = numerator * f.denominator - f.numerator*denominator;
		int d = denominator * f.denominator;
		return Fraction(n, d); // 빼기
	}

	Fraction mul(const Fraction &f){
		int n = numerator * f.numerator;
		int d = denominator * f.denominator;
		return Fraction(n, d); // 곱하기
	}

	Fraction div(const Fraction &f) {
		int n = numerator * f.denominator;
		int d = denominator * f.numerator;
		return Fraction(n, d); // 나누기
	}

	void display() const; // 분수값 출력
};
#endif
