> 인터페이스 : 메서드들의 묶음

```
interface {
  Method1()
  Method2(i int) error
}

type Loader interface{
  Load(filename string) error
}
```


인터페이스를 사용한 문자 정렬 예제

```
type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
  return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
  return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensivite) Swap(i, j int) {
  c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
  apple := CaseInsensitive([]string {
    "iPhone", "IPAD"
  })
  
  sor.Sort(apple)
  /// IPAD, iPhone 순서로 정렬됨
}
```

---

힙

힙은 자료중 가장 작은값은 O logN의 시간복잡도로 꺼낼 수 있는 자료구조이다.

```
type Interface interface{
  sort.Interface
  Push(x interfae{})
  Pop() interface{}
}
```

힙 선언을 해 보았다. (Push/Pop)

```
func (c *CaseInsensitive) Push(x interface{}){
  *c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
  len := c.Len()
  last := (*c)[len-1]
  
  *c = (*c)[:len-1]
  return last // 마지막 거부터 꺼냄
}

func ExampleCaseInsensitive_heap() {
  apple := CaseInsensitive([]string{
    "iPhone", "IPAD"
  })
  heap.Init(&apple)
  for apple.Len() >0 {
    fmt.Println(heap.Pop(&apple))
  }
  // IPAD, iphone 순서로 나옴
}
```

---

고언어 형 단언 예제

```
func Join(sep string, a ...interface{}) string {
  if len(a) == 0 {
    return ""
  }
  
  t := make([]string, len(a))
  
  for i:= range a {
    switch x := a[i].(type) {
      case string :
        t[i] = x
      case int :
        t[i] = strcov.Itoa(x)
      case fmt.Stringer :
        t[i] = x.String()
      }
  }
  return strings.Join(t, sep)
}
```

switch문 안에서 type 형 단언으로 자료형마다 구현을 다르게 한 코드를 볼 수 있다.

if문을 이용해 아래와도 같이 표현 가능하다.

```
if x, ok := a[i].(string); ok {
  t[i] = x
} else if x, ok := a[i].(int); ok {
  ...
```

















