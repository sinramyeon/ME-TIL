# .sh 파일

---

## sh파일 : script bash 파일
흔히 쉘 프로그래밍 파일 이라고 한다.

## 쉘 : 명령어를 시스템 커널에 전달

## sh파일 실행법? : sh filename.sh
* 윈도우일 경우 powershell이나 git bash 등에서 실행하면 된다.*

즉, sh파일을 통하여 bash에서 직접 필 명령어를 자동화할 수 있다고 생각하면 편할 것.

## 프로그래밍 언어? 스크립트 언어?

프로그래밍 언어 | 스크립팅 언어
----------------|----------------
빠르고 강력함   | 느리고 기능이 별로 없음
C, C++, Java 등 | Pearl, Lisp, Tcl 등
운영체제별로 이식이 어렵다 | 어느 운영체제에서나 사용할 수 있다
큰 규모의 어플리케이션에 사용 | 작은 프로그램에서 사용

---

쉘 스크립트 사용 중에는 *root계정을 쓰지 않도록 합시다.*
일반 사용자 계정을 만들어서 하는 것을 추천합니다.

예제_1)
hello.sh
```
#!/bin/bash # bash 번역기를 사용하세요
echo "Hello World"
```

`chmod 700 ./hello.sh`한 후 `./hello.sh`로 실행시 예제를 볼 수 있습니다.

예제_2) 모든 파일을 하나의 디렉토리로 옮기고, 그 디렉토리를 내용물 모두와 함께 지운 뒤, 다시 그 디렉토리를 만드는 프로그램
mkdir.sh
```
#!/bin/bash
mkdir trash
mv * trash
rm -rf trash
mkdir trash
echo "Deleted!"
```

[참고](http://coffeenix.net/doc/shell/introbashscript.htm)