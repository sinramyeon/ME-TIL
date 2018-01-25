# 가운데 정렬

## 방법 1:position

> poa(position: absolute),pof(position: fixed),

*(단점: 객체의 width, height가 고정이여야 함)*

### 방법 1.1
마이너스 마진값으로 정렬

```
.target {
    position: absolute
    left: 50%;
    top: 50%;
    margin: -height/2 -width/2;
}
```

### 방법 1.2
calc사용

(장점: 1.1보다 css줄을 한줄 줄일 수 있다.)

(단점: calc를 사용하면 계산하기 때문에 조금 더 느려진다고 함)

```
.target {
    position: absolute
    left: calc(50% - (width / 2));
    top: calc(50% - (height / 2));
}
```

### 방법 1.3
auto 마진값으로 정렬

```
.target {
    position: absolute
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: width;
    height: height;
    margin: auto;
}
```

## 방법 2:display: inline-block*
dib(display: inline-block)의 vam(vertical-align: middle)방법

(장점: 사용법이 간단, 모듈로 만들어놓고 적용만 하면 더 간단.)

(단점 : 돔이 한개 더 필요함)

```
.parent {
    white-space: no-wrap;
}
 
.parent:after {
    display: inline-block;
    height: 100%;
    vertical-align: middle;
    content: "";
}
.target {
    display: inline-block;
    vertical-align: middle;
}

```

### 원리 간단히 알아보기

dib & vam은 두개의 높이를 비교하며 정렬되는 방법입니다.

전체 높이를 재는 자(rule)가 꼭 한개 필요합니다.

정렬할 객체 1개여야하고, 비교할 객체 1개가 추가 필요한 겁니다.


```
<div class="parent">
    <div class="target">
        <span>가나</span><span>다라</span>
    </div>
    <div class="rule"></div>
</div>
```



가운데 정렬하고픈 target이 있다면

그 옆에 rule과 같은 돔이 한개 필요합니다. 그리고 그 rule은 정렬할 기준(.parent)과 동일한 높이값을 가져야합니다.

```
.parent {
    white-space: nowrap;
}
 
.target {
    display: inline-block;
    vertical-align: middle;
}
 
.rule {
    display: inline-block;
    height: 100%;
    vertical-align: middle;
}
```


상단처럼 제작하면 .target은 어디에 내 놓아도 가운데 정렬인 상태가 됩니다.

이제 불필요한 돔

rule을 가상 선택자로 만들면

가운데정렬할때마다 돔을 건드려야하는 문제점이 해결되며

사용하기에 조금 더 편해집니다.

```
.parent {
    white-space: no-wrap;
}
 
/* .parent:after === .rule */
.parent:after {
    display: inline-block;
    height: 100%;
    vertical-align: middle;
    content: "";
}
 
 
.target {
    display: inline-block;
    vertical-align: middle;
}
```

그리고 이제 scss로 만들어 두면 바텀정렬, 미들 정렬이 훨씬 수월해집니다.

아래는 extend로 사용하는 방법으로 작성해보았습니다.

```
%vamBoxParent {
    white-space: nowrap;
 
    > * {
        display: inline-block;
        vertical-align: middle;
        white-space: normal;
    }
 
    &:after {
        display: inline-block;
        vertical-align: middle;
        height: 100%;
        min-height: inherit;   
        content: "";
    }
}
 
 
%vabBoxParent {
    white-space: nowrap;
 
    > * {
        display: inline-block;
        vertical-align: bottom;
        white-space: normal;
    }
 
    &:after {
        display: inline-block;
        vertical-align: bottom;
        height: 100%;
        min-height: inherit;
        content: "";
    }
}
 
//이제 필요한 객체에 extend로 가져가서 적용.
.parent {
    @extend %vamBoxParent;
}

```

## 방법 3:table

display: table-cell에 vam속성 적용.



### 방법 3.1
기본적인 table방법

(단점: 돔이 최소 한개 더 필요함, table특성을 따름)

```
/* parent = table-row || table */
 
 
.target {
    display: table-cell;
    vertical-align: middle;
}
```

### 방법 3.2
table-cell만 사용하기


```
/* parent = anything */
 
 
.target {
    display: table-cell;
    width: inherit;
    vertical-align: middle;
}

```

#### 원리 간단히 알아보기

table-cell은 최소 table or table-row가 부모로 존재해야합니다.

그런데 하단처럼 .taget이 only-child라면

```
<div class="parent">
    <div class="target">
        <span>가나</span><span>다라</span>
    </div>
</div>
```

.target에만 table-cell로 정렬할 수 있습니다.

```
.parent {
     
}
 
.target {
    display: table-cell;
    width: 100%;
    vertical-align: middle;
}
```

그런데 저기에 width: 100%가 문제가 됩니다.

table-cell만 쓸 경우

width 퍼센트에 정상 반응하지 않았습니다.



그래서 width: inherit을 쓰니 100%와 동일하게 적용되었습니다.

```
.parent {
     
}
 
.target {
    display: table-cell;
    width: inherit;
    vertical-align: middle;
}
```



## 방법 4:flex*
(장점: 돔 추가하지 않아도 됨)

(단점:하위기종문제, 안드4.4이하, 피시의 경우 사용하는데 조금 더 신중해야합니다.)

```
.parent {
    display: flex;
    align-items: center;
}
```

그래서 하위기종에만 반응할 수 있도록 추가 보안 소스를 상단에 작업해 두면 보안이 가능합니다.

(align-items가 먹은 상태에서 lh를 무시하는것 확인했습니다. )

```
.parent {
    display: block;
    display: flex;
    line-height: height;
    align-items: center;
}
```

## 방법 5:line-height*
lh(line-height)변경

(장점: 돔을 추가하지 않아도 됨)

(단점 : 폰트만 사용가능, 한줄 이상일 경우 사용불가함)



### 방법 5.1
lh를 가운데 정렬할 높이값 만큼 적용

```
.target {
    height: height;
    line-height: height;
}
```


### 방법 5.2
(단점: oh(overflow: hidden)일 경우 사용불가)

lh를 0으로 주고 padding값을 높이값의 반으로 줌

```
.target {
    height: height;
    padding: height/2 0;
    line-height: 0;
    box-sizing: border-box;
}
```

때에 따라 필요한 방법을 써야 합니다.
