# Go Mutable

---


## 사전지식

### Mutable? Immutable?

> mutable은 값이 변한다는 뜻이고, immutable 은 값이 변하지 않는다는 의미이다.

이해하기 쉽도록 python으로 예를 들자면 아래와 같다.

* 숫자형(Number) : immutable

![](http://cfile10.uf.tistory.com/image/276C134655A74C8E09424E)

* 문자열 (string) : immutable

![](http://cfile10.uf.tistory.com/image/2177D13B55A74D97322487)

* 리스트 (list) : mutable

![](http://cfile4.uf.tistory.com/image/244AB53955A74E552B27F4)

* 튜플 (tuple) : immutable

![](http://cfile2.uf.tistory.com/image/2664A14255A74F0E296014)

* 딕셔너리 (Dict) : mutable

![](http://cfile10.uf.tistory.com/image/23722E4655A7504C1405B3)

파이썬에서는 숫자, 문자, 튜플은 변경할 수 없고 리스트와 딕트는 변경이 가능하다. Call-By-Value, Call-By-Reference를 떠올리면 된다. 자바를 처음 배울 때에 정리한 내용을 떠올려 보자.

---

### 문자열 비교하기

---

[참고](http://kmj1107.tistory.com/entry/JAVA-%EB%AC%B8%EC%9E%90%EC%97%B4string-%EB%B9%84%EA%B5%90-equals%EC%99%80-%EC%9D%98-%EC%B0%A8%EC%9D%B4%EC%A0%90-equals%EC%9D%98-%EB%B0%98%EB%8C%80)

1. EQUALS
 
 - 메소드이다.
 - 대상의 내용 자체를 비교한다.
 
 
 
 
2. ==

 - 연산자이다.
 - 주소값을 비교한다.


![참조](http://cfile10.uf.tistory.com/image/182CE2414E705D1C129150)

위 그림을 예로 들면,
a.equals(b)와 a==b는 true지만,
a==c는 false이게 된다.
값만을 비교하려면 a.equals(c)를 해야 한다.

----

#### Call by reference, Call by Value

[참조자료](http://twoday.tistory.com/entry/call-by-value%EC%99%80-call-by-reference-%EC%A0%9C%EB%8C%80%EB%A1%9C-%EC%95%8C%EA%B8%B0)
[각 언어별 방식](https://okky.kr/article/303162?note=1005863)
[예제](http://trypsr.tistory.com/74)
---

## Go 에서의 예시


예를 들자면, 아래와 같이 변수를 할당한다고 치자.
```
x := 1
x = 2
```

그럼 여기서 추가 메모리는 할당되지 않는다.
Go에서는 데이터 구조 상 string 외에는 모두 immutable하다.
이런 점이 고의 단점 중 하나로 뽑히기도 한다.

예시
```
package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
```

위를 실행하면, 아래와 같은 결과가 나온다.
```
Matt released their 43th song
Matt has a total of 42 songs
```

newRelease()로 분명 양을 늘려 주었지만, 값이 변하지 않았다. 여기서 실제 값이 변하게 하려면 포인터를 이용해야만 한다.

```
package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
```

딱 하나 고친 곳은
```
func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}
```
포인터로 받아오는 곳이다. 포인터 값을 받아오기 위해서는 &를 사용한다.




