# 자바스크립트 클로저

> 클로저는 독립적인 자유 변수를 참조하는 함수들이다.

---

### 문법
```
function init(){
	var name = "Mozilla" # 지역변수
	function displayName() { # 내부 함수인 클로저
	alert(name)
	}
	displayName();
}

init()
```

init이 생성한 지역 변수 name과 함수 displayName().
displayName()은 부모 함수에서 선언된 변수 name을 참조

```
function makeFunc() {
  var name = "Mozilla";
  function displayName() {
    alert(name);
  }
  return displayName;
}

var myFunc = makeFunc();
myFunc();
```

displayName() 안 내부 함수 실행 전 외부 함수로부터 반환

```
function makeAdder(x) {
  return function(y) {
    return x + y;
  };
}

var add5 = makeAdder(5);
var add10 = makeAdder(10);

console.log(add5(2));  // 7
console.log(add10(2)); // 12
```

add5, add10이 클로저 > 함수 본문을 공유하지만 서로 다른 변수를 기억

### 자바스크립트의 모든 함수는 클로저