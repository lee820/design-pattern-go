package decorator

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("Shape testing: ", ShapeTest)
	t.Run("Decorator testing: ", ShapeDecoratorTest)
}

func ShapeTest(t *testing.T) {
	circle := NewCircle()
	if circle != nil {
		circle.Draw()
	}
	rect := NewRectangle()
	if rect != nil {
		rect.Draw()
	}
}

func ShapeDecoratorTest(t *testing.T) {
	circle := NewCircle()
	rect := NewRectangle()

	redCircle := NewRedShapeDecorator(circle)
	redRectangle := NewRedShapeDecorator(rect)

	redCircle.Draw()
	redRectangle.Draw()
}
