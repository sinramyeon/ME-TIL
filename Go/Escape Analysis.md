# Language Mechanics On Escape Analysis

---

자바에서 포인터 분석...JVM 최적화 등을 배울 때울 기억하자... 기억이 잘 나지 않지만 [escape analysis](https://www.ibm.com/developerworks/java/library/j-jtp09275/index.html) 기법이 있었다.


Escape analysis란 건 프로그램에서 작성한 값의 배치를 정하기 위해 컴파일러에서 사용하는 프로세스다. 특히 컴파일러는 정적 코드 분석을 해서 값이 스택 프레임에 값을 구성 할 수 있는지 또는 값을 힙으로 옮겨야하는지 여부를 결정한다. Go에서는 바로 컴파일러에게 이걸 시키는 함수나 키워드 대신에, 컨벤션을 통해서만 수행한다.

* Escape analysis

(뭔지모르겠음)

---

## 힙

힙은 메모리의 두번째 영역이다. 값을 저장하는 데에 사용된다. 스택처럼 힙은 지알아서 정리를 못한다. 그렇기 때무에 힙을 쓰려면 더 많은 코스트를 필요로 한다. 가비지 콜렉터를 생각해 보자. 가비지 콜렉터가 실행되는 중에는 CPU 용량을 25퍼센트나 먹고 있다. 거기에 레이턴시도 발생시킨다. 가비지 콜렉터를 사용할 때의 장점? 힙에 남긴 메모리를 걱정할 필요가 없어진다.

힙의 값은 Go의 메모리 할당을 구성한다. (The heap is a second area of memory, in addition to the stack, used for storing values.) 더 이상 포인터로 참조되지 않는 힙의 없어져야 하니까 GC에 부담을 주게 된다. 확인하고 제거해야하는 값이 많을수록 GC가 매 실행시 더 많은 작업을 수행해야 하게 되므로, 페이싱 알고리즘은 힙 크기와 실행 속도의 균형을 맞추기 위해 끊임없이 노력하고 있다.(So, the pacing algorithm is constantly working to balance the size of the heap with the pace it runs at.)

---

## 스택 공유

고에서 고루틴은 다른 고루틴 스택을 가르키는 포인터를 가질 수 없다. 고루틴 스택 메모리는 스택 크기가 변화할 때 대체될 수 있기 때문이다. 런타임이 고루틴 스택 포인터를 추적해야 한다면, 관리해야할 게 너무 많아서 스택에서 포인터를 업데이트 할 때 레이턴시가 장난아니게 걸릴 수도 있다.

[예제](https://play.golang.org/p/pxn5u4EBSI)

---

## Escape Mechanics

함수의 스택 프레임 스코프 밖에서 값이 공유될 때, 힙에 그 값이 저장된다. 이러한 상황을 발견하고 프로그램의 무결성을 유지하는 것이 Escape analysis 알고리즘이 하는 일이다. 모든 값에 대한 액세스가 항상 정확하고 일관되며 효율적이라는 것을 증명해야 한다.

아래의 예제를 보자.

```
01 package main
02
03 type user struct {
04     name  string
05     email string
06 }
07
08 func main() {
09     u1 := createUserV1()
10     u2 := createUserV2()
11
12     println("u1", &u1, "u2", &u2)
13 }
14
15 //go:noinline
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
25
26 //go:noinline
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

`go:noinline`을 이용해서, 컴파일러가 main에서 직접 이 함수에 inline하지 못하게 하고 있다. inline하면 함수 호출이 지워지고 문제를 복잡하게 만든다.

두개의 다른 함수가 user값을 만들어다 리턴하고 있는데, 첫 번째 함수부터 알아보자.

```
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
```

이 함수에 의해 생성 된 user 값이 복사되고 호출 스택을 전달했기 때문에 함수가 반환값에 value semantics 사용하고 있다. 즉, 호출하는 함수가 값 자체의 복사본을 받고 있는 것이다.

17 ~ 20 행에서 user 값이 어떠헥 정해지는지 볼수 있다. 23 행에서 값 사본이 호출한 스택 위로 전달되어 호출자에게 다시 전달된다. 함수가 리턴하면 스택은 다음과 같이 생기게 된다.

![](https://www.goinggo.net/images/goinggo/81_figure1.png)

createuserV1을 호출한 후 user 값이 각 프레임 안에 생겼다.

다른 함수는 어떻게 하고 있을까?

```
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

리턴할 때 포인터를 쓰고 있는데 이 함수가 만들고 있는 user값은 콜스택이랑 공유되고 있기 때문이다. 함수를 호출할 때, 주소값 복사본을 받고 있다는 얘기다.

28-31줄에서 user값을 구성하기 위해 같은 스트럭트가 사용되고 있는데, 34줄이 조금 다르다. 콜스택을 백업하기 위해 user값의 복사본을 주는 게 아니고, user값 주소 복사본을 주었다.

![](https://www.goinggo.net/images/goinggo/81_figure2.png)


포인터가 더이상 유효하지 않은 곳을 참조하고 있다. main에서 다음 함수를 호출할 때는, 포인터가 참조하는 곳이 다시 프레임되고 재구성된다.

여기서 escape analysis가 무결성을 유지하는 것이다. 이런 경우에 컴파일러는 createUserV2의 스택 프레임 안에 user값을 생성하는 것이 안전하지 않다고 판단하므로 대신 힙에 값을 생성한다.

---

## 가독성

함수는 프레임 포인터를 통해 프레임 내부의 메모리에 직접 액세스 할 수 있지만 프레임 외부의 메모리에 액세스하려면 간접 액세스를 이용하지 않으면 안 된다. 즉, 힙을 벗어나는 값에 대한 액세스는 포인터를 통해 간접적으로 이루어져야 한다는 것이다.

```
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

user를 선언한 변수 u는, 34라인에서 리턴될때까지 고에서 메모리 어디에 이 값이 있는지 정해지지가 않는다.(Construction in Go doesn’t tell you where a value lives in memory, so it’s not until the return statement on line 34, do you know the value will need to escape.) 이 말인즉슨, u가 user타입을 대변하는 변수지만 user값에 접근하기 위해서는 포인터를 이용해야만 한단 것이다.

![](https://www.goinggo.net/images/goinggo/81_figure3.png)


createUserV2에 대한 스택 프레임의 u 변수는 스택이 아닌 힙에있는 값을 나타낸다. 즉, u를 사용하여 값에 접근하기 위해서는 포인터를 이용해야 한다. 

```
27 func createUserV2() *user {
28     u := &user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", u)
34     return u
35 }
```

return에서 , u의 복사본이 콜 스택에 전달된것을 알 수 있다.

```
34     return &u
35 }
```

이런식으로 `&`를 사용한다면 u가 콜 스택과 공유되고 있고, 힙에서 escape 되고 있는 것이다. 포인터는 곧 공유이다. 또 다른 예제를 생각해 보자.

```
01 var u *user
02 err := json.Unmarshal([]byte(r), &u)
03 return u, err
```

json.Unmarshal을 콜할땐 무조건 포인터값을 넘겨줘야 한다. 이 json.Unmarshal이 user값을 만들고 그 주소를 포인터값에 넘겨준다.

01줄에서 user 포인터 값을 만들고 초기값으로 지정된 후
02줄에서 u값을 json.Unmarshal에 넘겼다.
03줄에서는 u값 복사본을 넘긴다.

가독성이 그리 좋지 않다. 어떻게 해야 좀 더 명확하게 표시할 수 있을까?

```
01 var u user
02 err := json.Unmarshal([]byte(r), &u)
03 return &u, err
```

위와 같이 코드를 리팩토링한다.

01줄에서 user값을 만든 후 초기값으로 지정하고,
02줄에서 json.Unmarshal값과 공유했다
03줄에서는 u를  넘긴다.

훨씬 가독성이 좋다. 02줄에서 user값을 json.Unmarshal 콜스택과 공유하고 03줄에서는 user값을 호출자와 공유한다. 이렇게 user값을 escape시킬 수 있는 것이다.

& 연산자를 쓰는 것이 훨씬 가독성 좋은 코드를 만들 수 있다.

---

## 컴파일러 레포팅

컴파일러가 내린 결정을 보려면 이런 명령을 써볼 수 있다. go 빌드시 -m 옵션과 함께 -gcflags 를 사용하면 된다.

실제로 사용할 수있는 -m 수준은 4가지이지만, 2 단계를 넘어서는 정보는 너무 많다. 보통은 2단계 정도로 사용한다.

```
$ go build -gcflags "-m -m"
./main.go:16: cannot inline createUserV1: marked go:noinline
./main.go:27: cannot inline createUserV2: marked go:noinline
./main.go:8: cannot inline main: non-leaf function
./main.go:22: createUserV1 &u does not escape
./main.go:34: &u escapes to heap
./main.go:34: 	from ~r0 (return) at ./main.go:34
./main.go:31: moved to heap: u
./main.go:33: createUserV2 &u does not escape
./main.go:12: main &u1 does not escape
./main.go:12: main &u2 does not escape
```

createUserV1하고 createUserV2의 흐름을 확인 가능하다.


```
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
 
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

위 코드를 가동시 `./main.go:22: createUserV1 &u does not escape`라고 나온다.

createUserV1안 println을 해 봐야 user값이 힙에서 escpae하지 못했다는 뜻이다.


```
./main.go:34: &u escapes to heap
./main.go:34: 	from ~r0 (return) at ./main.go:34
./main.go:31: moved to heap: u
./main.go:33: createUserV2 &u does not escape
```

31행에 할당 된 u 변수와 연결된 User 값은 34행에서 반환되므로 escape 되고 있다. println에서는 여전히 escape가 일어나지 않았다.

u를 *user 타입으로 바꿔보자.

```
27 func createUserV2() *user {
28     u := &user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", u)
34     return u
35 }
```

```
./main.go:30: &user literal escapes to heap
./main.go:30: 	from u (assigned) at ./main.go:28
./main.go:30: 	from ~r0 (return) at ./main.go:34
```

위와 같이 결과가 나오며, u값은 escape 되고 있다.

---

## 결론

값이 어떻게 공유되는지에 따라 컴파일러가 하는 일이 바뀝니다. 콜스택에 값 공유 할 때마다, escape 될 것입니다.

 Each semantic comes with a benefit and cost. Value semantics keep values on the stack which reduces pressure on the GC. However, there are different copies of any given value that must be stored, tracked and maintained. Pointer semantics place values on the heap which can put pressure on the GC. However, they are efficient because there is only one value that needs to be stored, tracked and maintained. The key is using each semantic correctly, consistently and in balance.