## Iterator 와 ListIterator

--

[참조자료](http://aroundck.tistory.com/822)
[ListIterator API](http://docs.oracle.com/javase/7/docs/api/java/util/ListIterator.html)

--


#### Iterator

Set, List, 그리고 Map 엘리먼트에 순차 접근.
단방향 traverse만 지원
기존 element를 삭제할 수는 있으나 추가할 수 없다.

- hasNext();
- next();
- remove();

#### ListIterator

List type의 object만 travers할 수 있다.
양방향 traverse를 지원한다.
기존 element 삭제와 추가를 지원한다.

- add()
- hasNext()
- next()
- hasPrevious()
- previous()
- remove()
- nextIndex()
- previousIndex()