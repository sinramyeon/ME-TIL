# Channel

---

Go를 쓰고 싶은 이유 중 하나는 보통 멀티스레드 프로그래밍을 안전하고 빠르고 쉽게(!) 할 수 있다는 점이다. 고루틴은 만들기가 아주 쉽기 때문이다. 그렇다면 channel은 어떨까. channel 에 대해 알아보도록 한다.

## Go 채널

![](https://thebook.io/img/006806/5-1.jpg)
> 채널이란 : 데이터를 추고 받는 통로

공식 설명또한 알아보자.
> "By default, sends and receives block wait until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables."

![](https://camo.qiitausercontent.com/52962702fa3abcb1bb4b59f29f7e880ccc9108ed/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e616d617a6f6e6177732e636f6d2f302f35323631332f38386466333666322d343861372d616562622d373865612d3832616236316664653564642e706e67)

```
package main
 
func main() {
  // 정수형 채널을 생성한다 
  ch := make(chan int)
 
  go func() {
    ch <- 123   //채널에 123을 보낸다
  }()
 
  var i int
  i = <- ch  // 채널로부터 123을 받는다
  println(i)
}
```

채널은 수신자와 송신자가 서로를 기다리는 속성을 가지고 있어서, 고루틴이 끝날 때까지 기다릴 수 있다.

```
package main
 
import "fmt"
 
func main() {
    done := make(chan bool)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    }()
 
    // 위의 Go루틴이 끝날 때까지 대기
    <-done
}
```

이렇게 활용해볼 수 있다.

```
// 일정 시간마다 스케쥴 반복하기
func schedule(delay time.Duration) chan bool {

	stop := make(chan bool) // 채널을 만들어 놓고
	go func() { // 고루틴을 돌렸는데
		for {
			select {
			case <-time.After(delay): //일정 시간이 지나야
			case <-stop: // 채널에 값을 보낸 후 리턴한다!
				return
			}
		}
	}()

	return stop
}
```

#### 채널의 종류와 방향

> 채널의 종류는 두 가지로 나뉘는데, Unbuffered Channel과 Buffered Channel이 있다.

Unbuffered Channel에서는 수신자가 데이타를 받을 때까지 송신자가 데이타를 보내는 채널에 묶여 있게 된다. Buffered Channel을 사용하면 비록 수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이타를 보내고 계속 다른 일을 수행할 수 있다. 

채널 문법 `<-` 로 방향을 설정할 수 있는데, 무슨 소리냐면 이 채널이 보내는 애인지 받는 애인지를 정할 수 있다.

`func sender(c chan <- string)`
위는 보내기만 하는 애이다.

`func receiver(c <- chan string)`
그렇다면 위는 받기만 할 수 있다.


예를 들면,

```
// 전송 전용 용량 10 channel 작성 
c  : =  make ( chan <-  int ,  10 )

// 큐에 송신 
c  <-  10

// 전송 전용이므로 컴파일 에러가 뜬다. 
fmt . Println ( <- c )
```

```
// 수신 전용 용량 10 channel 작성 
c  : =  make ( <- chan  int ,  10 )

// 수신 전용이기 때문에 컴파일 에러 
c  <-  10

fmt . Println ( <- c ) 
```

그러면 이것을 어떻게 사용하고 어디에 쓸지에 대해 생각해 볼 수 있다.

전송 전용 채널의 사용 예제(피보나치 수열 구현하기)
```
package  main

import  ( 
    "fmt" 
)

func  fibonacci ( n  int ,  c  chan <-  int )  { 
    x ,  y  : =  0 ,  1 
    for  i  : =  0 ;  i  <  n ;  i ++  { 
        c  <-  x  // 무조건 보내기만 하고 있다.
        x ,  y  =  y ,  x + y 
    } 
    close ( c ) 
}

func  main ()  { 
    c  : =  make ( chan  int ,  10 ) 
    go  fibonacci ( cap ( c )  c ) 
    for  i  : =  range  c  { 
        fmt . Println ( i ) 
    } 
}
```

*잠깐*
피보나치 수열 구현시 select 문을 사용해 볼 수도 있다. 채널은 기본적으로 받고 보내기가 다 되기 때문에, select 문은 하나의 고루틴이 여러 채널과 통신할 때 사용한다. case로 여러 채널을 대기시키고 있다가 실행 가능 상태가 된 채널이 있으면 해당 케이스를 수행한다.
```
package main
 
import “fmt”
 
// 피보나치 2
func fibonacci2(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x: // 값을 더할 때와
            x, y = y, x+y
        case <-quit: // 끝낼 때를 나눴다.
            fmt.Println(“quit”)
            return
        }
    }
}
 
func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

수신 전용 채널의 사용 예제(소인수 분해하기)
```
package  main

import  ( 
    "fmt" 
)

func  getPrimeGen ()  <- chan  int  { 
    var  gen  =  make ( chan  int ) 
    go  func ()  { 
        var  primes  =  make ([] int ,  0 ) 
        for  k  : =  2 ;  ;  k ++  { 
            check  : =  true 
            for  _ ,  value  : =  range  primes  { 
                if  k  %  value  ==  0  {
                    check  =  false 
                    break 
                } 
            } 
            if  check  { 
                primes  =  append ( primes ,  k ) 
                gen  <-  k 
            } 
        } 
        close ( gen ) 
    } () 
    return  gen 
}

func  main ()  { 
    //받은 channel은 수신 만 
    primeGen  : =  getPrimeGen ()

    for  i  : =  0 ;  i  <  10000 - 1 ;  i ++  { 
        <- primeGen 
    } 
    // 소수의 10000 번째를 출력합니다. 
    fmt . Println ( <- primeGen )
```

#### 채널 닫기

채널을 더이상 안쓸때(전송할 값이 없음)는 닫는다. close()
물론 닫은 후 메시지를 전송하면 에러가 발생한다.



---

## 채널 잘! 사용하기

[채널 잘 사용하기1](https://thebook.io/006806/ch05/02/00_02/)


함수와 마찬가지로 채널도 값에 의한 호출 방식(Call by Value)을 사용해서 값을 전달한다. 즉, 실제 값이 복사되어 전달되기 때문에 bool, int, float64 등의 값을 전달하는 것은 안전하다. Go에서는 문자열과 배열도 변하지 않는 값이므로 채널의 값으로 사용해도 안전하다.
하지만 포인터 변수나 참조 값(슬라이스, 맵)을 채널로 전달할 때는 주소 값이 전달되므로 값을 보내는 고루틴과 값을 받는 고루틴에서 값을 동시에 수정하면 예상치 못한 결과가 발생할 수 있다. 그래서 포인터나 참조 값을 채널로 전달할 때는 여러 고루틴에서 값을 동시에 수정하지 않게 해야 한다. 가장 간단한 방법은 여러 고루틴에서 참조 값에 동시에 접근할 수 없게 뮤텍스로 제한하는 것이다.
참조 값을 동시에 변경할 위험이 있는 고루틴에는 참조 값을 직접 전달하는 것이 아니라, 인터페이스를 전달하는 것도 좋은 방법이다. 참조 값에 대한 읽기 전용 인터페이스를 만들고, 채널로 읽기 전용 인터페이스를 전달하면 안전하게 참조 값을 사용할 수 있다.

```
type mapGetter interface {
    Get(s string) interface{}
}
 
type stringMap map[string]interface{}
 
func (m stringMap) Get(s string) interface{} { return m[s] }
```

stringMap의 값을 변경할 필요가 없는 고루틴에는 stringMap을 직접 전달하지 않고 mapGetter 인터페이스를 채널로 전달한다. 그러면 개발자가 실수로 stringMap의 값을 변경하는 것을 막을 수 있다.

```
func process(g <-chan mapGetter) {
    ...
    map := <-g
    value := map.Get("key")
    ...
}
```

[채널 잘 사용하기2(한문주의)](https://qiita.com/awakia/items/f8afa070c96d1c9a04c9)

병렬처리에 대해 대조적 개념부터 이해해 보자.

모델 | 전달법 | 예를 든다면 ...
----|-------|------
Shared Memory | 여러 프로세스가 락을 걸면서 공통의 메모리에 액세스한다. | 모두가 원형 책상에 둘러싸고 앉아서 메모를 적고 있다. 그렇지만 메모지랑 펜 쟁탈전이 일어날수도 있다.
Message Passing	| 각 프로세스는 서로 같은 내용을 전달한다. | 각자 자신의 책상에 앉아 자기 걸로 메모를 적고 있다.

