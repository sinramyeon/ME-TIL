#include <iostream>
#include "Fraction.h"
using namespace std;

int main() {
	Fraction f1(1, 2);
	Fraction f2(1, 3);
	Fraction f3 = f1.add(f2);

	f1.display();
	cout << "+";
	f2.display();
	cout << "=";
	f3.display();
	cout << endl;

	Fraction f4 = f1.sub(f2);


	f1.display();
	cout << "-";
	f2.display();
	cout << "=";
	f4.display();
	cout << endl;

	Fraction f5 = f1.mul(f2);
	f1.display();
	cout << "* 1/3 = ";
	f5.display();
	cout << endl;


	Fraction f6 = f1.div(f2);
	f1.display();
	cout << "나누기 1/3 = ";
	f6.display();
	cout << endl;

	return 0;

}
