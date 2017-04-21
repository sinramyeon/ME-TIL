#### CI를 alaboza.

---

[참고자료](http://www.nextree.co.kr/p10799/)
[참고자료2](http://happystory.tistory.com/89)

---


##### Continuous Integration

소프트웨어 개발에서 유지보수로 연결되는 지점은 소스관리, 빌드, 배포이다.
이 활동을 지원하는 기법 중 하나 CI!

자동화된 일일 빌드를 지원하는 툴이라고 이해하면 되는데, 기본 원칙은 다음과 같다.

* 모든 소스 코드는 현재 실행되고, 어느 누구든 현재의 소스를 접근할 수 있는 단일 지점을 유지할 것
* 빌드 프로세스를 자동화시켜서 어느 누구든 소스로부터 시스템을 빌드하는 단일 명령어를 사용할 수 있게 할 것
* 테스팅을 자동화 시켜서, 단일 명령어를 통해 언제든지 시스템에 대한 건전한 테스트 수트를 실행할 수 있게 할 것
* 누구나 현재 실행파일을 얻으면 그게 제일 나은 실행파일이라는 확신을 얻게 만들 것


CI 서버는 형상관리 서버에 Commit된 소스를 주기적으로 폴링하여 컴파일, 단위테스트,**코드 인스펙션** 등의 과정을 수행하여 코드 결함 여부를 지속적으로 검증한다.

##### CI시스템 구축

1. CI server : 빌드 프로세스 관리 서버. Jenkins, Hudson, CruiseControl.Net, TeamCity, GitLab 등

2. SCM : 소스코드 형상관리 시스템. SVN, Git, Mercurial

3. Build Tool : 컴파일, 테스트, 정적 분석 등을 실시해 동작 가능한 소프트웨어를 생성. Ant, Maven, MSBuild, Make

4. Test Tool : 테스트 코드를 따라 자동으로 테스트를 수행해주는 도구. JUnit, CppUnit, CppTeset, MSTest, Selenium

5. Test Coverage Tool : 테스트 코드가 대상 소스 코드에 대해 어느정도 커버하는지 분석하는 도구. Emma, Cobertura, TestCoccon

6. Inspection Tool : 프로그램을 실행하지 않고 소스코드 자체의 품질을 판단. CheckStyle, FindBugs, Cppcheck, Valgrind


> Ex
> 1. 소스코드를 바이너리 파일로 컴파일
> 2. 바이너리 파일을 배포 형태로 패키징
> 3. 단위 테스트(선택사항)
> 4. 정적 분석 (선택사항)
> 5. 분석 결과 리포트(선택사항)
> 6. 패키징한 파일 서버 배포

---

*코드 인스펙션* :  개발팀에서 작성한 개발소스 코드를 분석하여 개발 표준에 위배되었거나 잘못 작성된 부분을 수정하는 작업
