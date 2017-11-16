# 스택, 포인터

----

스택, 포인터, 정적 분석 기법, 메모리 프로파일링, 데이터와 시맨틱을 Go로 풀어낸 4 시리즈를 정리했다.

[Language Mechanics On Stacks And Pointers](https://www.goinggo.net/2017/05/language-mechanics-on-stacks-and-pointers.html)

[Language Mechanics On Escape Analysis](https://www.goinggo.net/2017/05/language-mechanics-on-escape-analysis.html)

[Language Mechanics On Memory Profiling](https://www.goinggo.net/2017/06/language-mechanics-on-memory-profiling.html)

[Design Philosophy On Data And Semantics](https://www.goinggo.net/2017/06/design-philosophy-on-data-and-semantics.html)

----

포인터는 너무 너무 어렵다. 잘못 썼다간 어디서 초래됐는지도 모를 버그를 만들고 또 성능에도 영향을 끼친다. 멀티스레드나 동시성 프로그래밍을 하려면 어쩔 수 없이 거쳐갈 관문이지만 솔직히 Go에서는 초장부터 비껴나갈 수 없는 게 Pointer이기도 하다.

## 프레임 바운더리

프레임 스코프 안에서 함수가 실행되면 각 함수바다 개별적으로 메모리 공간이 제공된다. 프레임은 함수가 자기 위치에서 작동되도록 해 주고 플로우 컨트롤을 할 수 있게 도와준다. 함수는 프레임 포인터를 이용해서 프레임 안에 있는 메모리에 바로 접근할 수 있다. 그렇지만, 밖에서 메모리에 접근하려고 하면 직접적으로 접근하지는 못하게 되어 있다. 프레임 밖에서 함수가 메모리에 접근하려고 들면, 메모리가 꼭 함수랑 공유된 상태여야만 한다. 

함수가 호출되었을 때, 두 개의 프레임 사이 트랜지션이 일어난다.(The code transitions out of the calling function’s frame and into the called function’s frame.) 함수 호출을 만들기 위해 데이터가 필요한 경우 해당 데이터를 한 프레임에서 다른 프레임으로 전송해야 한다. 두 프레임 사이의 데이터 전달은 Go에서 "값"에 의해 수행된다. (Call by Value를 기억해 두자)


Call by Value식 값 전달의 이점은 역시 가독성이다.(알아먹기 쉽다)  함수 호출에서 확인할 수 있는 값은 다른 쪽에서 복사해서 가져온 값이다. 이 모든 기능을 통해 두 기능 간 트랜지션에 드는 비용을 숨기지 않는 코드를 작성할 수 있다.

아래의 예제를 보자.

```
01 package main
02
03 func main() {
04
05    // Declare variable of type int with a value of 10.
06    count := 10
07
08    // Display the "value of" and "address of" count.
09    println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
10
11    // Pass the "value of" the count.
12    increment(count)
13
14    println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
15 }
16
17 //go:noinline
18 func increment(inc int) {
19
20    // Increment the "value of" inc.
21    inc++
22    println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
23 }
```

위 프로그램을 실행 시 메인 고루틴을 만들어서 코드를 실행한다. (A goroutine is a path of execution that is placed on an operating system thread that eventually executes on some core.) 고 버전 1.8을 기준으로 고루틴 한개당 2048 바이트메모리를 스택에서 차지한다. 이 스택 사이즈는 아마 나중에 좀 바뀌어 있을 수도 있다.

스택은 각 개별 기능에 지정된 프레임 바운더리에 대한 물리적 메모리 공간을 제공한다. 메인 goroutine이 Listing 1의 기능을 수행 할 때 goroutine의 스택은 다음과 같이 보일 것이다.

![](https://www.goinggo.net/images/goinggo/80_figure1.png)

메인 함수를 위해 스택 중 일정 부분이 할당되어 있다. 이 할당된 부분을 스택프레임이라고 부르고, 이 프레임이 메인함수의 경계를 나타내 준다. 프레임은 함수가 호출 될 때 실행되는 코드의 일부로 설정된다. count 변수의 메모리가 메인 프레임의 주소 0x10429fa4에 저장되어 있음을 확인할 수 있다.

활성된 프레임 아래의 모든 스택 메모리는 유효하지 않지만(비활성상태) 활성 프레임 및 그 이상의 메모리는 유효한것을 확인할 수 있다.

## 주소

변수를 만들면 특정 메모리 공간에 이름을 할당해 준다. 만약, 변수가 선언되있고 값이 할당되어 있으면 메모리 안에 주소도 가지고 있는 것이다.

`09    println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")`

`&` 연산자는 변수의 위치값을 가져와 준다. 아래와 같이 추력되는 것을 확인할 수 있다.

`count:  Value Of[ 10 ]  Addr Of[ 0x10429fa4 ]`

## 함수 호출

`increment(count)`

위와 같이 함수를 호출한다는 것은, 고루틴이 스택 위에 있는 메모리에서 새로운 메모리 섹션을 찾게 된다는 것을 뜻한다.(사실은 이보다 좀 복잡하다.) 성공적으로 함수를 호출하기 위해서는, 데이터는 프레임 경계를 가로 질러 전달 되어야 하며 트랜지션 중 새로운 프레임에 배치되어야 한다.

`18 func increment(inc int) {`

`count`라는 변수로 int값 아규먼트가 전달되고 있는데, 값이 복사된 후 increment 함수를 위해 새 프레임에 전해진다. 함수는 자기 프레임 안에 있는 데이터만 직접적으로 읽고 쓸 수 있기 때문에, `inc`라는 변수로 받아서 저장하고 있다.

![](https://www.goinggo.net/images/goinggo/80_figure2.png)

이제는 하나는 main, 하나는 increment 함수 용으로 두 프레임이 스택 위에 생겼다. inc는 값 10을 갖고 있는데, 함수 호출 시 복사되고 전해진 값이다. inc 변수의 주소는 0x10429f98이며 프레임이 새로 생기면 스택 아래로 내려 가기 때문에 메모리가 밑에 생겨 있다.(The address of the inc variable is 0x10429f98 and is lower in memory because frames are taken down the stack, which is just an implementation detail that doesn’t mean anything.) 중요한 것은 goroutine이 main에 프레임 내에서 count 값을 가져 와서 inc 변수를 사용하여 increment용 프레임에 해당 값의 복사본을 저장한다는 것이다.

```
21    inc++
22    println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
```

위 코드 실행 시 아래와 같은 값을 확인할 수 있다.

`inc:    Value Of[ 11 ]  Addr Of[ 0x10429f98 ]`

![](https://www.goinggo.net/images/goinggo/80_figure3.png)

inc의 값이 바뀌었다.

`14    println("count:\tValue Of[",count, "]\tAddr Of[", &count, "]")`

그러나 count 값은 increment를 실행하기 전과 같을 것이다.

```
count:  Value Of[ 10 ]  Addr Of[ 0x10429fa4 ]
inc:    Value Of[ 11 ]  Addr Of[ 0x10429f98 ]
count:  Value Of[ 10 ]  Addr Of[ 0x10429fa4 ]
```

## 함수 리턴값

함수가 값을 리턴할 때 스택에는 무슨 일이 벌어질까? 답은 아무일도 없다는 것이다. increment 함수값 리턴시 메모리 상태를 확인해 보자.

![](https://www.goinggo.net/images/goinggo/80_figure4.png)

스택은 increment 함수와 관련된 프레임이 유효하지 않은 메모리로 간주된다는 것을 제외하고는 아까랑 똑같다. 이것은 메인 프레임이 현재 활성된 프레임이기 때문이다. increment 함수를 위해 프레임 된 메모리는 변경되지 않은 채로 남겨져 있다.

메모리가 다시 필요할 지 모르기 때문에 리턴 함수의 프레임 메모리를 정리하는 것은 시간 낭비다. 그래서 메모리는 그대로 남아 있다. 각 함수 호출 중에 프레임을 가져 왔을 때 해당 프레임의 스택 메모리가 지워진다.(It’s during each function call, when the frame is taken, that the stack memory for that frame is wiped clean.) 이 작업은 프레임에 있는 값의 초기화를 통해 수행된다. 모든 값은 적어도 각 타입별 초기값으로 초기화되기 때문에 스택은 모든 함수 호출에서 적절하게 자체적으로 정리가 된다.

## 값 공유

만약 increment 함수가 메인 프레임의 내부에있는 count 변수와 관련해 직접 조작해야 하는 경우에는 어떻게 해야 할까요? 여기서 pointer 를 사용해야 한다. 포인터는 값을 함수와 공유하기 위해 사용하는 것이다. 따라서 함수는 값이 자체 프레임 내에 직접 존재하지 않더라도 해당 값을 읽고 쓸 수 있는 것이다.

> 코드속 `&`만 봐도 `값 공유` 부터 떠올리자.

## 포인터 타입

모든 타입마다 *를 붙이자. 그러면 그게 포인터 타입니다. int 에는 *int가 있고 내가 선언한 User 스트럭트가 있다면 *User라고 선언하면 된다.

포인터 타입은 두 가지 공통점이 있다. 하나는 무조건 앞에 `*`가 붙어있다는 것이겠고, 다른 하나는 4나 8바이트의 메모리 사이즈를 갖는다는 것이다. 32비트에서는 4, 64에서는 8이다.

---

## 간접 메모리 접근

```
01 package main
02
03 func main() {
04
05    // Declare variable of type int with a value of 10.
06    count := 10
07
08    // Display the "value of" and "address of" count.
09    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
10
11    // Pass the "address of" count.
12    increment(&count)
13
14    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
15 }
16
17 //go:noinline
18 func increment(inc *int) {
19
20    // Increment the "value of" count that the "pointer points to". (dereferencing)
21    *inc++
22    println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
23 }
```

아까와 바뀐 곳을 확인해 보자.

`12    increment(&count)`

이번에는 count 값을 이용하지 않는다. count 주소를 이용한다. count 변수를 드디어 increment 함수랑 *공유* 하고 있는 것이다. 아까도 말했지만 포인터(&)는 무조건 공유다.

이번엔 integer값 대신 주소를 갖다주고 있을 뿐이다. 주소도 어쨌든 값이기 때문이다. 아까랑 마찬가지로 함수 호출이 되면 주소를 복사하고 전달해 준다.

건네준 곳에서 주소를 썼기 때문에, 이번엔 주소를 받아와야 한다.

`18 func increment(inc *int) {`

똑같이 다 주소인데 왜 *아무거나 라고 표현하지 않을지 궁금해 졌다. 값을 공유하는 이유는 수신하는츨 함수가 해당 값에 대한 읽기 또는 쓰기를 수행해야하기 때문이다. 값을 읽고 쓸 수있는 값의 타입이 필요하게 될 것이다. 컴파일러는 올바른 포인터 유형과 그 타입 값만 해당 함수가 쓰는지 확인해 준다.

![](https://www.goinggo.net/images/goinggo/80_figure5.png)

주소를 전달했을 때는 이런 상태가 된다. increment 함수 프레임 안에 있는 포인터가 이제 count 변수를 가리키고 있다. main 함수 프레임 안에 있는 그 변수 말이다.

> pointer 변수를 씀으로 인해서 간접적 RW가 가능해졌다.

`21    *inc++`

이번에는 * 문자가 연산자로 사용되며 포인터 변수에 적용되었다. *를 연산자로 사용하면 "포인터가 가리키는 값"을 의미하게 된다. 포인터 변수는 함수를 사용하는 함수 프레임 외부의 간접 메모리 접근을 허용해 준다.(포인터 역 참조 - dereference the pointer)  increment 함수는 프레임 내에서 여전히 간접 접근을 하려면 프레임 안에 포인터 변수를 갖고 있어야 한다.

![](https://www.goinggo.net/images/goinggo/80_figure6.png)

이번에는 main함수 안 count 값도 11로 늘었다.

```
count:  Value Of[ 10 ]   	   	Addr Of[ 0x10429fa4 ]
inc:    Value Of[ 0x10429fa4 ]  	Addr Of[ 0x10429f98 ]   Value Points To[ 11 ]
count:  Value Of[ 11 ]   	   	Addr Of[ 0x10429fa4 ]
```

inc 포인터 변수의 값이 count 변수의 주소랑 동일하다. 이렇게 하면 프레임 외부의 메모리에 대한 간접 액세스가 가능하도록 하는 공유 관계가 설정된다. increment 함수를 실행하면 main함수 안에 있는 변수 값이 변하게 되는 것이다.

---

### 포인터는 어렵지 않다!(어렵다...)

포인터 변수는 다른 변수와 같은 변수이기 때문에 그다지 특별할 것 없다. 메모리 할당도 되고 값을 가지고 있다. 포인터가 가리킬 수 있는 값 유형에 관계없이 모든 포인터 변수는 항상 동일한 크기를 갖는다.

## 결론

- Functions execute within the scope of frame boundaries that provide an individual memory space for each respective function.
- When a function is called, there is a transition that takes place between two frames.
- The benefit of passing data “by value” is readability.
- The stack is important because it provides the physical memory space for the frame boundaries that are given to each individual function.
- All stack memory below the active frame is invalid but memory from the active frame and above is valid.
- Making a function call means the goroutine needs to frame a new section of memory on the stack.
- It’s during each function call, when the frame is taken, that the stack memory for that frame is wiped clean.
- Pointers serve one purpose, to share a value with a function so the function can read and write to that value even though the value does not exist directly inside its own frame.
- For every type that is declared, either by you or the language itself, you get for free a compliment pointer type you can use for sharing.
- The pointer variable allows indirect memory access outside of the function’s frame that is using it.
- Pointer variables are not special because they are variables like any other variable. They have a memory allocation and they hold a value.




