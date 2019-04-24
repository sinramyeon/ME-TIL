코드의 덩어리를 만든 다음 그것을 호출하고 귀환할 수 있는 구조  > 서브루틴
Go에서는 이런 서브루틴을 **함수** 라고 부른다
Ex) 서브루틴, 함수, 프로시저, 메서드, 호출 가능 객체 등

---

Go는 Call by Value만 지원 -> 함수 내에서 넘겨받은 변수값을 변경해도 함수밖의 변수에는 영향을 주지 않음

```
func AddOne(nums []int) {
  for i:= range nums {
    nums[i]++
  }
}

func ExampleAddOnes() {
   n:= []int{1, 2, 3, 4}
   AddOne(n)
   fmr.Println(n)
   // 2, 3, 4, 5
}
```
---

가변인자 함수

```
func WriteTo(w io.Writer lines.. string) (n int64, err error)

lines := []string{"a", "b", "c"}
WriteTo(w, lines...)
WriteTo(w, "a", "b", "c")
```

---

Go 언어에서 함수는 일급시민(First class citizen)
> 함수 역시 값으로 변수에 담길 수 있고 다른 함수로 넘기거나 돌려받을 수 있음

*함수형 언어*

함수 리터럴(익명함수, 람다) 의 예제

```
func(a, b int) int {
  return a+b
}

func Example_functionLiteral() {
  printHello := func() {
    fmt.Println("Hello?")
  }
  
  printHello()
}
```

---

higher order fuction(고계함수)

함수를 넘겨받거나 반환하는 

```
func ReadFrom(r io.Reader, f func(line string)) error {
  scanner := bufio.NewScanner(r)
  for scanner.Scan() {
    f(scanner.Text())
  }
  
  if err := scanner.Err(); err != nil {
    return err
  }
  
  return nil
}


func ExampleReadFrom() {
  r := strings.NewReader("bill\ntom\njane\n")
  
  err := ReadFrom(r, func(line string) {
    fmt.Println(line)
  })
  
  if err != nil {
    fmt.Println(err)
  }

}
```

---


클로저(closure)

외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드

```
func ExampleReadFrom_append() {
  r:= strings.NewReader("bill\ntom\njane")
  
  var lines []string
  err := ReadFrom(r, func(line string) {
    lines = append(lines, line)
  })
  
  if err != nil{
    fmt.Println(err)
  }
  
  fmt.Println(lines)
  
  
}
```
---



제너레이터

```
func NewIntGenerator() func() int { // 함수를 반환하는 고계 함수
  var next int
  return func() int { // 정수를 반환하는 함수를 반환
    next++ // 클로저임
    return next
  }
}

func ExampleNewIntGenerator() {
  gen := NewIntGenerator()
  fmt.Println(gen(), gen(), gen())
  // 1, 2, 3 이 차례대로 출력
}
```

---


타입 선언

```
type VertexID int
type EdgeID int

func NewVertextIDGenerator() func() VertexID {
  var next VertexID
  return func() VertexID {
    next++
    return next
  }
}

type BinOp func(int, int) int

func OpThreeAndFour(f BinOp) {
  fmt.Println(f(3,4))
}

OpThreeAndFour(func (a, b int) int {
  return a+b
})
```

---

리시버receiver가 있는 함수 메서드 method

> func (recv T) MethodName(p1 T1, p2 T2) R1

OOP의 기본 메서드

```
type VertexID int

func (id VertexID) String() string {
  return fmt.Sprintf("VertextID(%d)", id)
}
```

포인터 리시버로도 응용이 가능하다

```
func ReadFrom(r io.Reader, adjList *[][]int) error

func (adjList *Graph) ReadFrom(r io.Reader) error
```

---

예제

1. 타이머

```
package main

import (
  "fmt"
  
  "time"
)

func CountDown(seconds int) {
  for seconds > 0 {
    fmt.Println(seconds)
    time.Sleep(time.Second)
    seconds--
  }
}

```

블로킹 타이머의 예제. 






