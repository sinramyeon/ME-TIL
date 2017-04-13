###디버깅 꿀팁

<hr/>

[참고한 링크](http://stackoverflow.com/questions/2342280/difference-between-logger-info-and-logger-debug)

<hr/>

#### 무언가 문제가 있을 때 디버깅을 하려면 어떻게 할까?
```
 protected static Log logger = LogFactory.getLog(Settings.class);
 logger.info;
 logger.debug;
```
*logger 와 Sysout 을 적극 이용하자!*

이 기능이 오류가 난다 하면 이 기능이 어떤 메서드나 액션에서 벌어지는지 확인하고, 이 기능을 호출하는 곳은 어디인지 찾아보자.
어떤 메서드를 어디서 호출하는지 찾으려면 메서드 이름을 드래그한 후 우클릭해서 Reference -> Workspace 로 찾는다.

!(DEBUG_ref1.png)

이렇게 메서드마다 타고 가서 log를 찍어서, 어디서부터 어디까지는 실행이 되고 안 되는지, 어떤 파라미터가 들어오고 어떤 결과가 생기는지 확인한다.

#### 이클립스 내에서 하자!

디버그 모드 적극 활용이 필요하다.

![참조](http://cfile22.uf.tistory.com/image/232D57485570C9DD0DF53D)

Run 시 디버그를 하는 모드로 시작해 보자.
의심되는 소스 좌측을 더블클릭해 break point를 생성하면, 그 곳에 오면 멈춘다.

![참조](http://cfile21.uf.tistory.com/image/233041485570C9E40CB473)

1) Skip All Breakpoints : 모든 브레이크 포인트 건너뜀
2) Resume(F8키) : 다음 브레이크포인트까지 진행함
3) Suspend : 쓰레드를 일시 정지하며 현재 수행문에 지정한 것과 같음
4) Terminate : 쓰레드 종료
6) Step Into(F5키) : 한단계 진행하는데 다음 라인이 함수 안이면 함수 안으로 들어감.
7) Step Over(F6키) : 함수 호출을 지나치고 현재 위치에서 한 단계 진행
8) Step Return(F7키) : 현재 함수 끝까지 바로 가서 리턴한 후 함수 호출부로 되돌아 감
9) Drop to Frame : 선택한 스택 프레임의 첫 행으로 이동. 처음부터 다시 하고자 할 때 *해당함수 처음으로 이동하니 놓친경우 유용히 쓸 수 있음*
10) Use Step Filters(Shift+F5) : 스텝 필터링

[초보 개발자를 위한 디버깅 하는 법](https://okky.kr/article/272227)

#### logger.info 와 logger.debug의 차이는?

로깅 환경에 달려 있습니다.
로그 레벨은 아래와 같습니다.
> DEBUG < INFO < WARN < ERROR < FATAL

FATAL 가장 크리티컬한 에러
ERROR 일반 에러
WARN 주의 하세용
INFO 정보
DEBUG 상세한 일반정보

로깅 레벨을 무엇으로 설정하냐에 따라 그 이상 레벨만 로깅하므로 나오는 로그가 다릅니다.
> EX) Log 레벨을 info 로 설정하면 info, warn, error, fatal 은 출력되나 debug 는 출력되지 않습니다.

Log Level | Summary
------------ | -------------
ALL | The ALL has the lowest possible rank and is intended to turn on all logging.
DEBUG | The DEBUG Level designates fine-grained informational events that are most useful to debug an application
ERROR | The ERROR level designates error events that might still allow the application to continue running.
FATAL | The FATAL level designates very severe error events that will presumably lead the application to abort.
INFO | The INFO level designates informational messages that highlight the progress of the application at coarse-grained level.
OFF | The OFF has the highest possible rank and is intended to turn off logging.
TRACE | The TRACE Level designates finer-grained informational events than the DEBUG.
TRACE_INT | TRACE level integer value.
WARN | The WARN level designates potentially harmful situations.

