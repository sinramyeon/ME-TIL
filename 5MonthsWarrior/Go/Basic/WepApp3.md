# 추가주제

## HTML 파일서버

자바스크립트, CSS, 이미지 파일 등 분리해서 쓰기

```
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String(
	  "addr",
	  ":8080",
	  "address of the webserver"
)

root = flag.String(
	"root",
	"/var/www",
	"root directory",
)
)

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(
		*addr,
		http.FileServer(http.Dir(*root)),
))
}
```

이제 /localhost:8000/css/style.css 등으로 파일 접근이 가능하다.

만약 주소와 파일 접근 경로가 다르다면

`http.Handle("/css", http.StripPrefix("/css"/ http.FileServer(http.Dir("path/to/cssfiles")), ))` 등으로 설정하여 사용 가능하다.

## 몽고디비 연동

```
type MongoAccessor struct {
	session *mgo.Session
	collection *mgo.Collection
}

func New(path, db, c string) task>accessor {
	session, err := mgo.Dial(path)
	if err != nil {
	return nil
	}

collection := session.DB(db).C(c)

return &MongoAccessor{
	session : session,
	collection : collection,
}
}

func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}

func idToObjectId(id task.ID) bson.ObjectId{
	return bson.ObjectIdHex(String(id))
}

func objectIdToID(objID bson.ObjectId) task.ID{
	return task.ID(objID.Hex())
}

func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
	t := task.Task{}
	err := m.collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

func (m *MongoAccessor) Put(id task.ID, t task.Task) error {
	return m.collection.UpdateId(idToObjectId(it), t)
}

func (m *MongoAccessor) Post(t task.Task) (task.ID, error) {
	objID := bson.NewObjectId()
	_, err := m.collection.UpsertId(objID, &t)
	return objectIdToID(obID), err
}

func (m *MongoAccessor) Delete(id task.ID) error {
	return m.collection.RemoveId(idToObjectId(id))
}
```


