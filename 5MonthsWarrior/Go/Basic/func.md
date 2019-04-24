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

가변인자 함수

```
func WriteTo(w io.Writer lines.. string) (n int64, err error)

lines := []string{"a", "b", "c"}
WriteTo(w, lines...)
WriteTo(w, "a", "b", "c")
```


