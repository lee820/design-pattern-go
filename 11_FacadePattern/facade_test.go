package facade

import "testing"

func Test(t *testing.T) {
	t.Run("Test ShapeMaker:", ShapeMakerTest)
}

func ShapeMakerTest(t *testing.T) {
	shapeMaker := NewShapeMaker()
	shapeMaker.DrawCircle()
	shapeMaker.DrawRectangle()
	shapeMaker.DrawSquare()
}
