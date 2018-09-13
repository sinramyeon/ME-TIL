#include <iostream>
#include "Fraction.h"

using namespace std;

void Fraction::display() const { // 인쇄하기
	cout << numerator;
	cout << "/";
	cout << denominator;
}
