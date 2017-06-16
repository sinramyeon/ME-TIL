### Pro_git2

----

[참고자료1](https://elegantcoder.com/git-merge-or-rebase/)
[참고자료2](https://blog.outsider.ne.kr/666)
[참고자료3](https://backlogtool.com/git-guide/kr/stepup/stepup1_4.html)

----

##### 리베이스와 머지

- git-merge? : Join two or more development histories together.
- git-rebase? : Forward-port local commits to the updated upstream head.

* _히스토리 관리를 별로 신경쓰지 않고 혼자서만 커밋하거나 아니면 믿을수 있는 소수만 커밋하지 않는다면 merge대신 rebase를 사용하라고 권장하고 있습니다_

기능 | 장점 | 단점
-- | -- | ---
Merge | 이해하기 쉽다.<br>원래 브랜치 컨텍스트를 유지한다. <br>Fast-Foward Merge를 하지 않는다면 브랜치별로 커밋을 분리해 유지, 특히 이런 분리는 기능 브랜치에 유용하다. <br>원래 브랜치의 커밋들은 변경 없이 계속 유지되기 때문에, 다른 개발자들의 작업과 공유되는 것에 대해 신경쓸 필요가 없다. | 단순히 모든 사람들이 같은 브랜치에서 작업할 때는 merge 기록이 어지럽다.
Rebase | 단순한 히스토리<br>여러 개발자들이 같은 브랜치를 공유할 때 커밋을 합치는 가장 좋은 방법이다! | 충돌시 아주 복잡해진다.<br>커밋 순서대로 Rebase를 하는데, 각 커밋마다 충돌해소를 순서대로 해줘야 한다.<br>해당 커밋들을 다른곳에도 푸시했다면 히스토리를 다시 쓰는데 부작용이 발생한다. push는 가능하지만 나만 쓰는 리모트 브랜치에만 가능하다.

###### 결론

1. 여러 개발자들이 같은 브랜치를 공유할 때는 Pull & Rebase 가 히스토리를 깔끔하게 유지하는데 좋음.
2. 완료된 기능 브랜치를 다시 합칠 때는 머지를 사용.
3. 기능 브랜치에 부모 브랜치의 변경 내용을 반영하고 싶을 때는?
   1. 아래 상황에서는 리베이스가 낫다.<br>
      1. 이 브랜치를 다른 곳에 푸시한 적 없는 경우.<br>
      2. Git을 사용하고 다른 사람이 이 기능브랜치를 체크아웃할 일이 없을 것이라 확신하는 경우.<br>
   2. 이 외의 상황에는 머지가 낫다.
