# LAMBDA & ANONYMOUS FUNCTION IN PYTHON

##   What are lambda functions in Python?

In Python, an anonymous function is a  [function](https://www.programiz.com/python-programming/function)  that is defined without a name. While normal functions are defined using the  `def`  keyword in Python, anonymous functions are defined using the  `lambda`  keyword.
Hence, anonymous functions are also called lambda functions.

**lambda == anonymous**
A **lambda abstraction** is another name for an anonymous function

## How to use lambda Functions in Python?

A lambda function in python has the following syntax.

```
# Program to show the use of lambda functions
double = lambda x: x * 2

print(double(5))

print((lambda x : x*2)(5))
```

**Output**

10
10

You can also put **Lambda in variable** .

```
lambdaAdd = lambda n,m:n+m
print(lambdaAdd(2,3))
print(lambdaAdd(4,5))
```  

Also, You can do **if eles in lambda**.

```
print((lambda n,m: n if n%2==0 else m)(1,3))
print((lambda n,m: n if n%2==0 else m)(2,3))
```

# ANONYMOUS FUNCTION IN GO

## What are anonymous functions in Go?

An anonymous function is a function which doesn’t contain any name. It is useful when you want to create an inline function. In Go language, an anonymous function can form a closure. An anonymous function is also known as  _**function literal**_.

**Syntax:**

```
package main

import "fmt"

func main() {

// Anonymous function
func(){
fmt.Println(``"Welcome! to GeeksforGeeks"``)
}()
}
```


## How to use anonymous Functions in Go?

In Go language, you are allowed to assign an anonymous function to a variable. When you assign a function to a variable, then the type of the variable is of function type and you can call that variable like a function call as shown in the below example.

**Example:**

```
package main
import "fmt"

func main() {

// Assigning an anonymous
// function to a variable
value := func(){

	fmt.Println(``"Welcome! to GeeksforGeeks"``)

}

value()

}
```

# But, Why use anonymous function?

1.  If no name is needed because the function is only ever called in one place, then why add a name to whatever namespace you're in.
2.  Anonymous functions are declared inline and inline functions have advantages in that they can access variables in the parent scopes. Yes, you can put a name on an anonymous function, but that's usually pointless if it's declared inline. So inline has a significant advantage and if you're doing inline, there's little reason to put a name on it.
3.  The code seems more self-contained and readable when handlers are defined right inside the code that's calling them. You can read the code in almost sequential fashion rather than having to go find the function with that name.

-> so ask yourself. **Will I use this function anywhere else?**

# I think this is related to Closure.

## Closures

**YES.** Function literals in Go are  **closures**: they may refer to variables defined in an enclosing function.

-   variables tare shared between the surrounding function and the function literal,
-   survive as long as they are accessible.

In this example, the function literal uses the local variable `n`  from the enclosing scope to count the number of times it has been invoked.

```
// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
    n := 0
    return func() {
        n++
        fmt.Println(n)
    }
}

func main() {
    f1, f2 := New(), New()
    f1() // 1
    f2() // 1 (different n)
    f1() // 2
    f2() // 2
```

But, What is Closure?
