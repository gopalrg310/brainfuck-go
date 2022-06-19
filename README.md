Brainfuck in go
Brainfuck compiler in Go - well not quite. A compiler outputs binaries. It's in between an interpreter and a compiler.

Intro
This is the possible implementation of a Brainfuck interpreter by any means. Nor it is the fastest.

The goal from this was to learn a bit about how interpreters and compilers work. Using Brainfuck looked to be the simplest way to achieve this.


 try this `go get github.com/gravip214/brainfuck-go`
 
Run below command to run the program

`go build`

`./brainfuck-go -h`

Usage of ./BF_Interpreter:
  -filename string
    	Enter input file name to interpret (default "input.bf")
  -input string
    	Enter input character if required

`./brainfuck-go -filename input.bf -input m`

To run the unit test

`go test -v`
