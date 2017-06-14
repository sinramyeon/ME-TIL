#### linux tar 압축풀기

---

##### 1. tar 로 압축하기

`tar -cvf 파일명.tar 폴더명`
abc라는 폴더를 aaa.tar로 압축 `tar -cvf aaa.tar abc`

##### 2. tar 압축 풀기

`tar -xvf 파일명.tar`
aaa.tar파일 압축 풀기 `tar -xvf aaa.tar`

##### 3. tar.gz 로 압축하기

`tar -zxvf aaa.tar.gz`

##### 4. tar.gz 압축 풀기

`tar -zxvf aaa.tar.gz`


##### 옵션

옵션 | 설명
-- | --
-c | 파일을 .tar로 압축
-p | 파일 권한 저장
-v | 압축하거나 풀 때 과정을 화면으로 출력
-f | 파일이름 지정
-C | 경로 지정
-x | tar 압축 풀기
-z | gzip일 경우 붙이기