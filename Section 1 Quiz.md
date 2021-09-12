# Quiz - Section 1

**Q1**: When was released the first stable version of Go?

- 2009
- 2010
- 2012

**A1**: 2012

---

**Q2**: What is $GOPATH?

- A directory structure
- Where Go code lives
- Where third party packages live
- Where binaries live
- All the above

**A2**: All the above

---

**Q3**: Do we need $GOPATH to write Go programs?

- Yes
- No
- Only the `pkg` and `bin` directories from $GOPATH

**A3**: Only the `pkg` and `bin` directories from $GOPATH

---

**Q4**: Which one of the following import lines matches the correct $GOPATH structure:

- github.com/foo/bar
- steevehook.com/q/w/e/r/t/y
- 0.0.0.0/w/t/f
- All the above

**A4**: All the above

---

**Q5**: Why do we need the `bin` directory

- To store Go binaries
- To execute Go binaries globally
- All the above

**A5**: All the above

---

**Q6**: What is the role of `vendor` directory?

- To store project packages
- To store third party packages inside a specific project
- To replace the `src` directory

**A6**: To store third party packages inside a specific project

---

**Q7**: What is the difference between `run` and `build` commands?

- Named executables
- `run` executes faster
- `run` compiles and executes the generated binary, `build` only compiles and generates a binary
- They are the same, just different names

**A7**: `run` compiles and executes the generated binary, `build` only compiles and generates a binary

---

**Q8**: Why is the Go binary bigger than the C binary?

- The Go compiler & linker are inefficient
- The Go binary has types, Garbage Collection, Scheduler packed in
- C is a much more primitive language

**A8**: The Go binary has types, Garbage Collection, Scheduler packed in
