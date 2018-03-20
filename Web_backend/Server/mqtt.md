# mqtt 란?

> mqtt : Message Queuing Telemetry Transport. ISO standard publish-subscribe-based messaging protocol

항상 뭔가의 프로토콜이라고 생각하고 넘어가던 mqtt의 정의와 활용에 대해 알아보자.

## 1. MQTT

애초에 MQTT는 발행/구독(Publish/Subscribe) 방식의 메세지 큐(message queue) - [큐 : 자료구조에서 First In First Out의 그 큐] 이다. 원격 검침(telemetry)영역에 사용하기 위해 개발되었다. Iot(사물인터넷)에서도 많이 활용한다고 한다. 이런 장비들 대부분이 소형이며 통신 대역폭과 전원이 한정적인 환경에서 동작하므로, 이는 배터리 용량이 제한적이고 통신 품질을 일정 수준으로 유지하기 어렵다는 점에서 스마트폰 환경과 매우 유사하다. 이러한 제한적인 환경을 고려하여 디자인 되었기 때문에 MQTT 프로토콜은 최적의 프로토콜로 주목받고 있다.

#### MQTT란?  publish subscribe 방식 메시지 큐
#### MQTT는 왜 쓰나요? 아주 경량이고 스케일하기 쉬운데다 확장성이 좋기 때문
#### MQTT는 어디에 쓸까? 검침, 오토모바일, 스마트폰, Iot기기 등등

> MQTT 프로토콜 설계의 의도

-   프로토콜이 차지하는 모든 면의 **리소스 점유(footprint)를 최소화**
-   느리고 품질이 낮은 **네트워크의 장애와 단절에 대비**
-   클라이언트 애플리케이션 동작에 **자원 활용이 극히 제한적임을 고려**
-   다수의 클라이언트 연결에 접합한 **Publish/Subscribe 네트워크 채용**
-   **신뢰성 있는 메시징**을 위한 QoS(Quality of Service) 옵션 제공.
-   **개방형 표준 메시징 프로토콜을 지향** – 제 3자(3rd Party) 기기 제조업체와 소프트웨어 개발업체의  
    용이한 도입, 적용을 유도

> MQTT의 특징

- 오픈소스 (http://www.mqtt.org)
- 오버헤드 최소화
-  단순한 pub / sub 메시지 체계를 가짐
-  Quality of Service(QoS) 제공으로 메시지 신뢰성 유지
- 연결 오류 시에도 보정할 수 있는 자체 기능도 있음

> MQTT 장점

- 주기적으로 폴링(polling)하는데 비해 베터리 소모량 및 패킷 전송량이 적다.

- 메시지 전송에 대한 신뢰성을 보장하기 위해 3단계 QoS 제공

> MQTT 단점
 
- 다수의 Publisher/Subscriber 가 하나의 공통된 Topic을 공유할 경우 메시지의 순차적 처리를 보장하지 못한다 (ex:그룹채팅 / 네트워크 환경에 의한 예외현상)

브로커로는 Mosquitto, HiveMQ, RabbitMQ, IBM MQ등이 있다. 나는 RabbitMQ만 이게 뭔지 모르는 상태로 사용해 봤지만(제일 럴닝커브가 낮고 꽁짜여서...), 인터넷에서 찾아보니 IBM MQ 추천이 많다. 

## 2. MQTT 활용

#### Publish / Subscribe

MQTT는 브로커된 publish / subscribe 패턴을 구현하고 있다. pub/ sub가 클라이언트에서 메시지를 받아다 전송하는 ("publisher")를 분리해 내는 일을 한다. 이게 무슨 뜻이냐면, publisher랑 subscriber가 각자가 존재하는지 전혀 모르고 있다는 뜻이다. 클라이언트끼리는 서로 모르지만, 메시지 브로커를 알고 있는 것이다. 그리고 이 메시지 브로커가 받아오는 모든 메시지를 필터링 하는 구조이다.

sender랑 receiver는 세가지 차원으로 분류해볼 수 있는데 다음과 같다.

-  공간 분리(Space decoupling) : publisher랑 subscriber가 서로 알 수 없도록 합니다(예:IP주소 및 포트별).  
- 시간 디커플링(Time decoupling) : publisher와 subscriber를 동시에 연결할 필요가 없습니다.  
- 동기화 디커플링(Synchronization decoupling) : 메시지를 게시하거나 수신하는 동안 두 구성 요소에 대한 작업이 멈추지 않습니다

#### MQTT Message Types

[mqtt 메시지 종류들](https://dzone.com/refcardz/getting-started-with-mqtt?chapter=4)

#### Qos 레벨

1. QoS 0 - 최대 1회 전달, 전송확인 X, 유실가능성이 있다.
2. QoS 1 - 최소 1회 전달, 전송확인 O, 중복전달 가능성이 있다.
3. QoS 2 - 4단계 HandShaking을 통해 정확히 1회 전달 (종단 간 지연이 늘어나게 되는 단점이 있다)
