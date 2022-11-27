package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
)

//struture OutputStack to store the index and the character it got interpreted
type OutputStack struct {
	Letter string
	Index  int
}

var response []OutputStack

type Stack []uint16

func (s *Stack) Push(v uint16) {
	*s = append(*s, v)
}

func (s *Stack) Pop() uint16 {
	l := len(*s)
	if l > 0 {
		op := (*s)[l-1]
		*s = (*s)[:l-1]
		return op
	}
	return 0
}

func (s *Stack) Top() uint16 {
	n := len(*s) - 1
	return (*s)[n]
}

func (s Stack) Len() int {
	return len(s)
}
const (
	arraylen = 65536
)

func main() {
	//Get all the optional input
	/*
		-filename string <default:input.bf>
		-input string <default:"">
	*/
	filename := flag.String("filename", "input.bf", "Enter input file name to interpret")

	us_inp := flag.String("input", "", "Enter input character if required")

	flag.Parse()

	/*
	    -input:
	   	*filename name that defined in var:filename
	*/
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("found error in reading input file " + err.Error())
		return
	}
	defer file.Close()
	/*
	    -input:
	   	*data from input defined filename
	   	*us_inp input defined to add in "," as putchar
	*/
	output, err := interpret_bf(file, *us_inp)
	if err != nil {
		fmt.Println("found error while interpreting input: " + err.Error())
		return
	}
	/*
	    -input:
	   	output - oupt generated from the function interpret_bf
	*/
	err = writeOutputData(output)
	if err != nil {
		fmt.Println("found error while interpreting input " + err.Error())
		return
	}
	fmt.Println("The given input successfully interpreted in output.txt")
}

//This will be the major function for that perform as interpreter
//The terms are followed that defined in
func interpret_bf(inputfile io.Reader, readinput string) (string, error) {
	result := ""
	i := 0
	buf := new(strings.Builder)
	_, err := io.Copy(buf, inputfile)
	if err!=nil{
		return result, err
	}
	input:=buf.String()
	data := make([]int16, arraylen)
	var bf_ptr, data_ptr uint16 = 0, 0
	var bf_stack Stack
	//Iterating input character one by one
	for i < len(input) {
		switch input[i] {
		case '>':
			data_ptr++
		case '<':
			data_ptr--
		case '+':
			data[data_ptr]++
		case '-':
			data[data_ptr]--
		case '.':
			response = append(response, OutputStack{string(rune(data[data_ptr] % 128)), i})
			result += string(rune(data[data_ptr] % 128))
		case ',':
			if len([]byte(readinput)) > 0 {
				data[data_ptr] = int16([]byte(readinput)[0])
			} else {
				data[data_ptr] = 0
			}
		case '[':
			bf_stack.Push(uint16(i))
		case ']':
			if len(bf_stack) == 0 {
				return result, errors.New("Compilation error.")
			}
			bf_ptr = bf_stack.Pop()
			
			if data[data_ptr] != 0 {
				i = int(bf_ptr) - 1
			}
		}
		i++
	}
	if len(bf_stack) != 0 {
		return result, errors.New("Compilation error.")
	}
	return result, nil
}

//This fucntion will be use to writer interpreted from input
func writeOutputData(output string) error {
	err := os.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	return nil
}