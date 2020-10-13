package state

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Context test:", StateTest)
}

func StateTest(t *testing.T) {
	context := NewContext()

	startState := NewStatartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())

	stopState := NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
