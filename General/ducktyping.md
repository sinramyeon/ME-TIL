# duck typing

---


객체지향 개념 중 상속 개념을 생각할때 덕타이핑을 떠올리자.

```
# python

class Duck :
	def sound(self) :
		print("꽥꽥")

class Dog :
	def sound(self) :
		print("멍멍")

def get_sound(animal) :
	animal.sound()

def main() :
	bird = Duck()
	dog = Dog()
	get_sound(bird) # 꽥꽥
	get_sound(dog) # 멍멍
```

A, B 클래스가 가진 함수가 같다면, A, B 클래스 모두 동일한 타입으로 본다.
*이렇게 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식을 덕 타이핑이라 한다.*

```
//Go

package main

import "fmt"

type Duck struct {
	//오리 구조체
}

func (d Duck) quack() {
	fmt.Println("꽥")
}

func (d Duck) feathers() {
	fmt.Println("오리 too much kawaii...")
}

type Person struct {
	// 사람 구조체
}

func (p Person) quack() {
	fmt.Println("꽥(죽음)")
}

func (p Person) feathers(){
	fmt.Println("사람 too less kawaii....")
}

type Quacker interface {
	quack()
	feathers() // 두 메서드를 갖는 인터페이스를 만들었다.
}


func inTheForest(q Quacker){
	q.quack()
	q.feathers()
}

func main(){
	var donald Duck
	var trump Person // 분명 다른 타입의 인스턴스인데

	inTheForest(donald)
	inTheForest(trump)
	// 이러헥 둘 다 호출이 가능하다.
}
```

보통 객체지향 언어에서 지원하며, 파이썬, 루비, 고, C++ 등에서 지원한다.
 동적 언어의 개발속도와 편리성, 컴파일 기반의 정적 언어의 안정감을 둘 다 가질 수는 없는 것인가? -> Golang의 덕 타이핑 지원의 예
