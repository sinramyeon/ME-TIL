Go 로깅 시스템
(Go in Action : Standard library > Log package 편 번역)입니다.

고를 쓰면서 제일 큰 불만은 대체 왜 이름을 고로 지었냐는 거다. 당최 검색하기 힘들어 죽겠다. 누구는 golang 누구는 go 누구는 google go로 쓰고 자빠졌으니 분노만이 일어난다. 시노니움을 프로그래밍 언어로 차용한 구글놈들을 다 때려주고 싶다....



어쨌든, go, golang, google go, 고, 고언어... 를 사용하여 로그를 짜보겠다.



8.2.1 로그 패키지
자신만의 로그를 만들기 위해서 고에서 기본으로 제공하는 log 패키지를 들여다봅시다. 로그를 찍는 이유는, 프로그램이 무얼 하고 있고, 언제 어디서 무슨 일이 일어나는지 추적하기 위해서입니다. 로그란, 한 줄 한 줄 마다 몇가지 설정과 함께 정보를 찍는 것입니다.



8.2 Sample trace line TRACE: 2009/11/10 23:00:00.000000 /tmpfs/gosandbox-/prog.go:14: message



8.2에서 로그 패키지에서 제공하는 샘플 로그를 확인해볼 수 있습니다. 이 로그는 'prefix'로 현재 시간을 표시해주고 있고, 로그가 찍힌 소스 코드를 나타내고 있으며, 몇 번째 라인에서 발생한지와, 무슨 로그 메시지가 나왔는지까지 포함하고 있습니다. 이러한 로그를 찍기 위해 log 패키지를 설정하는 법을 알아봅시다.



8.3

// This sample program demonstrates how to use the base log package.
package main

import (

"log"
)

func init() {
 log.SetPrefix("TRACE: ")
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
 }

 func main() {
 // Println writes to the standard logger.
 log.Println("message")

 // Fatalln is Println() followed by a call to os.Exit(1).
 log.Fatalln("fatal message")

 // Panicln is Println() followed by a call to panic().
 log.Panicln("panic message")
 }


위 프로그램을 실행해 보면, 8.2와 비슷한 로그를 찍어보실 수 있으실 겁니다.



8.4

func init() {
 log.SetPrefix("TRACE: ")
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
 }


init()메서드가 보이시죠? 

이 메서드는 main()함수 실행 전 프로그램을 초기화해주는데에 사용됩니다. 보통 로그 설정을 하는 것도 init()안에서 합니다. 그렇게 하면 프로그램이 시작할 때 로그 설정을 바로 할 수 있기 때문입니다. 확인해보면 SetPrefix()로 로그 맨 앞을 꾸미는 접두사를 만들어주고 있는데, 모든 로그가 찍힐 때 맨 앞에서 TRACE: 를 확인해볼 수 있습니다. 보통 Prefix는 대문자로 씁니다. log 패키지 안에는 몇 가지 flag도 있는데, 각 로그 라인마다 다른 정보를 관리할 수 있습니다. 아래에서 flag를 어떻게 사용하는지 확인해 보세요.



8.5

const (
// Bits or'ed together to control what's printed. There is no control
// over the order they appear (the order listed here) or the format
// they present (as described in the comments). A colon appears after
// these items:
// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
// the date: 2009/01/23
Ldate = 1 << iota
// the time: 01:23:23
Ltime
// microsecond resolution: 01:23:23.123123. assumes Ltime
Lmicroseconds
// full file name and line number: /a/b/c/d.go:23
Llongfile
// final file name element and line number: d.go:23
// overrides Llongfile
Lshortfile
// initial values for the standard logger
LstdFlags = Ldate | Ltime
)


8.5에서 log 패키지 내 소스 코드를 바로 확인해볼 수 있습니다. 플래그는 상수로 지정되어 있고, Ldate라는 상수를 보시면 특별한 문법으로 지정되어 있는 것을 확인해볼 수 있습니다.



8.6

// the date: 2009/01/23
Ldate = 1 << iota


