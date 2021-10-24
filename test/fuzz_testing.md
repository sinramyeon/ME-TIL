## Fuzz Testing  

**Fuzz Testing**  or Fuzzing is a software testing technique of putting invalid or random data called FUZZ into software system to  discover security vulnerabilities or bugs in software applications. 

Fuzz testing is a novel way toUnlike traditional software testing methodologies, fuzz testing essentially “pings” code with random (or semi-random) inputs in an effort to crash it and thus identify “faults” that would otherwise not be apparent.

The purpose of fuzz testing is inserting data using automated or semi-automated techniques and testing the system for various exceptions like system crashing or failure of built-in code, etc.

![Fuzz Testing](https://cdn.guru99.com/images/3-2016/032816_0730_FuzzTesting1.png)

## Why to do Fuzz Testing?

-   Usually, Fuzzy testing finds the most serious security fault or defect.
-   Fuzz testing gives more effective result when used with  [Black Box Testing](https://www.guru99.com/black-box-testing.html), Beta Testing, and other debugging methods.
-   Fuzz testing is used to check the Vulnerability of software. It is very cost effective testing techniques.
-   Fuzz testing is one of the black box testing technique. Fuzzing is one of the most common method hackers used to find vulnerability of the system.

## How to do Fuzz Testing

The steps for fuzzy testing include the basic testing steps-

**Step 1)**  Identify the target system

**Step 2)**  Identify inputs

**Step 3)**  Generate Fuzzed data

**Step 4)**  Execute the test using fuzzy data

**Step 5)**  Monitor system behavior

**Step 6)**  Log defects

## FUZZAMPLE

[google/gofuzz](https://github.com/google/gofuzz) 를 사용

Readme에 보면 아예 이렇게 적혀 있음

 ```
This is useful for testing:

-   Do your project's objects really serialize/unserialize correctly in all cases?
-   Is there an incorrectly formatted object that will cause your project to panic?
```

#### 1. single variable에 적용하기

```
f := fuzz.New()
var myInt int
f.Fuzz(&myInt) // myInt gets a random value.
```

#### 2. map에 적용하기

```
f := fuzz.New().NilChance(0).NumElements(1, 1)
var myMap map[ComplexKeyType]string
f.Fuzz(&myMap) // myMap will have exactly one element.
```

#### 3. nil pointer 에 적용하기
```
f := fuzz.New().NilChance(.5)
var fancyStruct struct {
  A, B, C, D *string
}
f.Fuzz(&fancyStruct) // About half the pointers should be set.
```
## Two types of fuzz testing


#### Coverage-guided fuzz testing
focuses on the source code while the app is running, probing it with random challenges in an effort to uncover bugs. New tests are constantly being generated and the goal is to get the app to crash. A crash means a potential problem, and data from the coverage-guided fuzz testing process will allow a tester to reproduce the crash which is helpful when trying to identify at-risk code.

#### Behavioral fuzz testing 
Using specs to show how an application should work it uses random inputs to judge how the app really works; the difference between the expected and the reality is generally where bugs or other potential security risks can be found.
