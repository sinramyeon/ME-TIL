#### hoisting

---

[참고자료](http://mohwaproject.tistory.com/entry/%EC%9E%90%EB%B0%94%EC%8A%A4%ED%81%AC%EB%A6%BD%ED%8A%B8-%ED%95%A8%EC%88%98-%ED%98%B8%EC%9D%B4%EC%8A%A4%ED%8C%85)
[참고자료2](http://chanlee.github.io/2013/12/10/javascript-variable-scope-and-hoisting/)

---

###### 자바스크립트의 함수 범위

자바스크립트는 함수 수준(function level)의 범위를 가진다.
함수 내 정의된 변수는 지역 범위를 가지며,
해당 함수와 내부 함수에서 접근이 가능하다.

Ex)

```
var name = "HERO"

function hoistingTest(){
  var name = "HERO0926"
  
  console.log(name); // HERO0926이 출력된다.

}

console.log(name); //HERO가 출력된다.
```

지역변수는 함수 내에서 전역변수보다 높은 우선순위를 가짐.

Ex2)

```
for (var i = 1 ; i <=10 ; i++) {
  console.log(i); //10까지 출력
}

function aNumber(){
  console.log(i); #i
}

aNumber(); // 11 출력
```

##### 호이스팅 Hoisting

변수의 정의가 그 범위에 따라 선언과 할당으로 분리.
변수가 함수 내에서 정의되었을 경우, 선언이 함수의 최 상위로
함수 바깥에서 정의되었을 경우, 전역 컨테스트이 최 상위로 변경됨.

Ex)
```
function showName() {
     console.log("First Name : " + name);
     var name = "Ford";
     console.log("Last Name : " + name);
}
showName();

// First Name : undefined
// Last Name : Ford
// First Name이 undefined인 이유는 지역변수 name이 호이스트 되었기 때문입니다.

function showName() {
     var name; // name 변수는 호이스트 되었습니다. 할당은 이후에 발생하기 때문에, 이 시점에 name의 값은 undefined 입니다.
     console.log("First name : " + name); // First Name : undefined
     name = "Ford"; // name에 값이 할당 되었습니다.
     console.log("Last Name : " + name); // Last Name : Ford
}

// 다음 두 변수와 함수는 myName으로 이름이 같습니다.
var myName; // string
function myName() {
     console.log("Rich");
}
// 함수 선언은 변수명을 덮어 씁니다.
console.log(typeof myName); // function

//반대로 변수에 값이 할당될 경우에는 변수가 함수선언을 덮어 씁니다.

var myName = "Richard";
function myName() {
     console.log("Rich");
}
console.log(typeof myName); //string
```


Ex2)

```
console.log(internalFunction);// 함수가 Hoisting 됐을뿐, 아직 Execution Context는 생성되지 않았으므로, 에러가 발생한다.

functionDeclaration();// 함수가 성공적으로 호출되며 Function Declaration가 프린트 된다.

console.log();// Undefined 가 프린트 된다. 변수에 "vart"라는 값이 어싸인되는것은 아래 코드가 읽히는 시점에서기 때문.

function functionDeclaration() {
    console.log("Function Declaration");
    
    function internalFunction(){
    }
}

var t = "vart";

functionExpression(); // 에러가 발생

console.log(novalue); // Context에서 선언된적 없는 변수인 novalue를 호출할 수가 없어 에러가 발생

var functionExpression = function() {
    console.log("Function Expression");
}
```

