# Flag Handling of Golang Standard Library `flag`

TIL that `cmd -flag value` and `cmd --flag=value` behave equally.  
So I wonder why one would need external packages like `pflag`, since Go 1.16 also provides a `flag.Func()` interface for handling specific parsing like _2s -> 2 * time.Second_. See [Alex Edwards blog entry on custom CLI flags](https://www.alexedwards.net/blog/custom-command-line-flags)


## Known Issues

The `test.sh` script expects you're on an Apple Mac M1 with `go` installed via `homebrew`.  
If you're not, please change the `GO_EXEC` variable to the according location.

# Sample Output
Running the (reconfigured) test script shows:

```
./test.sh
'-foo baz' translates to 'foo: baz'
'--foo=qux' translates to 'foo: qux'
```
