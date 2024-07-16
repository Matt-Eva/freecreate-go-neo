package err

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
)

type Error struct {
	Callstrs []string
	E       error
}

func New(msg string) Error {
	callStrs := callers()
	err := errors.New(msg)
	return Error{
		callStrs,
		err,
	}
}

func NewFromErr(e error) Error {
	callStrs := callers()
	return Error{
		callStrs,
		e,
	}
}

func (e Error) Log() {
	calls := ""
	for _, call := range e.Callstrs{
		calls += call + "\n"
	}
	msg := fmt.Errorf("ERROR: %w\n" + calls, e.E)
	log.Println(msg)
}

func callers()[]string{
	pc := make([]uintptr, 50)
	callers := runtime.Callers(2, pc)
	callStrs := make([]string,0)
	for i := 2; i <=callers; i++{
		_, file, line, _ := runtime.Caller(i)
		callStr := file + ": " + "line " + strconv.Itoa(line)
		callStrs = append(callStrs, callStr)
	}
	return callStrs
}
