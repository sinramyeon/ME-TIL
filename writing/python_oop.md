객체지향, OOP, 자바를 잡아 할 때부터 열심히 공부...는 하려고 했지만 사실 안했던 부분이다. 당최 이 객체라는게 무어고 이걸 어떻게 적용시킨단 말인가. 파이썬을 쓰면서 완전히 막장 지향형 코드를 짜고 있는 듯 하여 파이썬에서의 oop를 짚고 가기 위해 부족하지만 정리해 봅니다.



class & objects
class와 object 이해

클래스 - 객체 생성을 위해 변수와 메소드를 정의하는 일종의 틀. A class is logical grouping of attributes(variables) and methods(functions)



객체 - 클래스의 인스턴스. Object is an instance of a class that has access to all the attributes and methods of that class

라고 정의되어 있는데, 파이썬에서는 어떻게 표현되냐면 아래와 같다.



클래스

class ClassName :
    class body
인스턴스

objectName = ClassName() #카멜 케이스랑 스네이크 케이스를 지켜 주자.


파이썬에선 모든것이 객체이다.



numbers = [1,2,3]
type(numbers)
<class 'list'>


예제

class Employee:
    # 클래스 내 변수 셋업
    name = "Ben"

    def changeName (self):
        # 메서드 안에서 이름을 변경했음
        # self 로 메서드 안 변수에 접근할 수 있다.
        Employee.name = "Mark"

employee = Employee()
print(employee.name) # Ben
employee.changeName()
print(employee.name) # Mark

attributes & methods
클래스 어트리뷰트? 인스턴스 어트리뷰트? 애트리뷰트?


한글로 쓰는게 더 어렵다. 우선 instance attribute는 클래스의 인스턴스가 갖고 있는 애들이다. 여러 인스턴스가 있다면 인스턴스 별로 각자 다른 애트리뷰트(어트리뷰트???) 를 가지고 있다.

attribute를 클래스 단위에서 선언할 수도 있는데, 이렇게 클래스 자체가 갖고 있는 attribute를 class attribute라고 부른다. 클래스의 모든 인스턴스에서 공유하고, 값도 같다. 보통은 클래스 선언 바로 아래 써 준다.



클래스 어트리뷰트 예제

An attribute that is common across all instances of a class is called a class attribute



class Employee :
    numberOfWorkingHours = 40

employeeOne = Employee()
employeeTwo = Employee()

employeeOne.numberOfWorkingHours

>>> 40 # 둘다 40


class attribute값을 변경하고 싶다면 클래스명.변수명 으로 접근하여 변경할 수 있다.



Employee.numberOfWorkingHours = 45
employeeTwo.numberOfWorkingHours

>>> 45 # 둘다 45


인스턴스 어트리뷰트 예제



An attribute that is specific to each instance of a class is called an instance attribute.

employOne.name = "John"
employTwo.name = "NotJohn"

employOne.name
>>> John
employTwo.name
>>> NotJohn


주의

employeeOne.numberOfWorkingHours = 1
employeeOne.numberOfWorkingHours

>>> 1


employeeTwo.numberOfWorkingHours

>>> 45 # 1이 아니다.


여기서는 클래스 어트리뷰트와 같은 변수 이름을 사용해서 클래스 어트리뷰트와 같다고 착각할 수 있으나 이것은 인스턴스 어트리뷰트를 만드는 문법이다.

클래스 어트리뷰트를 다루고 싶으면 클래스이름.어트리뷰트 이름 으로 호출하는것을 잊지 않는다.



self?


self를 자바의 this 비슷하게 생각하며 헷갈릴 수 있다고 생각한다. 클래스 내의 함수를 메서드라고 부르고 보통 메서드의 인자로 self를 넣게 된다. 



self는 객체의 인스턴스 그 자체라고 이해하면 쉽다.

class Employee :
    def employeeDetails() :
        pass

employee = Employee()
employee.employeeDetails()


위와 같이 self인자를 미사용하고도 ide에선 오류가 나지 않는다. 그러나 이 메서드를 시행하면 에러가 발생한다.



Employee.employeeDetails(employee)


