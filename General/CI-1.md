# CI란?

## CI

> Continuous Integration

프로그래밍이 끝나면 항상 찾아오는 코드리뷰 아니... 머지, 빌드, 디플로이 시간이 있다. 항상 SCM(Source Code Management) 툴을 이용해서 다같이 하는 작업이라 익숙하실 것이다. SCM이 뭐냐고? Git이다(SVN 을 사용하시는 분들 죄송합니다)

그렇지만 점점 다양화, 다각화되는 개발 사업에서 매주 금요일 밤 5시 반에 디플로이를 하는 미친 짓을 계속 할 수는 없는 것이다. **CI란,  짧은 주기로 소프트웨어 빌드, 테스트, 컴파일 과정을 자동화하는 것이다.**




![](https://cdn-images-1.medium.com/max/1600/1*ezikvlWu0UAwFfh3ko8LXQ.png)

깃랩에서 설명한 CI/CD의 개념 이미지. 이해하는 데에 도움이 될 수 있다.

## CD

>Continous Deploy
**CI란,  짧은 주기로 소프트웨어를 끊어 배포하고 그 과정 전체를 자동화하는 것이다.**

**즉 컴파일 - 빌드 - 테스트 - 배포 과정 자체의 자동화** 및 **자주하기** 라고 이해하면 쉽다.

### CI/CD가 왜 필요한지에 대하여

원래 회사에서는 매주 나하고싶을때 릴리즈를 배포하는 아주 안 좋은 버릇이 있었다. 그러다 보니 매일 급하게 bugfix 로 땜빵한 코드 체인지가 늘어가고, 이놈이 커밋을 머지하기 전에 제대로 테스트를 돌리고 한건지? 왜 머지했는데 빌드에러가 뜨는지? 등의 온갖 고난을 겪을 수 밖에 없었다.

그렇다면 만약

> 버전관리 시스템 -> 컴파일 -> 빌드 -> 테스트 -> 머지 -> 배포

를 이놈들이 커밋할 때마다(나 포함) 자동화되게 한다면 어떨까? 엄청난 편의성과 오류 감소를 기대할 수 있지 않을까?

1. 커밋마다 컴파일 빌드 테스트 자동화
2. 코드 검증 결과를 빠른 전달 가능
3. 장애 및 오류 감소!
4. Profit!

###  CI/CD와 관련된 툴들 소개

Jenkins, Travis, JUnit, Mocha, Maven 등을 어디선가 들어본 적이 있으실수도 있다. 현재 사내에서는 Gitlab CI(꽁짜고 뭔가 설치가 쉬웠다.) 와 내가 얼기설기 짠 테스트 툴(`go test -v --cover` 를 때려박아 사용중)을 사용하고 있지만, 대략 다음과 같은 툴을 한번 검색해보시면 감이 쉽게 오실 수 있다.

#### CI 서버
Jenkins, Travis CI, Gitlab CI

#### 빌드 툴
Maven, Gradle, Ant

#### 테스트 툴
Junit, Mocha

---

# GITLAB CI 설치하기

사실 설치에 대해서는 [공식 문서]([https://docs.gitlab.com/runner/install/](https://docs.gitlab.com/runner/install/)) 나 많은 선배님들의 블로그 [올라운드플레이어 님의 블로그]([https://allroundplaying.tistory.com/21](https://allroundplaying.tistory.com/21)) 에 너무 설명이 잘 되어 있어서 그닥 설명할 게 없는 느낌이 든다. 특히 **올라운드플레이어** 님 블로그만 봐도 눈을 감고도 설치할 수 있는 단순함마저 보인다.

### Gitlab CI runner 란?

GITLAB CI 에는 Gitlab CI Runner 란 것이 있는데, 이 러너란 놈을 깔면 우리가 하라고 한 걸 하고 그 결과를 깃랩 서버에 전달해 준다. 즉 이 Gitlab CI Runner 를 CI 서버로 쓸 곳에 계속 실행시켜 놓고 CI/CD를 이어나간다고 생각하면 된다.

# GITLAB CI 설정하기

사실 gitlab runner 설치 자체는 단순하다. gitlab 내에서 CI 관련 설정을 하는 것이 더욱 복잡하게 느껴질 수 있다.

