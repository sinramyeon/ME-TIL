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
값만을 비교하려면 a.equals(c)를 해줘야겠지?

----

#### Call by reference, Call by Value

[참조자료](http://twoday.tistory.com/entry/call-by-value%EC%99%80-call-by-reference-%EC%A0%9C%EB%8C%80%EB%A1%9C-%EC%95%8C%EA%B8%B0)
[각 언어별 방식](https://okky.kr/article/303162?note=1005863)
[예제](http://trypsr.tistory.com/74)


##### Call by reference 