에러가 나는 이유는 파이썬에서 함수를 콜할때 위와 같이 부르기 때문이다. 객체 employee가 self에 들어간다. 여기서 employee가 들어갈 부분이 비어 있기 때문에 에러가 나는 것이다.



def employeeDetails(self) :


로 변경하면 에러가 나지 않는다. 물론 self대신 아무 변수로나 써도 되지만 파이썬 컨벤션이 이러므로 self를 넣어주자.



class Employee :
    def employeeDetails(self) :
        self.name = "Matthew"
        print(self.name)
		age = 30
		print(age)
		
employee = Employee()
employee.employeeDetails()
>> Matthew
>> 30


self.로 부르지 않아도 잘만 사용할 수 있는 것 같은데 왜 self.name식으로 써야만 할까?



class Employee :
    def employeeDetails(self) :
        self.name = "Matthew"
        print(self.name)
		age = 30
		print(age)
	
	def printEmployeeDetails(self) :
	    print(self.name)
	    print(age)
		
employee = Employee()
employee.printEmployeeDetails()

>>> age is not defined 에러


self 없이는 age변수의 라이프사이클은 employeeDetails()안 뿐이다.



파이썬에서는 형이나 범위 강제 문법이 없고 @를 쓰지 않으니 self를 명시적으로 표현하는게 중요하다.


static method, instance method


스태틱, 인스턴스 메서드의 차이를 알아보겠다. 대략 개념을 알고는 있는데 차이를 활용해서 잘 못 쓰고 있지는 않은지 싶다.



class Employee :
    def employeeDetails(self) :
        self.name = "Ben"

    def welcome(self) :
        print("Welcome")

employee = Employee()
employee.employeeDetails()
print(employee.name)

employee.welcome()

>>> Ben
>>> Welcome


보통 여기서 welcome()에 쓸데없는 self를 없애고 싶으면 static method를 사용한다. static method란 self파라미터를 기본으로 받지 않는 메서드를 뜻한다.



class Employee :
    def employeeDetails(self) :
        self.name = "Ben"

    @staticmethod
    def welcome() :
        print("Welcome")

employee = Employee()
employee.employeeDetails()
print(employee.name)

employee.welcome()

>>> Ben
>>> Welcome


@staticmethod 데코레이터를 붙여주면, self 바인딩을 자동적으로 없애줘서 위와 같은 사용이 가능하다. static method에는 class에 대한 아무런 정보가 담겨있지 않게 된다.



# s.py
class S:
   nInstances = 0
   def __init__(self):
      S.nInstances = S.nInstances + 1

   @staticmethod
   def howManyInstances():
      print('Number of instances created: ', S.nInstances)

>>> from s import S
>>> 
>>> a = S()
>>> b = S()
>>> c = S()

>>> S.howManyInstances()
('Number of instances created: ', 3)
>>> a.howManyInstances()
('Number of instances created: ', 3)


보다시피 static method는 인스턴스로도, 클래스로도 호출할 수 있다.



init()


객체를 초기화하는 __init__에 대해 알아보자.



class Employee :
   def enterEmployeeDetails(self) :
       self.name = "Mark"

    def displayEmployeeDetails(self) :
        print(self.name)


employee = Employee()
employee.displayEmployeeDetails() # name을 만들어주기 전 사용하려고 하면 no attribute 에러가 발생한다.


초기화시켜줘야할 값이 있다면 메서드로 사용하지 말고 __init__ 을 사용하자. __init__은 메서드 호출 전 이니셜라이즈를 먼저 시켜준다.



class Employee :

    def __init__(self) :
        self.name = "Mark"
        
    def displayEmployeeDetails(self) :
        print(self.name)


이렇게 하면, 따로 호출하지 않아도 알아서 name에 값이 할당된다.



class Employee :

    def __init__(self, name) :
        self.name = name
        # self.name > 인스턴스 애트리뷰트인 name
        # name > 파라미터로 넣은 name
        
    def displayEmployeeDetails(self) :
        print(self.name)

employee = Employee("Mark")
employeeTweo = Employee("Matthew")

employee.displayEmpoyeeDetails()
>>> Mark
employeeTwo.displayEmployeeDetails()
>>> Matthew


