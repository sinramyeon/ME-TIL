Go는 상속이나 제네릭 등이 없다. 조금 생경하게 느껴질 수도 있다. 이것 때문에 Go에서는 부모-자식간의 객체 관계를 맺기가 조금 어려워 질 수도 있다.

아래의 코드를 보자.

자식
package child

import "Parent"

type Child struct {
  parent *Parent
}

func (child *Child) PrintParentMessage() {
  child.parent.PrintMessage()
}

func NewChild(parent *Parent) *Child {
  return &Child{parent: parent }
}
부모
package parent

import (
  "fmt"
  "child"
)

type Parent struct {
  message string
}

func (parent *Parent) PrintMessage() {
  fmt.Println(parent.message)
}

func (parent *Parent) CreateNewChild() *child.Child {
  return child.NewChild(parent)
}

func NewParent() *Parent {
  return &Parent{message: "Hello World"}
}
메인함수에서 이용하기
package main

import (
  "parent"
)

func main() {
  p := parent.NewParent()
  c := p.CreateNewChild()
  c.PrintParentMessage()
}


여기서는 Import cycle not allowed' error: 가 발생하면서 빌드가 되지 않는다.



Import cycle이란?

Cross-refering packages. 각 패키지들이 자기 서로 참조하는것을 뜻한다. 고에서는 허용되지 않는다.





이것은 그지같은설계를 막는 하나의 방법이지만, 처음 맞닥뜨리면 귀찮기 짝이 없다.



Import cycle을 어떻게 해결하지?

인터페이스로 해결

package child

type IParent interface {
  PrintMessage()
}

type Child struct {
  parent IParent
}

func (child *Child) PrintParentMessage() {
  child.parent.PrintMessage()
}

func NewChild(parent IParent) *Child {
  return &Child{parent: parent }
}
package parent

import (
  "fmt"
  "child"
)

type Parent struct {
  message string
}

func (parent *Parent) PrintMessage() {
  fmt.Println(parent.message)
}

func (parent *Parent) CreateNewChild() *child.Child {
  return child.NewChild(child.IParent(parent))
}

func NewParent() *Parent {
  return &Parent{message: "Hello World"}
}


다시 부모, 자식 코드를 고쳤고, 자식은 부모를 임포트하지 않고 있다.



package main

import (
  "parent"
)

func main() {
  p := parent.NewParent()
  c := p.CreateNewChild()
  c.PrintParentMessage()
}
이렇게 임포트 에러를 피할 수 있다.



Duck typing으로 해결


A -> B -> C -> A 식으로 돌고도는 패키지 의존성을 만들어 버렸다면, InterfaceA를 만들던가 설계를 다시 하는 것이 좋겠다.



패턴 공부하기

클린 아키텍쳐란

Go에 적용하기


Go 에서 확인해볼 수 있는 의존성 피하기의 예제

가장 대표적으로는 io모듈이 os모듈과 분리되어 있다.

파일을 다루는데도 os.File클래스는 전혀 이용하지 않는 것을 확인할 수 있다. 

io.Writer랑 *os.File이 인터페이스로 구성되어 있다 한다.(확인해보지는 않음)

패키지 의존성 확인하기


go list -f '{{join .Deps "\n"}}' 로 현재 디렉토리에서 가져다 쓰는 패키지들을 확인해 볼 수 있다.
