# ES6

---

## let

### 블럭 범위 내에서 변수 선언

```
var x = 3;
if(true) {
	let x = 5;
	console.log(x);
    // 5 출력
}
console.log(x) //3 출력
```

### 블럭 내 중복 선언 시 문법 에러 발생

```
(function () {
  let x;
  let x; //-> Uncaught SyntaxError: Identifier 'x' has already been declared
})();
```

### Hoisting 적용 안됨

## 스코핑

### JS는 정적 함수 스코핑을 하는 언어

```
function foo() {
	
	var x = 1;
	{
		var x = 2;
	}
	return x; //var x = 2 를 할당문으로 인식
}

var x = foo();
console.log(x); //2 출력
```

### let 이용시
```
function foo_() {
    let x = 1;
    {
        let x = 2;
    }
    return x;
}
 
var x_ = foo_();
console.log(x_);
// 1
```

### let을 var대신 모두 사용하는것이 좋다.