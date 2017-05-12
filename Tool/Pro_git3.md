### Pro_git3

----

[참고자료](https://backlogtool.com/git-guide/kr/stepup/stepup1_4.html)

----

#### 병합시 생기는 충돌 해결하기

마스터 브랜치 체크아웃 후
issue 브랜치를 병합해 본다.

```
$ git checkout master
Switched to branch 'master'
$ git merge issue2
Updating b2b23c4..8f7aa27
Fast-forward
 myfile.txt |    2 ++
 1 files changed, 2 insertions(+), 0 deletions(-)
 ```
 
 아래와 같이 된다.
 
 issue3도 병합한다.
 
 ![병합](https://backlogtool.com/git-guide/kr/img/post/stepup/capture_stepup2_7_1.png)
 
 
 ```
 $ git merge issue3
Auto-merging myfile.txt
CONFLICT (content): Merge conflict in myfile.txt
Automatic merge failed; fix conflicts and then commit the result.
 ```
 
 이렇게 CONFLICT(충돌)이 났을 경우에는, 앞서 각각 브랜치에서 변경한 내용이 **같은 파일 내 같은 행**에 있기 때문이다.
 충돌부분은 아래와 같이 표시되어 찾기 쉽다.
 
 ```
 원숭이도 이해할 수 있는 Git 명령어
add: 변경 사항을 만들어서 인덱스에 등록해보기
<<<<<<< HEAD
commit: 인덱스의 상태를 기록하기
=======
pull: 원격 저장소의 내용을 가져오기
>>>>>>> issue3
 ```
 
 충돌부분 수정 후 다시 커밋하면 아래와 같이 된다.
 
 ![병합2](https://backlogtool.com/git-guide/kr/img/post/stepup/capture_stepup2_7_2.png)
 
 
 