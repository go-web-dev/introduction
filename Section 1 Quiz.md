# Quiz - Section 1

**Q1**: When was released the first stable version of Go?

1) 2009
   <br/> *Explanation*: This is when Google adopted and open sourced the language, not when we got Go major version 1
2) 2010
   <br/> *Explanation*: Still 2 more years before Go reaches major version 1
3) 2012
   <br/> *Explanation*: Correct. This is the year Go reached major version 1, which means no breaking changes till this
   point

**Correct Answer**: `2`

---

**Q2**: What is $GOPATH?

1) A directory structure
   <br/> *Explanation*: True, but have a look at all the options
2) Where Go code lives
   <br/> *Explanation*: Yes, while this is legacy, it still works. But it's not only about that
3) Where third party packages live
   <br/> *Explanation*: Yes, but not only for that
4) Where binaries live
   <br/> *Explanation*: Yes, but not only for that
5) All the above
   <br/> *Explanation*: Correct. GOPATH is a directory structure where we can store code, binaries and third party packages

**Correct Answer**: `5`

---

**Q3**: Do we need $GOPATH to write Go programs?

1) Yes
   <br/> *Explanation*: We don't really need the src directory with Go modules enabled
2) No
   <br/> *Explanation*: We still need at least the src and bin directories used by go get and go install commands
3) Only the `pkg` and `bin` directories from $GOPATH, if the Go version is > 1.11
   <br/> *Explanation*: Correct, as of Go 1.11, these are the only relevant directories needed

**Correct Answer**: `3`

---

**Q4**: Which one of the following import lines matches the correct $GOPATH structure:

1) github.com/foo/bar
   <br/> *Explanation*: We have a valid host/user/project path, but check all the answers
2) steevehook.com/q/w/e/r/t/y
   <br/> *Explanation*: We have a valid host/user/project/package path, but check all the answers
3) 0.0.0.0/w/t/f
   <br/> *Explanation*: 0.0.0.0 is also a host, but check all the answers
4) All the above
   <br/> *Explanation*: Correct, all the above are valid import paths that match the GOPATH structure

**Correct Answer**: `4`

---

**Q5**: Why do we need the `bin` directory

1) To store Go binaries
   <br/> *Explanation*: Correct, but check all the answers
2) To execute Go binaries globally
   <br/> *Explanation*: Correct, only if you set the $PATH variable to point to the $GOBIN variable, but check all the answers
3) All the above
   <br/> *Explanation*: Correct. If we configure $GOBIN and $PATH variables to point to the bin directory, we can store and run Go binaries globally

**Correct Answer**: `3`

---

**Q6**: What is the role of `vendor` directory?

1) To store project packages
   <br/> *Explanation*: We can store packages that belong to a project anywhere inside the project itself
2) To store third party packages inside a specific project
   <br/> *Explanation*: Correct. If we create a directory named vendor inside the project, the packages from the vendor directory will get picked up instead of the ones stored in the src directory
3) To replace the `src` directory
   <br/> *Explanation*: The src directory was not replaced, it was rather made obsolete by introducing Go Modules

**Correct Answer**: `2`

---

**Q7**: What is the difference between `run` and `build` commands?


1) `run` executes faster
   <br/> *Explanation*: Not necessarily. They run almost at the same speed, going through almost the same process
2) `run` compiles and executes the generated binary, `build` only compiles and generates a binary
   <br/> *Explanation*: Correct. The run command will compile and run the generate binary, while build will only compile and generate a binary
3) They are the same, just different names
   <br/> *Explanation*: They are not the same. The run command compiles and runs the generated binary, while the build command only compiles and generates the binary

**Correct Answer**: `2`

---

**Q8**: Why is the Go binary bigger than the C binary?

1) The Go compiler and linker are inefficient
   <br/> *Explanation*: It's hard to tell. The Go compiler and linker are inspired from the C one, and they are equally efficient
2) The Go binary has a rich runtime packed into the final binary
   <br/> *Explanation*: Correct. The linker will not only use the object created in the compilation phase, but also pack in the GC, Go Scheduler, types and the entire Go runtime
3) C is a much more primitive language
   <br/> *Explanation*: Correct. C is a more primitive language, but that has to do less with the final generated binary

**Correct Answer**: `2`
