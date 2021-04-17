# CLOSURE

## What is Closure

A closure is a special type of anonymous function that references variables declared outside of the function itself. That means that using variables that weren’t passed into the function as a parameter, but instead were available when the function was declared.

This is very similar to how a regular function can reference global variables. You aren’t directly passing these variables into the function as a parameter, but the function has access to them when it is called.

Let’s take a look at a closure in action. In this next example we are going to create a closure that keeps track of how many times it has been called, and returns that number.

```go
import "fmt"

func main() {
  n := 0
  counter := func() int {
    n += 1
    return n
  }
  fmt.Println(counter()) // 1
  fmt.Println(counter()) // 2
}
```

Notice how our anonymous function has access to the `n` variable, but this was never passed in as a parameter when `counter()` was called. This is what makes it a closure!

## Why we use Closure

#### Closures provide data isolation

One problem with the previous example is a problem that can also pop up when using global variables. Any code inside of the  `main()`  function has access to  `n`, so it is possible to increment the counter without actually calling  `counter()`. That isn’t what we want; Instead we would rather have  `n`  isolated so that no other code has access to it.

To do this we need to look at another interesting aspect of closures, which is the fact that they can still reference variables that they had access to during creation even if those variables are no longer referenced elsewhere.

```go
package main

import "fmt"

func main() {
  counter := newCounter()
  fmt.Println(counter())
  fmt.Println(counter())
}

func newCounter() func() int {
  n := 0
  return func() int {
    n += 1
    return n
  }
}
```

In this example our closure references the variable `n` even after the `newCounter()` function has finished running. This means that our closure has access to a variable that keeps track of how many times it has been called, but no other code outside of the `newCounter()` function has access to this variable. This is one of the many benefits of a closure - we can persist data between function calls while also isolating the data from other code.

The closure property is used extensively where data isolation is required. The state provided by the closures makes them immensely helpful in that regard. When we want to create a state encapsulated function we use closures to do that.

---

```go
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	printInt := intSeq()
	fmt.Println(printInt()) // 1
	fmt.Println(printInt()) // 2

	printInt2 := intSeq()
	fmt.Println(printInt2()) // 1
	fmt.Println(printInt2()) // 2
}
```
