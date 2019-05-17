
## mockup test, mock 이란?

목업 테스트 라는 말을 들어본 적이 있을 수 있다(없을 수도 있다)...

mock 오브젝트는 **테스트 더블Test Double**중 하나이다. 실제 객체가 사용되는 모습을 따라해서 만든 카피품이라고 생각하면 된다.

보통 mock 객체를 다른 객체에서 사용하는 테스트를 하는 데 사용한다.

보통 Stub와 엮여 언급된다. Test Double 에 대해 알아보면 된다.

### Test Double 은 또 뭐임?

테스트 더블이란 -> 테스트시 실제 객체를 대신할수 있는 객체를 뜻함.

### Stub

로직은 없고 원하는 값을 반환한다. 예를 들어 항상 같은 값을 반환하는 객체를 생각할 수 있다.

### Mock

Mock은 이 객체를 이용하면 > 이런 메서드가 호출된다 라고 이해하면 쉽다. 객체 사이의 상호작용을 테스트하는 데 사용한다.

### Fake

실제 구현처럼 동작하게 간단히 만든다. 

자바 계열에선 Junit, Mockito를 사용해서 위와 같은 테스트를 진행한다. 
의존성 제거 및 테스트 속도 개선의 효과가 있다.

go언어에서는 어떻게 mock 테스트를 할까?
golang/mock 패키지 안에 [gomock](github.com/golang/mock/gomock)이나 [mockgen](github.com/golang/mock/mockgen)에서 참고할 수 있다.
