#### git 사용법

----

[참고자료](https://github.com/milooy/TIL/blob/master/git/git-guide.md)

---

##### git 기본 사용법

git init

` # 레포지토리 초기화. '.git'폴더 생성`

git clone /로컬저장소/경로

`로컬 저장소 경로를 복제`

git clone 사용자@호스트:/원격저장소/경로

`원격저장소를 복제`

git status

` # 레포지토리 상태 표시`

git add 파일명

`스테이지 영역(커밋 전 임시 영역)에 파일 추가`

git add *
git add --all

`파일 전부 추가`

git commit -m "커밋"

`커밋 - 스테이지 영역에 기록된 시점들 파일을 실제 레포지토리 변경내역에 반영`

git log

`커밋 로그 확인`

git log README.md

`README.md 수정내역 확인`

git log -p

`커밋에서 변경된 내용도 함께 확인`

git diff

`working tree, stage, commit 사이의 변경 확인`

git diff HEAD

`최신 커밋과의 차이 확인`

git branch

`브랜치 목록 표시`

git checkout -b 'NEWBRANCH'

`NEWBRANCH 라는 브랜치 만들고 거기로 이동`

git checkout master

`master 브랜치로 돌아옴`

git branch -d 'NEWBRANCH'

`NEWBRANCH 브랜치 삭제`

git push origin 'NEWBRANCH'

`새 브랜치 원격 저장소로 전송`

git tag 1.0.0 번호

`소프트웨어 새 버전마다 꼬리표 달기`

git merge --no--ff "NEWBRANCH"

`현 브랜치에서 NEWBRANCH 브랜치를 머지 커밋도 함께 남기면서 머지`
[참고자료](https://git-scm.com/book/ko/v1/Git-%EB%B8%8C%EB%9E%9C%EC%B9%98-%EB%B8%8C%EB%9E%9C%EC%B9%98%EC%99%80-Merge%EC%9D%98-%EA%B8%B0%EC%B4%88)


git log --graph

`브랜치 그래프 확인`

git reset --hard 번호

` HEAD Stage Working tree 모두 특정 커밋 상태로 복원`

git fetch origin
git reset --hard origin/master

`로컬에 있는 모든 변경 내용과 확정본 포기`

git checkout -- '파일이름'

`로컬 변경 내용을 변경 전 상태로 복원`

git commit --amend

`방금전 한 커밋 메시지 수정`

git push

`커밋을 반영합니다.`
