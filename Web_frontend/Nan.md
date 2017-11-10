# Nan

---

## global NaN 속성(property) :  Not-A-Number(숫자가 아님)


Property attributes of NaN
Writable  |	no
--------- | ----
Enumerable | no
Configurable | no

`NaN`은 global object(표준 내장 객체)의 속성입니다.

NaN의 초기값은 Not-A-Number 즉, Number.NaN 값과 같습니다.

최신 브라우저에서, NaN은 설정불가(non-configurable), 쓰기불가(non-writable) 속성입니다. 이는 그 경우가 아닐 때라도, 재정의(overriding)를 피합니다.

프로그램에서 NaN을 사용하기는 조금 드뭅니다. 그것은 Math 함수가 실패(Math.sqrt(-1)))하거나 숫자를 파싱하려 한 함수가 실패(parseInt("blabla"))했을 때 반환되는 값입니다.

---

### NaN에 대한 테스트

NaN은 다른 NaN 값을 포함하여 다른 어떤 값과 같지 않음(`==`, `!=`, `===` 및 `!==`를 통해)을 비교합니다. 값이 NaN인지 가장 명확하게 결정하기 위해 Number.isNaN() 또는 isNaN()을 사용하세요. 아니면 자체 비교를 하세요: NaN 및 NaN만, 자신과 같지 않음을 비교합니다.

```
NaN === NaN;        // false
Number.NaN === NaN; // false
isNaN(NaN);         // true
isNaN(Number.NaN);  // true

function valueIsNaN(v) { return v !== v; }
valueIsNaN(1);          // false
valueIsNaN(NaN);        // true
valueIsNaN(Number.NaN); // true
```

Nan 비교는 `Number.isNaN()` 이나 `global isNan()`으로 하는게 좋습니다. Number.isNan() 메서드는 전달된 값이 NaN인지 결정합니다. 오리지널 global isNaN()의 더 강력한 버전.

```
Number.isNaN(NaN);        // true
Number.isNaN(Number.NaN); // true
Number.isNaN(0 / 0)       // true

// 예를 들면 이들은 global isNaN()으로는 true가 됐을 것임
Number.isNaN("NaN");      // false
Number.isNaN(undefined);  // false
Number.isNaN({});         // false
Number.isNaN("blabla");   // false

// 이들 모두 false를 반환함
Number.isNaN(true);
Number.isNaN(null);
Number.isNaN(37);
Number.isNaN("37");
Number.isNaN("37.37");
Number.isNaN("");
Number.isNaN(" ");
```

---

### Nan 의 짜증남

`typeof NaN === 'number' // true `

typeof로는 NaN과 숫자 구별도 안되고 자기와 비교도 안된다.

```
NaN === NaN    // false 
NaN !== NaN    // true 
```

특히 typeof에서는 null도 몰라본다.

`typeof null     // object `

그러므로 null값 비교는 `===` 연산자로 하는것이 좋다.

이런 일을 필하려면 [JSLint](http://www.jslint.com/)를 자주 사용하는 수 밖에 없다.

