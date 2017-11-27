# 정말 쉬운 파이썬 os

---

go를 쓰다보면 정말 그리운 몇 가지 파이썬 기능들이 있다.
그 중 하나는 심플하기 그지없는 시스템 관련 기능일 것이다.
`import os`로 os모듈에서 쓸 수 있는 간편한 기능을 모았다.

## open() 생성

`fout = open("test.txt", "wt")`

있으면 열고 없으면 만든다

## exists() 존재여부 확인

`os.path.exists("./no.txt")`

존재여부를 반환한다.

## isfile() 타입 확인

세트로 물론 isdir()도 있다.

## isabs() 절대경로 확인

`os.path.isabs("/big/fake/name")`

인자가 절대 경로인지 반환한다.

## copy() 복사하기

`import shutil`

os모듈 외 shutil에 들어있다.

`shutil.copy("oops.txt", "ohno.txt")`

## rename() 이름 바꾸기

`os.rename("oops.txt", "ohno.txt")`

## chmod() 퍼미션 바꾸기

`os.chmod("oops.txt", stat.S_IRUSR)`

chomod에서 바로 777등으로 못 쓰고 0o400인지 하는 8진수 값을 쓰는게 조금 짜증난다.

`import stat`

stat모듈을 가져와서 값을 가져다 쓰ㄴ느것이 좋다.

## chown() 오너십 바꾸기

`os.chown("oops", uid, gid)`

## abspath() 절대경로 얻기

`os.path.abspath("text.txt")`

굉장히 편한 함수 중 하난데, `/usr/hero/text.txt` 식으로 절대 경로를 반환해 준다.

## remove() 삭제하기

`os.remove("text.txt")`

## mkdir() / rmdir() 폴더 만들고 지우기

`os.mkdir("폴더")`와 `os.rmdir("폴더")`로 간단한 폴더 생성 및 삭제가 가능하다.

## listdir() 하위 콘텐츠 나열

`os.listdir("폴더")`