iota는, 상수 지정시 특별한 목적을 갖게 됩니다. 컴파일러가 모든 상수 구문을 블록이 끝날 때까지 아니면 새 할당문이 발견될 때까지 중첩시키기 때문입니다.(원문 : It instructs the compiler to duplicate the expression for every constant until the block ends or an assignment statement is found.)

 iota 키워드의 또 다른 기능은 앞의 각 상수에 대한 iota의 값이 1 씩 증가한다는 것입니다. 초기 값은 0입니다. 더 자세히 보겠습니다.



8.7

const (
Ldate = 1 << iota // 1 << 0 = 000000001 = 1
Ltime // 1 << 1 = 000000010 = 2
Lmicroseconds // 1 << 2 = 000000100 = 4
Llongfile // 1 << 3 = 000001000 = 8
Lshortfile // 1 << 4 = 000010000 = 16
...
)


8.7은 상수 선언시 일어나는 일을 보여줍니다. << 연산자는 값이 나타내는 비트의 왼쪽 비트 시프트를 수행합니다. 각각의 경우에 값 1에 대한 비트 패턴은 왼쪽 iota에서 시프트됩니다. 이것은 각 상수에 고유 한 비트를 주는 효과가 있습니다. 플래그 연산에 제격입니다. LstdFlags 상수는 각 상수에 고유한 비트 위치가 있다는 것을 보여줍니다.



8.8

const (
...
LstdFlags = Ldate(1) | Ltime(2) = 00000011 = 3
)


8.8에서 LstdFalgs 상수가 할당 연산자를 사용해서 iota 체인을 끊는 것을 확인해볼 수 있습니다. 파이프 연산자 (|)가 함께 사용되거나 비트가 함께 사용된것을 보아서 LstdFlag에 3일 지정된것을 알수 있습니다. O'ring 비트가 조이닝 비트와 동일하므로 각 개별적으로 설정된 비트는 최종 값으로 표현됩니다. 비트 1과 2가 함께 합쳐져 3을 이룹니다.

이번엔 log 플래그를 설정하는 방법을 다시 살펴 보겠습니다.



8.9

func init() {
 ...
 log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
 }


여기에서는 Ldate, Lmicroseconds 및 Llongfile 플래그를 함께 묶어 해당 연산의 값을 SetFlags 함수에 넣었습니다. 이 플래그들은 함께 묶였을 때 값 13과 비트 4, 3 및 1 (00001101)을 나타냅니다. 각 상수는 개별 비트를 나타내며, 파이프 연산자를 사용하여 플래그를 함께 조인 해 적용하고 싶은 모든 로그 옵션을 넣을 수 있습니다. 로그 패키지가 그 후 입력한 정수값을 받아들여 올바른 비트를 적용하도록 설정된 비트를 검사합니다.

로그 패키지가 초기화되면 main() 함수를 확인해서 어떻게 로그를 쓰는지 알아봅시다.



8.10

func main() {
// Println writes to the standard logger.
log.Println("message")
// Fatalln is Println() followed by a call to os.Exit(1).
log.Fatalln("fatal message")

// Panicln is Println() followed by a call to panic().
log.Panicln("panic message")
}


Listing 8.10은 세 가지 다른 함수를 사용하여 로그 메시지를 작성하는 방법을 보여준다 : Println, Fatalln, Panicln. 

이 함수들은 포맷된 버전도 가지고 있고, ln 대신 f를 써서 호출합니다.

 Fatal 계열은 로그 메시지를 작성한 다음 os.Exit (1) 함수 호출을 하여 프로그램을 종료하는데에 사용합니다. 

Panic 은 로그 메시지를 작성한 다음 패닉을 발생시키는 데 사용됩니다. 복구되지 않으면 프로그램을 종료하고 스택트레이스를 남깁니다. 

Print계열은 로그 메시지를 작성하는 표준 방법입니다. 로그 패키지에 대한 좋은 점 중 하나는 로거가 멀티고루틴에 안전하다는 것입니다. 

이 말은 여러 goroutine이 같은 로거 값에서 이러한 함수를 동시에 호출 할 수 있음을 의미합니다. 이제 로그 패키지를 사용하고 구성하는 방법을 알았으니 사용자 정의 로거를 작성하여 작성할 수 있는 여러 로깅 레벨에 대해 알아봅시다.



