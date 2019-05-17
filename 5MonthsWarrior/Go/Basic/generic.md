# 제네릭

> 알고리즘을 표현하면서 자료형을 배제할 수 있는 프로그래밍 패러다임

보통 컨테이너에 많이 제네릭을 이용한다. 컨테이너가 담고 있는 자료형에 상관없이 구현을 하기 위해서이다.

고에서는 인터페이스 활용으로 제네릭을 대체하는 경우가 있다.

```
func FieldNames(s interface{}) ([]string, error) {
  t := reflect.TypeOf(s)
  if t.Kind() != reflect.Struct {
    return nil, errors.New("FieldNames : s is not a struct")
  }
  
  names := []string{}
  n := t.NumField()
  for i:=0; i<n; i ++ {
    names = append(names, t.Field(i).Name)
  }
  
  return names, nil

}

s := struct {
  id int
   Name string
   Age int
}{}

fmt.Println(FieldNames(s))
```
