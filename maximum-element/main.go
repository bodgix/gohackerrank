package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	bst "github.com/bodgix/gobst"
)

type commandType int

const (
	_                = iota
	push commandType = iota
	delete
	print
)

type stack []int

func (s *stack) pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

type command interface {
	run(io.Reader, stack, bst.Bst) (stack, bst.Bst, error)
}

type commandPush struct{}

func (commandPush) run(r io.Reader, s stack, t bst.Bst) (stack, bst.Bst, error) {
	n, err := readInt(r)
	if err != nil {
		return s, t, err
	}
	t.Insert(n)

	s = append(s, n)
	return s, t, nil
}

type commandDelete struct{}

func (commandDelete) run(r io.Reader, s stack, t bst.Bst) (stack, bst.Bst, error) {
	value := s.pop()
	t.Delete(value)
	return s, t, nil
}

type commandPrint struct{}

func (commandPrint) run(r io.Reader, s stack, t bst.Bst) (stack, bst.Bst, error) {
	fmt.Println(t.Max())
	return s, t, nil
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
	t := bst.NewBst()

	if _, err = readInt(os.Stdin); err != nil {
		log.Fatal("Error reading input data ", err)
	}

	for c, err = readInt(os.Stdin); err == nil; c, err = readInt(os.Stdin) {
		var cmd command
		if cmd, err = newCommand(commandType(c)); err != nil {
			log.Fatal(err)
		}
		if s, t, err = cmd.run(os.Stdin, s, t); err != nil {
			log.Fatal(err)
		}
	}
	if err != io.EOF {
		log.Fatal("Error reading input data ", err)
	}
}
