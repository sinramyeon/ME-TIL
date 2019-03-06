### 오버로딩과 오버라이딩 그 차이는?

---

[참고자료](http://private.tistory.com/25)
[참고자료2](http://hyeonstorage.tistory.com/185)

---

#### 오버로딩
##### 같은 이름의 메서드 여러개를 가지면서 매개변수의 유형과 개수가 다르도록 하는 기술

같은 이름 메서드를 여러개 정의, 매개변수의 유형과 개수를 다르게 -> 다양한 유형의 호출에 응답


```
class OverloadingTest{
    //이름이 cat인 메서드
    void cat(){
        System.out.println("매개변수 없음");
    }
    
    //매개변수 int형이 2개인 cat 메서드
    void cat(int a, int b){
        System.out.println("매개변수 :"+a+", "+b);
    }
    
    //매개변수 String형이 한 개인 cat 메서드
    void cat(String c){
        System.out.println("매개변수 : "+ c);
    }
    
}
 
public class OverTest {
 
    public static void main(String[] args) {
        
        //OverloadingTest 객체 생성
        OverloadingTest ot = new OverloadingTest();
        
        //매개변수가 없는 cat() 호출
        ot.cat();
        
        //매개변수가 int형 두개인 cat() 호출
        ot.cat(20, 80);
        
        //매개변수가 String 한개인 cat() 호출
        ot.cat("오버로딩 예제입니다.");
        
    }
 
}

```

> 실행 결과<Br>
> 매개변수 없음<Br>
> 매개변수 : 20. 80<Br>
> 매개변수 : 오버로딩 예제입니다.

---

오버로딩인 단순 변수차이, 그럼 오버라이딩은?

#### 오버라이딩
##### 상속 관계에 있는 클래스 간에 같은 이름 메서드를 하위 클래스에서 재정의

상위 클래스 내 멤버 변수와 메서드를 하위 클래스에서 상속해서 사용할 수 있습니다. 하위 클래스에서 메서드를 재정의해서도 사용할 수 있음.

```
class Woman{ //부모클래스
    public String name;
    public int age;
    
    //info 메서드
    public void info(){
        System.out.println("여자의 이름은 "+name+", 나이는 "+age+"살입니다.");
    }
    
}
 
class Job extends Woman{ //Woman클래스(부모클래스)를 상속받음 : 
 
    String job;
    
    public void info() {//부모(Woman)클래스에 있는 info()메서드를 재정의
        super.info();
        System.out.println("여자의 직업은 "+job+"입니다.");
    }
}
 
public class OverTest {
 
    public static void main(String[] args) {
        
        //Job 객체 생성
        Job job = new Job();
        
        //변수 설정
        job.name = "유리";
        job.age = 30;
        job.job = "프로그래머";
        
        //호출
        job.info();
        
    }
 
}

```

> 실행결과<Br>
> 여자의 직업은 프로그래머입니다.<Br>
> 여자의 이름은 유리, 나이는 30살입니다.

super 예약어로 부모 클래스에 있는 내용을 쓸 수 있습니다.

---

구분 | 오버로딩 | 오버라이딩
------------ | ------------- | -------------
메서드 이름 | 동일 | 동일
매개변수,타입 | **다름** | 동일
리턴 타입 | **상관없음** | 동일
