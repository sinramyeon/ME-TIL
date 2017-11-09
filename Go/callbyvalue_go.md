# Go와 Call by Value

---

[going go의 copying interface values in go](https://www.goinggo.net/2016/05/copying-interface-values-in-go.html)
와 golang의 type assertion에 대하여


고는 Call by Value 기반으로 만들어졌다. 값을 함수에 넘겼을 때,슬라이스를 Iterate 할 때, Type assertion을 할 때 이 점을 상기해볼 수 있다.

Golang에서 Type Assertion은 인터페이스 유형 변수가 보유한 값이 원하는 인터페이스를 구현하는지 또는 구체적인 유형인지 여부를 확인하는 데 사용된다.

```
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) //hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic
	fmt.Println(f) //panic: interface conversion: interface {} is string, not float64

}
```
Go에서는 인터페이스를 쓸 때도 그 형이 무엇인지 정확히 해 주어야 한다. 

Type assertion의 정의를 확인하자.

`PrimaryExpression.(Type)`


이 PrimaryExpression은 반드시 인터페이스 형으로 받아야 하고 아닐 시 컴파일 오류가 발생한다.

```
type I interface {
    walk()
    quack()
}
type S struct{}
S{}.(I)  // invalid type assertion: S literal.(I) (non-interface type S on left)
```

모든 변수에는 static 정적 말고도 dynamic 동적 타입(불변 말고 가변)도 있습니다. 지금은 인터페이스 타입 안에 있는 거죠. 프로그램 실행 과정에서 인터페이스 유형의 변수는 동일한 정적 유형을 갖지만 동적 유형은 원하는 인터페이스를 구현하는 다른 값이 할당 될 때 변경 될 수 있습니다.

```
type I interface {
    walk()
}
type A struct{}
func (a A) walk() {}
type B struct{}
func (b B) walk() {}
func main() {
    var i I
    i = A{}  // dynamic type of i is A
    fmt.Printf(“%T\n”, i.(A))
    i = B{}  // dynamic type of i is B
    fmt.Printf(“%T\n”, i.(B))
}
``` 

그렇다면 인터페이스에 있는 값을 포인터 없이 복사해서 저장하면 어떻게 되는 것인지에 대해 알아보자.


```
06 package main
07 
08 import (
09     "fmt"
10     "unsafe" 
11 ) 
12  
13 // notifier provides support for notifying events.
14 type notifier interface {
15     notify() 
16 } 
17  
18 // user represents a user in the system.
19 type user struct{
20     name string 
21 } 
22  
23 // notify implements the notifier interface.
24 func (u user) notify() {
25     fmt.Println("Alert", u.name)
26 }
```

14줄에서 interface형으로 notifier를 선언하고 안에는 간단한 메서드 notify()를 생성했다. 19줄을 보면 user 스트럭트를 만들어 뒀고 notify 함수에선 user를 구현하고 있다.


```
28 // inspect allows us to look at the value stored
29 // inside the interface value.
30 func inspect(n *notifier, u *user) {
31     word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
32     value := (**user)(unsafe.Pointer(word))
33     fmt.Printf("Addr User: %p  Word Value: %p  Ptr Value: %v\n", u, *value, **value)
34 }
```

30줄에 inspect 함수를 보면, 인터페이스의 두번째 값에 대한 포인터를 주고 있는 것을 확인할 수 있다. 이 포인터를 사용하여 인터페이스의 두 번째 단어 값과 두 번째 단어가 가리키는 user 값을 알아낼 수 있다.

```
36 func main() {
37 
38     // Create a notifier interface and concrete type value.
39     var n1 notifier
40     u := user{"bill"}
41
42     // Store a copy of the user value inside the notifier
43     // interface value.
44     n1 = u
45 
46     // We see the interface has its own copy.
47     // Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
48     inspect(&n1, &u)
49 
50     // Make a copy of the interface value.
51     n2 := n1
52 	
53     // We see the interface is sharing the same value stored in
54     // the n1 interface value.
55     // Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
56     inspect(&n2, &u)
57 
58     // Store a copy of the user address value inside the
59     // notifier interface value.
60     n1 = &u
61 
62     // We see the interface is sharing the u variables value
63     // directly. There is no copy.
64     // Addr User: 0x1040a120  Word Value: 0x1040a120  Ptr Value: {bill}
65     inspect(&n1, &u)	
66 }
```

36줄에서 시작하는 메인함수를 보면 n1 변수로 notifier 인터페이스를 만들고 있고, 값은 초기값으로 설정됩니다. 40줄에는 user 를 만들어서 bill 로 설정했습니다.

44줄에서 n1에 bill을 대입했을때 어떻게 되는지 보겠습니다.

```
42     // Store a copy of the user value inside the notifier
43     // interface value.
44     n1 = u
45 
46     // We see the interface has its own copy.
47     // Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
48     inspect(&n1, &u)
```

실행시 값은 아래와 같이 할당되게 됩니다.

![](https://www.goinggo.net/images/goinggo/69_figure1.png)

인터페이스인 n1이 저희가 할당한 user값 복사본을 갖고 있게 됩니다. 인터페이스 안에 저장된 user 값의 주소값은, 원래 있던 user 주소값이랑 완전 다른 게 된 거죠.

이 코드를 작성하여 포인터가 아닌 값이 할당 된 인터페이스 값의 복사본을 만들면 어떤 일이 발생 할지를 알 수 있었습니다. 아래와 같이 해보면 어떨까요?

```
50     // Make a copy of the interface value.
51     n2 := n1
52 	
53     // We see the interface is sharing the same value stored in
54     // the n1 interface value.
55     // Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
56     inspect(&n2, &u)
```

51줄에서 새 인터페이스로 인터페이스 값을 복사해 왔습니다.
이렇게 하면 어떻게 될까요?

![](https://www.goinggo.net/images/goinggo/69_figure2.png)

우리가 인터페이스 값의 복사본을 만들 때, 그 때 그대로 복사되는 걸 확인할 수 있을 것입니다. n1 내부에 저장된 user 값은 이제 n2와 공유됩니다. 각 인터페이스 값에 자체 복사본이 없습니다. 둘 다 동일한 user 값을 공유합니다.

만약 값 자체 말고 주소를 전달한다면 어떨까요?

```
58     // Store a copy of the user address value inside the
59     // notifier interface value.
60     n1 = &u
61
62     // We see the interface is sharing the u variables value
63     // directly. There is no copy.
64     // Addr User: 0x1040a120  Word Value: 0x1040a120  Ptr Value: {bill}
65     inspect(&n1, &u)
```

![](https://www.goinggo.net/images/goinggo/69_figure3.png)

그 인터페이스는 이제 u 변수가 참조하는 user 값을 가리키고 있음을 알 수 있습니다. u 변수의 주소는 이제 인터페이스 안에 저장되어 있던 그 값으로 변해 있습니다. 즉, u 변수와 연관된 user값의 모든 변경 사항은 인터페이스 값을 보면 알 수 있다는 뜻이 됩니다.

결론

주소가 저장되고 있는 한 주소가 복사되는거지 값이 저혼자 바뀌는것이 아니다.

