package err

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
)

type Error struct {
	Callstr string
	E error
}

func New(msg string)Error{
	_, file, line, _ := runtime.Caller(1)
	callStr := file + ": " + "line " + strconv.Itoa(line)
	err := errors.New(msg)
	return Error{
		callStr,
		err,
	}
}

func NewFromErr(e error)Error{
	_, file, line, _ := runtime.Caller(1)
	callStr := file + ": " + "line " + strconv.Itoa(line)
	return Error{
		callStr,
		e,
	}
}

func (e Error) Log(){
	msg := fmt.Errorf(e.Callstr + " %w", e.E)
	log.Println(msg)
}