예제

class PreciousStone:
    numberOfPreciousStones = 0
    preciousStoneCollection = []
    def __init__(self, name):
        self.name = name
        # 클래스 변수를 증가시킴
        PreciousStone.numberOfPreciousStones += 1
      
        if PreciousStone.numberOfPreciousStones <= 5:
            PreciousStone.preciousStoneCollection.append(self)
        else:
            del PreciousStone.preciousStoneCollection[0]
            PreciousStone.preciousStoneCollection.append(self)

    @staticmethod # 스태틱 메서드
    def displayPreciousStones():
        for preciousStone in PreciousStone.preciousStoneCollection:
            print(preciousStone.name, end = ' ')
        print()

preciousStoneOne  = PreciousStone("Ruby")
preciousStoneTwo  = PreciousStone("Emerald")
preciousStoneThree  = PreciousStone("Sapphire")
preciousStoneFour  = PreciousStone("Diamond")
preciousStoneFive  = PreciousStone("Amber")
preciousStoneFive.displayPreciousStones()
preciousStoneSix = PreciousStone("Onyx")
# 첫번째 보석을 없애고 다시 출력
preciousStoneSix.displayPreciousStones()

Abstraction & Encapsulation


추상화와 캡슐화는 떨어질 수 없는 개념이다.



기본적 개념은 아래와 같다.



Abstracting something means to give names to things, so that the name captures the core of what a function or a whole program does.

Hiding the implementation details from the end user is called as encapsulation



추상화는 공통의 속성이나 기능을 묶어 이름을 붙이는 것이고 캡슐화는 함수와 변수를 하나로 묶는 것이다. 데이터를 함수를 통해서만 접근할 수 있도록 하는 것인데, 캡슐화에 성공하면 자연스럽게 은닉화도 할 수 있게 되는 것.



캡슐화의 대표적 예제를 확인하자.

#!/usr/bin/env python
 
class Car:
 
    def __init__(self):
        self.__updateSoftware()
 
    def drive(self):
        print 'driving'
 
    def __updateSoftware(self):
        print 'updating software'
 
redcar = Car()
redcar.drive()
#redcar.__updateSoftware()  not accesible from object.


객체를 바로 불러낼 수 없고 클래스 안에서만 사용할 수 있는것을 확인할 수 있다. 이럴 때 파이썬에서는 redcar._Car__updateSoftware() 식으로 호출한다.

클래스 인스턴스.__클래스명__캡슐화된 메서드 로 호출하며, 메서드 작성 시에도 알아보기 쉽게 캡슐화할 메서드(* 파이썬엔 private가 없다는 걸 기억하자. 고에서 소문자 대문자 쓰는것처럼 파이썬이는 무조건 __다.) 에는 __를 붙이자.



class Library :

	def __init__(self, listOfbooks) :
	    self.availableBooks = listofBooks
    # 리스트 할당
   
    def displayAvailableBooks(self) :
        print()
        for book in self.availableBooks :
            print(book)

    def lendBooks(self, requestedbBook) :
        if requestedbBook in self.availableBooks :
            self.availableBooks.remove(requestedBook)
        else :
            print()
    def addBooks(returnedBook) :
        self.availableBooks.append(returnedBook)
       


class Customer :
    def requestBook(self) :
        print()
        self.book = input()
        return self.book
        
    def returnBook(self) :
        self.book = input()
        return self.book


java에서는 private를 썼고 go에서는 소문자를 썼다면 python에서는 언더바(__) 이다. java에서는 protected를 썼다면 python에서는 언더바(_)이다.

__init__을 생각해보면 어렵지 않다. 상속에서 다시 한번 다루겠다.

여기서 그럼 getter, setter는 어디에 있냐는 생각을 할 수 있다. getter와 setter는 많은 객체지향 언어에서 캡슐화를 위해 사용하고 있다. ide에서 자동으로 만들어 주기까지 한다. 그렇지만 파이썬에서도 똑같이 하지 말고 꼭 전에 학습한 @property를 떠올린다.



@property

