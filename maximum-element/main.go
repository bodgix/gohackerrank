package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type commandType int

const (
	_                = iota
	push commandType = iota
	delete
	print
)

type stackElement struct {
	value       int
	currentMax  int
	initialized bool
}

type stack []stackElement

func (s *stack) pop() stackElement {
	var val stackElement
	if len(*s) > 0 {
		val = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	} else {
		val = stackElement{}
	}
	return val
}

func (s stack) peek() stackElement {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return stackElement{}
}

type command interface {
	run(io.Reader, stack) (stack, error)
}

type commandPush struct{}

func (commandPush) run(r io.Reader, s stack) (stack, error) {
	n, err := readInt(r)
	if err != nil {
		return s, err
	}
	var newMax int
	last := s.peek()
	if !last.initialized || last.currentMax < n {
		newMax = n
	} else {
		newMax = last.currentMax
	}
	s = append(s, stackElement{n, newMax, true})
	return s, nil
}

type commandDelete struct{}

func (commandDelete) run(r io.Reader, s stack) (stack, error) {
	s.pop()
	return s, nil
}

type commandPrint struct{}

func (commandPrint) run(r io.Reader, s stack) (stack, error) {
	lastElement := s.peek()
	fmt.Println(lastElement.currentMax)
	return s, nil
}

func newCommand(c commandType) (command, error) {
	switch c {
	case push:
		return commandPush{}, nil
	case delete:
		return commandDelete{}, nil
	case print:
		return commandPrint{}, nil
	default:
		return nil, fmt.Errorf("Unknown command %d", c)
	}
}

func readInt(r io.Reader) (int, error) {
	var d int
	n, err := fmt.Fscanf(r, "%d", &d)
	if err != nil {
		return d, err
	}
	if n != 1 {
		return d, errors.New("Expected 1 result")
	}
	return d, nil
}

func main() {
	var err error
	var c int
	s := make(stack, 0)

	if _, err = readInt(os.Stdin); err != nil {
		log.Fatal("Error reading input data ", err)
	}

	for c, err = readInt(os.Stdin); err == nil; c, err = readInt(os.Stdin) {
		var cmd command
		if cmd, err = newCommand(commandType(c)); err != nil {
			log.Fatal(err)
		}
		if s, err = cmd.run(os.Stdin, s); err != nil {
			log.Fatal(err)
		}
	}
	if err != io.EOF {
		log.Fatal("Error reading input data ", err)
	}
}
