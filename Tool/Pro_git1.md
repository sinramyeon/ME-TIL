### Pro_git

----

[참고자료1](https://tuwlab.com/ece/22218)
[참고자료2](http://mobicon.tistory.com/category/Git%2C%20GitHub)

----

##### 특정 커밋을 선택해서 반영하기 - cherry-pick

다른 branch에 있는 commit을 선별적으로 현재 branch에 반영할 수 있다.

`git cheery-pick {Commit ID}`

이때, 현재 브랜치는 작업 내역을 반영할 브랜치여야 하고, 가져오려는 커밋이 여러개인 경우 먼저 작성한 커밋부터 순서대로 가져온다.

현재 브랜치에 붙는 커밋은 커밋 아이디가 달라진다.

Cherry-pick은 새 commit 을 하는 것과 같다.

##### 여러개의 커밋을 반영하기 - rebase

대상 브랜치의 변경사항을 모두 가져와서 현재 브랜치에 반영할 수 있다.

`git rebase {가져올 브랜치 이름}`

결과적으로 대상 branch는 변경 없이 현 브랜치에만 새 커밋이 생성되어 덧붙여진다.

##### 커밋 순서 재정령, 첨삭

rebase 명령 뒤 -i옵션을 덧붙이면 commit의 순서를 바꾸고 첨삭하거나 몇 개의 커밋을 하나로 합치는 등의 작업 수행이 가능하다.

`git rebase -i HEAD~3`
현 branch head로부터 3개의 커밋을 편집 가능하다.

