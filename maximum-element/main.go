package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type commandType int

const (
	_                = iota
	push commandType = iota
	delete
	print
)

type stack []int

func (s *stack) pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

type max struct {
	max         int
	initialized bool
}

type command interface {
	run(io.Reader, stack, max) (stack, max, error)
}

type commandPush struct{}

func (commandPush) run(r io.Reader, s stack, m max) (stack, max, error) {
	n, err := readInt(r)
	if err != nil {
		return s, m, err
	}
	if !m.initialized || m.max < n {
		m.max = n
		m.initialized = true
	}

	s = append(s, n)
	return s, m, nil
}

type commandDelete struct{}

func (commandDelete) run(r io.Reader, s stack, m max) (stack, max, error) {
	s.pop()
	if len(s) == 0 {
		m.initialized = false
	} else {
		m.max = findMax(s)
	}
	return s, m, nil
}

type commandPrint struct{}

func (commandPrint) run(r io.Reader, s stack, m max) (stack, max, error) {
	fmt.Println(m.max)
	return s, m, nil
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

func findMax(s stack) int {
	tmp := make(stack, len(s))
	copy(tmp, s)
	sort.Ints(tmp)
	return tmp[len(tmp)-1]
}

func main() {
	var err error
	var c int
	s := make(stack, 0)
	var m max

	if _, err = readInt(os.Stdin); err != nil {
		log.Fatal("Error reading input data ", err)
	}

	for c, err = readInt(os.Stdin); err == nil; c, err = readInt(os.Stdin) {
		var cmd command
		if cmd, err = newCommand(commandType(c)); err != nil {
			log.Fatal(err)
		}
		if s, m, err = cmd.run(os.Stdin, s, m); err != nil {
			log.Fatal(err)
		}
	}
	if err != io.EOF {
		log.Fatal("Error reading input data ", err)
	}
}
