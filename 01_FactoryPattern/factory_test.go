package factory

import "testing"

func Test(t *testing.T) {
	t.Run("Shape Factory:", ShapeFactoryTest)
}

func ShapeFactoryTest(t *testing.T) {
	sf := NewShapeFactory()

	circle := sf.GetShape("circle")
	if circle != nil {
		circle.Draw()
	}

	square := sf.GetShape("square")
	if square != nil {
		square.Draw()
	}

	rect := sf.GetShape("rectangle")
	if rect != nil {
		rect.Draw()
	}
}
