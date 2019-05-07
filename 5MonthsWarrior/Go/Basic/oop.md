# 객체지향

Go는 객체지향을 완전히 지원하지 않는다.

## 다형성

> 메서드가 호출되었을 때 어떤 자료형이냐에 따라 다른 구현을 할 수 있게 한다

```
package polymorph

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

func TotalArea(shapes []Shape) float32 {
	var total float32
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type RectangleCircum struct{ Rectangle }

func (r RectangleCircum) Circum() float32 {
	return 2 * (r.Width + r.Height)
}

func NewRectangle(width, height float32) *RectangleCircum {
	return &RectangleCircum{Rectangle{width, height}}
}

type WrongRectangle struct{ Rectangle }

func (r WrongRectangle) Area() float32 {
	return r.Rectangle.Area() * 2
}
```

go에서는 다형성을 위해 인터페이스를 이용한다.

## 상속

상속을 통해 객체지향에서는 클래스의 구현을 재사용한다.

```
// 남이 만든ㄷ거

type Rectangle struct {
  Width, Height float32
}

func (r Rectangle) Area() float32{
  return r.Width * r.Height
}

// 내가 추가한거

type RectangleCircum struct { Rectangle }

func (r RectangleCircum) Circum() float32{
  return 2*(r.Width + r.Height)
}
```

go 언어에서는 오버라이딩은 어떻게 구현할까? 구조체 내장을 이용한다.

```
type WrongRectangle struct { Rectangle }

func (r WrongRectangle) Area() float32 {
  return r.Rectangle.Area() * 2
}
```

캡슐화(객체 안 정보를 바깥에 숨기는 것) 는 어떻게 표현할까? Go에는 상속이 없기 때문에 protected는 없지만 public, private는 있다.

이런 경우에는 internal pacakage 를 이용한다.

pacakge 경로에 internal을 넣으면, 패키지 경로 안 internal 이 있는 범위에서만 참조가 가능하다.

