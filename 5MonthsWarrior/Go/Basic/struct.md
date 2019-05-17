# 구조체
필드들의 묶음


```
type Task struct{
  title string
  done bool
  due *time.Time
}

var myTask = Task{"laundry", false, nil}
```

# consta iota

bool 을 쓸 곳에 enum을 쓰면 확장성이 좋아진다!

```
type Task struct {
  title string
  status status
  due *time.Time
}

type status int

const (
UNKNOWN status = iota
TODO status
DONE status 
)
```

iota로 순서대로 0,1,2, 변수 할당이 가능하다.

# 고와 제네릭

> 제네릭이란? 클래스 내부에서 사용할 데이터 타입을 외부에서 지정하는 기법


```
package org.opentutorials.javatutorials.generic;
 
class Person<T>{
    public T info;
}
 
public class GenericDemo {
 
    public static void main(String[] args) {
        Person<String> p1 = new Person<String>();
        Person<StringBuilder> p2 = new Person<StringBuilder>();
    }
 
}
```

클래스 안에서 쓸 것을 클래스 생성시 마음대로 지정 가능.
컴파일 단계에서 자료타입 오류 검출 가능, 중복 제거와 타입 안정성을 지키는데 도움이 됨.

go에는 제네릭이 없다. 고는 심플함을 기반으로 디자인됐는데 언어학적으로 제네릭이 복잡함을 더하기 때문이라고 한다(안그래도어려움)

```
func TestFib(t *testing.T) {
  cases := []struct{
    in, want int
    }{
      {0, 0},
      {5, 5},
     }
     for _, c := range cases {
      got := seq.Fib(c.in)
      if got != c.want {
        t.Error()
      }
     }
  }
}
```
Go에서틑 이런식으로 테스트를 하곤 하지만 복잡하기 때문에 if문을 이용하던지 테이블 기반 테스트를 하는걸 권장하는 듯 하다.
테이블 기반 테스트는 이런 느낌이다.

```
var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	{"%+a", "[%+a]"},
	{"%#a", "[%#a]"},
	{"% a", "[% a]"},
	{"%0a", "[%0a]"},
	{"%1.2a", "[%1.2a]"},
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
	{"%-+1.2a", "[%+-1.2a]"},
	{"%-+1.2abc", "[%+-1.2a]bc"},
	{"%-1.2abc", "[%-1.2a]bc"},
}
func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
```

상속/비상속을 고에서는 이런식으로 구조체로 구현한다.

```
type Deadline time.Time

type Task struct {
  Title string
  Status status
  *Deadline
}
```

Task 구조체 안의 Status에서 status 관련 메서드를 호출 가능하고, Deadline 필드 이름이 생략되어 있으므로 Task 구조체에서 Deadline 의 메서드를 사용 가능하다.

```
type Address struct {
  City string
  State string
}

type Telephone struct {
  Mobile string
  Direct string
}

type Contact struct {
  Address
  Telephone
}

func ExampleContact(){
  var c Contact
  c.Mobile = "010-123-4567"
  c.Address.City = "SF"
  c.Direct = "N/A"
}
```

직렬화 serialization 이란 객체의 상태를 보관이나 전송가능한 상태로 바꾸는 것이다.
반대로 보관되거나 전송받은걸 다시 객체로 바꾸는걸 역직렬화 deserialization이라고 한다.

대표적 예제의 json을 보자.

```
func Example_marshalJSON() {
	t := Task{
		"Laundry",
		Done,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC))
		
		b, err := json.Marshal(t)
		
		if err != nil {
			log.Println(err)
			return
		}
		
		fmt.Println(string(b))
	}
}

func Example_unmarshalJSON() {
	b := []byte{`{"Title" : "Buy Milk"}`}
	t := Task{}
	err := json.Unmarshal(b, &t)
	
	if err != nil {
		return
	}
	
	fmt.Println(t.Title)
}
```

위와 같이 직렬화/역직렬화가 가능하고 json 태그를 붙여 구조체의 필드 이름과 json의 필드 이름을 다르게 할 수도 있다.


```
type MyStruct struct {
	Title string `json:"title"`
	Internal string `json:"-"`
	Value int64 `json:",omitempty"`
	ID int64 `json:",string"`
}
```

omitempty에서는 ->	Value int64 `json:",omitempty"` 에서 값이 0이면 json 결과를 내지 않는다.
,string의 id에서는 64비트정수를 10진수 문자열로 바꿔서 전달한다. 자바스크립트에서 숫자형은 8바이트 실수형이기 때문에 정수값이 53비트를 넘어서면 정확도가 떨어지기 때문이다. (낮은자리 숫자들이 0으로 바뀜)

---

Gob

go언어 직렬화는 gob 이 있는데 go언어에서만 읽고 쓸 수 있다.

```
func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoer(&b)
	data := map[string]string{"N":"J"}
	
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	
	const width = 16
	
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}	
	}
	
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	
}
```

---












