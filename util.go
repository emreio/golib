package golib

const PACKAGENAME = "golib/PACKAGENAME"

type MathUtil struct {
   loggingEnabled bool
}

func (mathUtil *MathUtil) Add(a int, b int) int {
	 if mathUtil.loggingEnabled {
	 	fmt.Println("[Add operation started.]")
	 }
	return a + b
}

func (mathUtil *MathUtil) Subtract(a int, b int) int {
	 if mathUtil.loggingEnabled {
		fmt.Println("[Subtract operation started.]")
	 }
	return a - b
}
