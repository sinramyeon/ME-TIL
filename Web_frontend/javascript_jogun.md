### 조건문

---

#### 삼항 연산자

(불 표현식)?(참일 때 실행하는 문장) : (거짓일 때 실행하는 문장)

```
<script>

//변수 선언
var input = prompt("숫자를 입력해 보세요...", "");
var number = Number(input);

//조건문
(number > 0) ? alert("자연수입니다.") : alert("자연수가 아닙니다.");

</script>
```

#### 짧은 조건문

(불 표현식) || (불 표현식이 거짓일 때 실행할 문장)
(불 표현식) && (불 표현식이 참일 때 실행할 문장)

```
<script>
ture || alert("실행이..안됨...");

//좌변이 거짓이므로 우변이 참인지 거짓인지 검사
false || alert("실행이...됨...");
</script>
```

```
<script>
//변수 선언
var input = Number(prompt("숫자를 입력해야 합니다.", "숫자");

//조건문
input % 2 == 0 || alert("짝수가 아닐 때, 즉 홀수.");
input % 2 == 0 && alert("짝수일 때.");
</script>
```