class Test:

    def __init__(self):
	    self.color = "red"

	def set_color(self,clr):
		self.color = clr

	def get_color(self):
		return self.color

if __name__ == '__main__':

	t = Test()
	t.set_color("blue")

	print(t.get_color())


위와같이 get, set 메서드를 직접 만들 수도 있지만 @property를 사용해도 된다.



class Test:

    def __init__(self):
        self.__color = "red"

    @property
    def color(self):
        return self.__color

    @color.setter
    def color(self,clr):
        self.__color = clr

if __name__ == '__main__':

    t = Test()
    t.color = "blue"

    print(t.color)


@property 로 red값을 가져온 후 t.color(self, clr) 세터로 밖에서 호출 시 set해줄 수 있다.

@property 를 사용하는 목적을 간단하게 말하면

변수를 변경 할 때 어떠한 제한을 두고 싶어서
get,set 함수를 만들지 않고 더 간단하게 접근하게 하기 위해서
하위호환성에 도움이 됨.
(출처: http://hamait.tistory.com/827 [HAMA 블로그])



또 이걸 쓰는 편이 훨씬 파이썬을 잘해보인다(실제로 다른 곳에서도 get, set은 ugly하다는 평이 많음...) *참고 "There should be one-- and preferably only one --obvious way to do it."



아래와 같은 예제를 참고하자.

@property
def color(self) :
    return self.color

#a.color 식으로 접근

@color.setter
def color(self, val) :
     self.color = val
#a.color = "red" 로 변경하면 setter가 작동되어 객체지향형 코딩이 가능!
python naming convention


java등을 쓰다 오면 public, private, protected를 못 써서 불안할수도 있다. 다시 한번 정리하자면 아래와 같다.

# public -> 그냥쓰기
# protected -> 클래스로만 접근할수 있게 하려면 _멤버네임
# private -> __멤버네임


파이썬에서는 사실 그냥 다 public을 쓰는걸 권장하는 것 같기도 하다.

class Car :
    numberOfWheels = 4
    _color = "Black"
    __yearOfManufacture = 2018

car = Car()
print(car.numberOfWheels)

class CarWithProctected(Car) :
    def __init__(self) :
        print(self._color)

bmw = CarWithProtected() # _color가 그렇지만 밖에서도 접근 안되는건 아니다.
car.__yearOfManufacture # __yearOfManufacture 를 접근하면 여기서는 에러가 발생한다.


__ 로 접근시 네임 맹글링이 발생해서 에러가 발생하게 된다. 꼭 숨겨야할게 있으면 __를 아용한다.

Name mangling is done by prepending the member name with an underscore and class name.



접근 방식은 _className__memberName와 같이 이루어진다.



Inheritance
inheritance : inheritance is when an object or class is based on another object (prototypal inheritance) or class (class-based inheritance), using the same implementation.



상속은 세가지 종류로 나눌 수 있다.



single inheritance
class Apple :
    manufacturer = "Apple Inc."
    contactWebsite = "www.apple.com/contact"

    def contactDetails(self) :
        print(self.contactWebsite)

class MacBook(Apple) :
    # Apple 클래스를 상속받는 맥북 클래스
    def __init__(self) :
        self.yearOfManufacture = 2018
    def manufactureDetails(self) :
        print(self.yearOfManufacture, self.manufacturer)
# 부모 함수 내 변수를 self.로 불러내고 있음

macBook = MacBook()
macBook.manufactureDetails()

>>> 2018 Apple Inc.

multiple inheritance

class OperatingSystem :
    multitasking = True
    name = "Mac Os"

class Apple :
    website = "www.apple.com"
    name = "Apple"

class MacBook(Apple, OperatingSystem) :
    # 다중 상속 가능
    def __init__(self) :
        if self.multitasking is True :
            print("This is a multi tasking system. Visit {} for more details.", self.website)
            # True이므로 www.apple.com이 출력된다.

            print(self.name) # Apple이 출력된다.
            # 같은 변수가 중복될 경우 먼저 상속한 클래스 것 부터 따르게 된다.

multilevel inheritance

class MusicalInstruments :
    numberOfMajorKey = 12


class StringInstruments(MusicalInstruments) :
    typeOfWood = "Tonewood"

class Guitar(StringInstruments) :
    # 3단으로 상속받는중
    def __init__(self) :
        self.numberOfString = 6
        print(self.numberOfString, self.typeOfWood, sel.fnumberOfMajorKeys)

guitar = Guitar() # 접근 가능

Polymorphism
In programming languages and type theory, polymorphism is the provision of a single interface to entities of different types. A polymorphic type is one whose operations can also be applied to values of some other type, or types.[2] There are several fundamentally different kinds of polymorphism



class Employee :
    def setNumberOfWorkingHours(self) :
        self.numberOfWorkingHours = 40
    def displayNumberOfWorkingHours(self) :
        print(self.numberOfWorkingHours)
employee = Employee()
employee.setNumberOfWorkingHours()

employee.displayNumberOfWorkingHours() # 40

class Trainee(Employee) :
    # 상속 후
    # 부모 메서드를 오버라이딩 가능
    def setNumberOfWorkingHours(self) :
        self.numberOfWorkingHours = 45
        # 변경
    def resetNumberOfWorkingHours(self) :
        super().setNumberOfWorkingHours() # super -> 부모클래스 호출

trainee = Trainee()
trainee.setNumberOfWorkingHours()
trainee.displayNumberOfWorkingHours() # 45

trainee.resetNumberOfWorkingHours()

trainee.displayNumberOfWorkingHours() # 40


오버라이딩을 할 때 다중상속에서 diamond shape problem이라는것을 발생시킬 수 있다.







대략 이런 상황을 일컬어 diamond shape problem이라고 한다.



class A :
    def method(self) :
        print("A")

class B(A) :
    pass

class C(A) :
    pass

class D(B, C) :
    pass

d = D() # A, B, C를 모두 부모를 가짐.

# 만약

# 1. b와 c에서 a에 있는 메서드를 모두 오버라이딩하지 않음
# 2. b에서만
# 3. c에서만
# 4. b, c 모두 오버라이딩 한다면?

# 1. 

d.method() # A 출력

# 2.

class B(A) :
    def method(self) :
        print("B")

d = D()
d.method() # B 출력

# 3.

class C(A) :
    def method(self) :
        print("C")

d = D()
d.method() # C 출력


# 4.

class B(A) :
    def method(self) :
        print("B")

class C(A) :
    def method(self) :
        print("C")

d = D()
d.method() # B 출력

# B가 먼저 D에서 불러오니까, B가 인쇄됨.
# class D(C, B) 라면 C가 출력됨.

#### multiple inheritance 를 지양하자 ####
다중 상속을 피하는 것이 좋다고 알 수 있다.



이번에는 오버라이딩 대신 오버로딩을 해 본다.

구분	오버로딩	오버라이딩
메소드 이름	동일	동일
매개변수와 타입	다름	동일
리턴타입	상관없음	동일
class Square :
    def __init__(self, side) :
        self.side = side
    def __add__(one, two) :
        return( (4 * one.side) + (4 * two.side) )

squareOne = Square(5)
squareTwo = Square(10)

squareOne + squareTwo
파이썬에서는 형이 없으니 오버로딩을 이용하여 형을 체크하는 것이 생기기도 하였다. 3.4부터 생긴 singledispatch를 확인해보면 된다.

from functools import singledispatch

@singledispatch
def fun(arg, verbose=False):
    if verbose:
        print('Let me just say,', end=' ')
    print(arg)

# .register(자료형) 으로 자료형을 체크한다.

@fun.register(int)
def _(arg, verbose=False):
    if verbose:
        print('Strength in numbers, eh?', end=' ')
    print(arg)

@fun.register(float)
def _(arg, verbose=False):
    if verbose:
        print('Half of your number:', end=' ')
    print(arg / 2)

@fun.register(type(None))
def _(arg):
    print('Nothing.')

"""
>>> fun('1', verbose=True)
Let me just say, 1

>>> fun(1, verbose=True)
Strength in numbers, eh? 1

>>> fun(1.0, verbose=True)
Half of your number: 0.5

>>> fun(None)
Nothing.
"""