8.2.2 사용자 정의 로거


사용자 정의 로거를 작성하려면 사용자 고유의 로거 유형 값을 작성해야합니다. 생성 한 각 로거는 고유 한 대상에 맞게 구성 할 수 있으며 자신만의 접두사와 플래그를 가집니다. 예제를 확인해 보세요.



8.11

01 // This sample program demonstrates how to create customized loggers.
02 package main
03
04 import (
05 "io"
06 "io/ioutil"
07 "log"
08 "os"
09 )
10
11 var (
12 Trace *log.Logger // Just about anything
13 Info *log.Logger // Important information
14 Warning *log.Logger // Be concerned
15 Error *log.Logger // Critical problem
16 )
17
18 func init() {
19 file, err := os.OpenFile("errors.txt",
20 os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
21 if err != nil {
22 log.Fatalln("Failed to open error log file:", err)
23 }
24
25 Trace = log.New(ioutil.Discard,
26 "TRACE: ",
27 log.Ldate|log.Ltime|log.Lshortfile)
28
29 Info = log.New(os.Stdout,
30 "INFO: ",
31 log.Ldate|log.Ltime|log.Lshortfile)
32
33 Warning = log.New(os.Stdout,
34 "WARNING: ",
35 log.Ldate|log.Ltime|log.Lshortfile)
36
37 Error = log.New(io.MultiWriter(file, os.Stderr),
38 "ERROR: ",
39 log.Ldate|log.Ltime|log.Lshortfile)
40 }
41
42 func main() {
43 Trace.Println("I have something standard to say")
44 Info.Println("Special Information")
45 Warning.Println("There is something you need to know about")
46 Error.Println("Something has failed")
47 }


8.11에서 Trace, Info, Warning, and Error의 4가지의 다른 로그 수준을 지정하는 것을 볼 수 있습니다. 각 로거 수준마다 다른 수준의 중요도를 가진 로그를 보여줍니다. 어떻게 작동하는지 알아봅시다.



12, 13
25 Trace = log.New(ioutil.Discard,
26 "TRACE: ",
27 log.Ldate|log.Ltime|log.Lshortfile)
28
29 Info = log.New(os.Stdout,
30 "INFO: ",
31 log.Ldate|log.Ltime|log.Lshortfile)
32
33 Warning = log.New(os.Stdout,
34 "WARNING: ",
35 log.Ldate|log.Ltime|log.Lshortfile)
36
37 Error = log.New(io.MultiWriter(file, os.Stderr),
38 "ERROR: ",
39 log.Ldate|log.Ltime|log.Lshortfile)
새 로그를 만들고 싶다면 로그 패키지 내 New 메서드를 사용합니다. New 함수는 새롭게 만들어진 값 주소를 리턴해줍니다. New를 호출할 때는 몇 가지 파라미터를 넘겨줘야 합니다.

14
// New creates a new Logger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *Logger {
return &Logger{out: out, prefix: prefix, flag: flag}
}


첫 번째 파라미터는 로그를 어디에 쓸 것인지 종착지를 입력해 줍니다. io.Writer인터페이스를 구현한 값을 제공합니다. 두 번째 파라미터는 prefix입니다. 로그 플래그가 세 번째 파라미터입니다.



8.15, 16

25 Trace = log.New(ioutil.Discard,
26 "TRACE: ",
27 log.Ldate|log.Ltime|log.Lshortfile)
//The Discard variable has some very interesting properties.
// devNull is a named type using int as its base type.
type devNull int
// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
var Discard io.Writer = devNull(0)
// Implementation of the io.Writer interface.
func (devNull) Write(p []byte) (int, error) {
return len(p), nil
}


8.16은 Discard 변수의 선언과 구현을 보여준다. Discard 변수는 devNull 유형의 0 값을 가진 io.Writer인터페이스로 선언됩니다. 이 제공됩니다. 이 변수에 쓰여진 것은 모두 버려집니다.(Anything written to this variable is discarded based on the implementation of the Write method for the devNull type.) Discard 를 사용하면 출력 할 때 로깅 수준을 비활성화할 수 있습니다.

