# Marco's Golang Playground

Well, indeed this is no real [Golang Playground](https://play.golang.org) since I advocate testable code over executable code.  
At least for now (2022-01-13Z) Golang Playgrounds do not offer this kind of workflow, but I won't change my behavior for them.

# What's inside of this repository?

## `error[_test].go`

With these files I fiddle around with the possibilities and restrictions of Golang `error`s.  
Since I'm used to chain a bunch of error types together I somewhat miss that feature in Golang.  
In addition the approach of `defer`ring error-prone functions and let them manipulate the named `error` return value throws away any previously error which may have lead to the `defer` call.


