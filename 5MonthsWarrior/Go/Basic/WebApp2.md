# HTML 템플릿 작성

```
func (s status) String() string {
	switch s {
		case UNKNOWN : 
			return "UNKOWN"
		case TODO :
			return "TODO"
		case DONE :
			return "DONE"
		default :
			return ""

}
}

func (s status) MarshalJSON() ([]byte, error) {

	str := s.String()
	if str == "" {
	return nil, errors.New("unknown value")
}
return []byte(fmt.Sprint(str)), nil
}
```

위 코드를 핸들링하는 htmlHandler를 만든다

```
const htmlPrefix = "/task."
var tmp = template.Must(template.ParseGlob("html/*.html"))

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
	return
}

getID := func() (ID, error) {
	id := task.ID(r.URL.Path[len(htmlPrefix):])
	if id == "" {
	return id, errors.New("ID EMPTY")
}
return id, nil
}

id, err := getID()
if err != nil {
	return
}


t, err := m.Get(id)
err = tmpl.ExecuteTemplate(w, "task.html", &Response{
	ID: id,
	Task : t,
	Error : ReseponseError(err),
})

if err!= nil {
	return
}
}

```

task.html을 불러와 /task/1에서 보여준다.

---

# 라우터 사용

유명 라우터로는 gorilla/mux 등이 있다.

```
const (
apiPathPrefix = "/api/v1/task/"
htmlPathPrefix = "/task/"
idPattern = "/{id:[0-9]+}/"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix(htmlPathPrefix).Path(idPattern).Methods("GET").HandlerFunc(htmlHandler)
s:= r.PathPrefix(apiPathPrefix).Subrouter()
s.HandleFunc(idPattern, apiGetHandler).Methods("GET")
s.HandleFunc(idPattern, apiPutHandler).Methods("PUT")
s.HandleFunc("/", apiPostHandler).Methods("POST")
s.HandleFunc(idPattern, apiDeleteHandler).Methods("DELETE")

http.Handler("/", r)
log.Fatal(ListenAndServe("8884"), nil))
}
```

url구조에서 id 패턴을 정해주고 그에 맞는 메서드로 보낼 수 있다.

apiHandler는 아래와 같이 구현가능하다.

```
func getTasks(r *http.Request) ([]task.Task, error) {
	var result []task.Task
	if err := r.ParseForm(); err != 	nil {
	return nil, err
}

encodedTasks, ok := r.PostForm["task"]

if !ok {
	return nil, errors.New("not ok")
}

for _, encodedTask := range encodedTasks {

var t task.Task
if err := json.Unmarshal(]
byte(encodedTask), &t); err != nil {
	return nli, err
}

result = append(result, t)
}
return result, nil
}
```

이제 id는 이렇게 가져온다.

`id := task.ID(mux.Vars(r)["id"])`

이제 나머지 핸들러들을 아래와 같이 작성한다.

```
func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
t, err := m.Get(id)
err = json.NewEncoder(w).Encode(Response{
	ID: id,
	Task : t,
	Error : ResponseError{err},
	})
if err != nil {
	log.Println(err)
}
}

func apiPutHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	tasks, err := getTasks(r)
	if err != nil {
	return
}

	for _, t := range tasks {
	err = m.Put(id, t)
	err = json.NewEncoder(w).Encode(Response {
	ID: id,
	Task : t,
	Error : ResponseError{err},
	})
if err != nil {
	log.Println(err)
}
}
}

func apiPostHandler(w http.ResponserWriter, r *http.Request) {
	tasks, err := getTasks(r)
	if err != nil {
	return
}

for _, t := range tasks {
	err = m.Post(id, t)
	err = json.NewEncoder(w).Encode(Response {
	ID: id,
	Task : t,
	Error : ResponseError{err},
	})
if err != nil {
	log.Println(err)
}
}
}

func apiDeleteHandler(w http.ResponseWriter, r  *http.Request) {

	id := task.ID(mux.Vars(r)["id"])
	err := m.Delete(id)
	err = json.NewEncoder(w).Encode(Response {
	ID: id,
	Task : t,
	Error : ResponseError{err},
	})
if err != nil {
	log.Println(err)
}
}
