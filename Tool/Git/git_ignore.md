# Git 잘못 푸시한 파일 삭제하기

## \bin, \pkg 라거나 잘못된 폴더를 아무 생각 없이 add --all 해서 push 한 경우

`git rm - r --cached "폴더명"(혹은 파일명)`
`git commit -m "지웠음"`
`git push origin master`

---

# .Gitignore 파일 만들기

## 아무생각 없이 .gitignore를 .txt 등으로 만들지 말자...

```
# .gitignore 파일 생김새

*.exe # .exe확장자 빼고
!must.exe # 그렇지만 must.exe파일은 꼭 포함
GIT/ # GIT 폴더 전체 무시
```