자바스크립트에서 처음 클로저를 배웠을 때를 기억해 보자.

고에도 이 클로저 문법이 있다.



클로저는 독립적인 자유 변수를 참조하는 함수들이다.


문법

function init(){
	var name = "Mozilla" # 지역변수
	function displayName() { # 내부 함수인 클로저
	alert(name)
	}
	displayName();
}

init()


init이 생성한 지역 변수 name과 함수 displayName(). displayName()은 부모 함수에서 선언된 변수 name을 참조



function makeFunc() {
  var name = "Mozilla";
  function displayName() {
    alert(name);
  }
  return displayName;
}

var myFunc = makeFunc();
myFunc();


displayName() 안 내부 함수 실행 전 외부 함수로부터 반환



function makeAdder(x) {
  return function(y) {
    return x + y;
  };
}

var add5 = makeAdder(5);
var add10 = makeAdder(10);

console.log(add5(2));  // 7
console.log(add10(2)); // 12


add5, add10이 클로저 > 함수 본문을 공유하지만 서로 다른 변수를 기억



자바스크립트의 모든 함수는 클로저이다.
그럼, Go에서의 클로저 활용법과 scope에 대해서 알아보자.



고는 익명함수를 지원한다. 그리고 이 익명함수로 클로저 형태를 구성할 수 있다.

익명함수는 일반 함수와 달리 함수를 정의할 때 이름이 없다. 함수가 들어있는 변수를 함수처럼 호출해서 사용한다.



클로저는 함수 외부에서 선언된 변수를 안에서 참조하는 특수한 형태의 익명함수이다.



보통 클로저를 구성할 때 함수 안에 이름이 없는 익명 함수를 사용합니다. 이때 사용되는 익명 함수를 람다 함수라고 합니다. 클로저와 람다 함수를 혼동하기 쉬운데 둘은 다른 개념이므로 구분할 필요가 있습니다. C++, Java 등 다른 언어에서는 람다 함수라는 용어를 사용하지만 Go 언어에서는 람다 함수라는 용어를 잘 사용하지 않습니다.



package main
 
func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
 
func main() {
    next := nextValue()
 
    println(next())  // 1
    println(next())  // 2
    println(next())  // 3
 
    anotherNext := nextValue()
    println(anotherNext()) // 1 다시 시작
    println(anotherNext()) // 2
}


위의 코드의 예에서 확인할 수 있는 것과 같이, nextValue()에서 익명함수를 하나 리턴하고 있다. Go언어에서 함수는 일급함수라서 리턴값으로도 쓸 수 있다. 그런데, 안에서 보면 return func()에 있는 i는 밖에 있는 i를 참조하고 있다.



package main

import "fmt"

func main() {  
  counter := newCounter()
  fmt.Println(counter())
  fmt.Println(counter())
}

func newCounter() func() int {  
  n := 0
  return func() int {
    n += 1
    return n
  }
}


이 예제에서 클로저는 n을 newCounter()가 다 끝난 후에 참조하고 있다. 즉, 클로저가 호출 된 횟수를 추적하는 변수에는 액세스 할 수 있지만 newCounter () 함수 외부의 다른 코드는 n에 액세스 할 수 없습니다. 클로저를 통해 변수를 고립시킬 수 있는 것입니다. 함수 호출간에 데이터를 유지하면서 다른 코드에서 데이터를 분리 할 수도 있습니다.



GO 와 클로저 활용
데이터를 고립시키기


만약 함수가 끝난 후에도 변수에 접근할 수 있게 만들고 싶다고 해 보자. 몇 번 함수가 호출되었는지나, 피보나치 수열을 만들 고 싶을 때. 그렇지만, 함수 본인 외엔 다른곳에서 접근할 수 없게 하려고 하면 클로져가 제격이다.



package main

import "fmt"

func main() {  
  gen := makeFibGen()
  for i := 0; i < 10; i++ {
    fmt.Println(gen())
  }
}

func makeFibGen() func() int {  
  f1 := 0
  f2 := 1
  return func() int {
    f2, f1 = (f1 + f2), f2
    return f1
  }
}


클로져가 있으면, 함수나 인터페이스를 만든 후 귀찮게 던져 줄 필요가 없어진다.



type Generator interface {  
  Next() int
}

func doWork(g Generator) {  
  n := g.Next()
  fmt.Println(n)
  // ... do work with n
}

// 클로져 사용시
func doWork(f func() int) {  
  n := f()
  fmt.Println(n)
  // ... do work with n
}

함수를 감싸고 미들웨어를 만들 수 있다


다시한번 말하지만 고언어 함수는 일급함수다. 함수를 파라미터로도 사용할 수 있다는 이야기이다. 아래 예제에선 웹 서버를 만들면서 보편적으로 쓰이는 일급함수를 보여준다.



package main

import (  
  "fmt"
  "net/http"
)

