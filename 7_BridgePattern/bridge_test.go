package bridge

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("start: ", NewShapeCircleTest)
}

func NewShapeCircleTest(t *testing.T) {
	redCircle := NewShapeCircle(5, 6, 8, NewRedCircle())
	if redCircle != nil {
		redCircle.Draw()
	} else {
		fmt.Println("red circle test fail.")
	}
	greenCircle := NewShapeCircle(1, 2, 4, NewGreenCircle())
	if greenCircle != nil {
		greenCircle.Draw()
	} else {
		fmt.Println("green circle test fail.")
	}
}