Info와 Warning 로그는 stdout을 목적지로 쓰고 있습니다.



8.17, 18

29 Info = log.New(os.Stdout,
30 "INFO: ",
31 log.Ldate|log.Ltime|log.Lshortfile)
32
33 Warning = log.New(os.Stdout,
34 "WARNING: ",
35 log.Ldate|log.Ltime|log.Lshortfile)
The declaration of the Stdout variable is also interesting.
// Stdin, Stdout, and Stderr are open Files pointing to the standard
// input, standard output, and standard error file descriptors.
var (
Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

os/file_unix.go
// NewFile returns a new File with the given file descriptor and name.
func NewFile(fd uintptr, name string) *File {


위에서 보듯이 세 변수를 선언할 때 Stdin, Stdout, Stderr를 목적지로 사용하고 있는데요. 파일을 포인터로 선언해서 쓰고 있으며 io.Writer 인터페이스를 구현하고 있습니다.

Error 로그를 살펴보겠습니다.



37 Error = log.New(io.MultiWriter(file, os.Stderr),
38 "ERROR: ",
39 log.Ldate|log.Ltime|log.Lshortfile)


위에서 보듯이 에러 로그는 io.MultiWriter라는 것을 사용하고 있습니다. io.MultiWriter 함수는 io.Writer 인터페이스값을 리턴하는데, 파일을 열면서(생성하면서) 동시에 출력도 할 수 있습니다. MultiWriter 함수는 io.Writer 인터페이스를 구현하는 여러 변수를 허용하는 가변 함수이며, 이 함수는 io.Writer 값이 전달 된 모든 값을 묶어 하나의 io.Writer값을 리턴해 줍니다. 그래서 log.New와 같은 기능으로 단일 작성자 내에서 여러 작성자를 허용 할 수 있습니다. 이제 사용자 정의 로거를 만드는 방법을 알고 있으므로 어떻게 사용할 수 있는지 살펴 보도록 하겠습니다.



8.21

42 func main() {
43 Trace.Println("I have something standard to say")
44 Info.Println("Special Information")
45 Warning.Println("There is something you need to know about")
46 Error.Println("Something has failed")
47 }
8.21에서 보면 만든 한 로그마다 하나의 메시지를 입력해서 출력하는 것을 볼 수 있습니다.



8.22

func (l *Logger) Fatal(v ...interface{})
func (l *Logger) Fatalf(format string, v ...interface{})
func (l *Logger) Fatalln(v ...interface{})
func (l *Logger) Flags() int
func (l *Logger) Output(calldepth int, s string) error
func (l *Logger) Panic(v ...interface{})
func (l *Logger) Panicf(format string, v ...interface{})
func (l *Logger) Panicln(v ...interface{})
func (l *Logger) Prefix() string
func (l *Logger) Print(v ...interface{})
func (l *Logger) Printf(format string, v ...interface{})
func (l *Logger) Println(v ...interface{})
func (l *Logger) SetFlags(flag int)
func (l *Logger) SetPrefix(prefix string)
Logger 타입을 사용하여 쓸 수 있는 메서드들입니다.

결론


로그 패키지는 오랜 역사를 갖고 있고 로그를 어떻게 적용할지에 기반해서 만들어 졌습니다. stdout과 stderr는 오랫동안 많은 CLI 기반 프로그램의 전통적 방식이었습니다. 프로그램이 로그를 출력 할 때 stdout, stderr 를 활용하는 것은 아주 괜찮은 방법 같습니다.

표준 라이브러리의 로그 패키지에는 로깅에 필요한 모든 것이 있습니다. 이왕이면 표준 라이브러리의 사용을 권장합니다. 구현을 신뢰할 수 있을 뿐만 아니라 표준 라이브러리가 커뮤니티에서 널리 사용되기 때문입니다.



참고 사이트



허접하기 그지없지만 나도 책을 따라서 하나 짜 봤다. 자동으로 로그가 저장되게 했는데 어떨지 모르겠다. 모듈로 정리해서 나중에 글에 덧붙이도록 하겠다...ㅠ