func main() {  
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {  
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}


클로져를 딱히 필요로 하지는 않지만, 좀 더 로직을 가지고 핸들러를 감싸는 데 쓰는 것은 좋다. 타이머 미들웨어의 예제를 봐보자.



package main

import (  
  "fmt"
  "net/http"
  "time"
)

func main() {  
  http.HandleFunc("/hello", timed(hello))
  http.ListenAndServe(":3000", nil)
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {  
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    f(w, r)
    end := time.Now()
    fmt.Println("The request took", end.Sub(start))
  }
}

func hello(w http.ResponseWriter, r *http.Request) {  
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}


timed()함수에서 함수를 아규먼트로 받고 또 받은 함수와는 다른 함수로 리턴을 하고 있다.



Middleware is a common pattern when writing web applications, and can be useful for more than just timers. It can also be used to verify that a user is authenticated, log web requests, write default headers, and much more.


보통은 접근 불가능한 데이터에 접근하기


클로저로 다른 곳에서는 접근을 할 수 없는 변수를 만들 수 있다. 글로벌 변수 없는 코드로 예를 들어보자.



package main

import (  
  "fmt"
  "net/http"
)

type Database struct {  
  Url string
}

func NewDatabase(url string) Database {  
  return Database{url}
}

func main() {  
  db := NewDatabase("localhost:5432")

  http.HandleFunc("/hello", hello(db))
  http.ListenAndServe(":3000", nil)
}

func hello(db Database) func(http.ResponseWriter, *http.Request) {  
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, db.Url)
  }
}


Database 객체에 http.HandleFunc()에서 접근하고 있다. 위 함수로는 글로벌 변수 등의 사용자 변수를 통해서는 접근할 수 없게 이루어졌다. http.handleFunc로 hello(db)를 이용한 클로져로만 접근하는 것이다.


sort패키지를 이용한 바이너리 검색

sort 패키지 이용시에는 클로저를 활용하는것이 좋다.

package main

import (  
  "fmt"
  "sort"
)

func main() {  
  numbers := []int{1, 11, -5, 7, 2, 0, 12}
  sort.Ints(numbers)
  fmt.Println("Sorted:", numbers)
  index := sort.SearchInts(numbers, 7)
  fmt.Println("7 is at index:", index)
}


sort.Search() 함수를 사용할 때는 클로저를 사용한다.

package main

import (  
  "fmt"
  "sort"
)

func main() {  
  numbers := []int{1, 11, -5, 8, 2, 0, 12}
  sort.Ints(numbers)
  fmt.Println("Sorted:", numbers)

  index := sort.Search(len(numbers), func(i int) bool {
    return numbers[i] >= 7
  })
  fmt.Println("The first number >= 7 is at index:", index)
  fmt.Println("The first number >= 7 is:", numbers[index])
}

defer 할때


자바스크립트의 콜백함수를 떠올려보자.

doWork(a, b, function(result) {  
  // use the result here
});
console.log("hi!");  


우리가 근본적으로 하고자 하는 일은 우리의 프로그램이 변수 a와 b를 파라미터로 doWork () 함수를 실행하도록 명령하는 것입니다. 마지막 인수는 함수가 완료된 후에 실행되었으면 좋겠는 거구요. 따라서 doWork ()가 끝나면 doWork ()의 ​​결과로 우리 함수를 호출합니다.



이 접근법의 이점은 doWork ()가 비동기 함수로 코딩 될 수 있다는 것입니다. doWork ()를 호출하고 "hi!"를 인쇄 한 후 우리 코드를 계속 사용할 수 있습니다.



이 예제가 심히 어려운 것은 아니지만... 여러 중첩함수를 사용하면 일반적으로 콜백지옥이라고 하는 그걸 만들 수도 있습니다...

doWork1(a, b, function(result) {  
  doWork2(result, function(result) {
    doWork3(result, function(result) {
      // use the final result here
    });
  });
});
console.log("hi!");  


이미 무슨 소린지도 모르겠습니다. promises 지옥이 펼쳐질거 같습니다. 이런 코드 방식은 고에서도 가능합니다.



go func() {  
  result := doWork1(a, b)
  result = doWork2(result)
  result = doWork3(result)
  // Use the final result
}()
fmt.Println("hi!") 


이 접근 방식에는 두 가지 주요 이점이 있습니다. 첫 번째는 아주 알아 보기가 쉽습니다. javascript 예제에서는 콜백지옥이 펼쳐졌지만 doWork1 () 및 doWork2 ()가 Go 코드에서 완료된 후에 doWork3 ()이 실행된다는 것을 쉽게 알 수 있습니다. go func로 실행하고 있고 있기 때문에이 모든 작업이 동시에 발생합니다.



이 방법의 두 번째 이점은 doWork1 ()의 작성자가 비동기 버전의 함수 작성에 대해 걱정할 필요가 없다는 것입니다. Go에서 doWork1 ()을 호출하면 전체 프로그램이 해당 함수가 실행을 중지 할 때까지 기다린 다음 결과를 얻으면 계속 진행할 것입니다. 이 함수가 동시에 작동해야하는 경우 이전 예제와 같은 goroutine에서 실행하면 됩니다.



참고자료

고 블로그 시리즈3 고 블로그 시리즈2 고 블로그 시리즈1
