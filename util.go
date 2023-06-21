package golib

import (
	"fmt"
)

const PACKAGENAME = "golib/PACKAGENAME"

type MathUtil struct {
	LoggingEnabled bool
}

type Add interface {
	Add(a int, b int) int
}

type Subtract interface {
	Subtract(a int, b int) int
}

type Multiply interface {
	Multiply(a int, b int) int
}

func (mathUtil *MathUtil) Add(a int, b int) int {
	if mathUtil.LoggingEnabled {
		fmt.Println("[Add operation started.]")
	}
	res := a + b

	if mathUtil.LoggingEnabled {
		fmt.Println("[Add operation ended.]")
	}

	return res
}

func (mathUtil *MathUtil) Subtract(a int, b int) int {
	if mathUtil.LoggingEnabled {
		fmt.Println("[Subtract operation started.]")
	}
	res := a - b
	if mathUtil.LoggingEnabled {
		fmt.Println("[Subtract operation ended.]")
	}
	return res
}

func (mathUtil *MathUtil) Multiply(a int, b int) int {
	if mathUtil.LoggingEnabled {
		fmt.Println("[Multiply operation started.]")
	}
	res := a * b
	if mathUtil.LoggingEnabled {
		fmt.Println("[Multiply operation ended.]")
	}
	return res
}
