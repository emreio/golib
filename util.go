package golib

import (
	"fmt"
)

const PACKAGENAME = "golib/PACKAGENAME"

// Root struct
type MathUtil struct {
	LoggingEnabled bool
}

// Aritmetic Add Operation Interface
type Add interface {
	Add(a int, b int) int
}

type Subtract interface {
	Subtract(a int, b int) int
}

type Multiply interface {
	Multiply(a int, b int) int
}

type Divide interface {
	Divide(a int, b int) (int, error)
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

func (mathUtil *MathUtil) Divide(a int, b int) (int, error) {
	return 1, nil
}
