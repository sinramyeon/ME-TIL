# 웹 애플리케이션 만들기

헬로 월드~

```
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
})
log.Fatal(http.ListenAndServe(":8080", nil))
}
```

curl http://localhost:8080/ 으로 확인해보자.

### RESTful API

Restful은 대략, 클라이언트-서버 구조가 있고, state가 없고, 캐시가 가능하면서 URL로 어떤 자원에 접근할지 지정하고, GET, PUT, POST, DELETE 메서드를 이용하는 것이라고 이해하면 된다.

### Data Access Object

```
type ID string

type DataAcess interface{
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}

type MemoryDataAccess struct {
	tasks map[ID]task.Task
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess {
	tasks : map[ID]task.Task{},
	nextID : int64(1),
}
}

var ErrTaskNotExist = errors.New("task does not exist")

func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
	return task.Task{}, ErrTaskNotExist
}

return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; 	!exists {
	return ErrTaskNotExist
}

m.tasks[id] = t
return nil
}

func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; 	!exists {
	return ErrTaskNotExist
}

delete(m.tasks, id)
return nil
}
```

위와 같은 코드는 REST api에서 다음과 같이 호출한다.

> GET /api/v1/task/{id} 
> PUT /api/vi/task/{id}
> task : {task}
> POST /api/v1/task/
> task : {task}
> DELETE /api/v1/task/{id}

위와 같은 요청이 올때 응답은 아래와 같이 한다.

```
type ResponseError struct {
	Err error
}

fun (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
	return []byte("null"), nil
}

return []byte(fmt.Sprint(err.Err)), nil
}

func (err *ResponseError) UnmarshalJSON(b []byte) error {

	var v interface{}
	if err := json.Unmarshal(b,v); 	err != nil {
	return err
}

if v == nil {
	err.Err = nil
	return nil
}

switch tv := v.(type) {
	case string :
		if tv == 	ErrTaskNotExist.Error() {
	err.Err = ErrTaskNotExist
	return nil
}

err.Err = erros.New(tv)
return nil

default : 
	return errors.New("unmarshal failed")
	}
}
}
```

이제 응답 구조체를 정의한다.

```
type Response struct {
	ID ID `json:"id:omitempty"`
	Task task.Task `json:"task"`
	Error ResponserERrror `json:"error"`
}
```

*Json Field 주의점*

Json ID 필드가 int64형이면 ID필드에 JSON 태그로 string 을 추가해줘야 한다. 자바스크립트엔 정수형이 따로 없어 실수형이기 때문에 53비트까지만 정확하게 읽을 수 있다. 그러므로 int64필드에 string 태그를 달아주는 버릇을 가지면 좋다.

메인함수에서 호출은 아래와 같이 한다.

```
var m = NewMemoryDataAccess()
const PathPrefix = "api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}
```

핸들로 구현은 아래와 같이 한다.

```
func apiHandler(w http.ResponseWriter, r *http.Request) {

 getID := func() (ID, error) {
	id := task.ID(r.URL.Path[len(pathPrefix):]
	if id == "" {
	retun id, erros.New("ID is empty")
}

return id, nil
}

getTasks := func() ([]task.Task, error) {
	var result []task.Task
	if err := r.ParseForm(); err !=nil {
	return nil, err
}

encodedTasksk, ok := r.PostForm["task]

if !ok {
	return nli, errors.New("task prameter expected")
}

for _, encodedTask : =range encodedTasks {
	var t task.Task
	if err := json.Unmarshal ([]byte(encodedTask), &t); err != 	nil{
	return nil, err
}
result = append(result, t)
}

return result, nil

}

...

}
```

핸들러 안에서 메서드에 따라 다른 코드를 수행되게 하려면 switch 문 등을 이용한다.

```
switch r.Method {
case "GET" :
case "PUT" :
case "POST" :
case "DELETE" :
...
}

```

구현하면 아래와 같다.

```
switch r.Method {
	case "GET" : 
		id, err := getID()
		if err != nil {
	log.Println(err)
	return
}
		t.err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response {
	ID : id,
	Task: t,
	Error : ResponseError{err},
})
	if err != nil {
	log.Println(err)
	return
}

case "PUT" :
	id, err := getID()
	if err != nil {
	return
	}
	tasks, err := getTasks()
	if err != nil {
	log.Println(err)
	return
	}
	for _, t := range tasks {
	err = m.Put(id, t)
	err = json.NewEncoder(w).Encode{Response {
	ID: id,
	Task: t,
	Error: ResponseError{err},
})

	if err != nil {
	log.Println(err)
	return
}

case "POST" :
	tasks, err:= getTasks()
	if err != nil {

	log.Println(err)
	return
	}
	for _, t := range tasks {
	id, err := m.Post(t)
	err = 	json.NewEncoder(w).Encode(Response{
	ID: id, Task: t, Error : ResponseRror{err},
	})
if err!= nil{
	log.Println(err)
	return
	}

case "DELETE" :
	id, err := getID()
	if err!= nil {
	log.Println(err)
	return
	}
err = m.Delete(id)

err = json.NewEncoder(w).Encode(Response{
	ID : id, Error : ResponseError{err},

})

if err!= nil {
	log.Println(err)
	return
	}
}
