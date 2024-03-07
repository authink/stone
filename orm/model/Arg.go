package model

import "reflect"

type Arg interface {
	ArgType() string
}

type Argument struct{}

func (a *Argument) ArgType() string {
	return reflect.TypeOf(a).Name()
}
