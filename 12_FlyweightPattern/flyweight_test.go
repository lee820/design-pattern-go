package flyweight

import (
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Run("Shape Factory test:", ShapeFactoryTest)
}

func ShapeFactoryTest(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	colors := []string{"red", "green", "blue", "white", "black"}
	shapeFactory := NewShapeFactory()

	for i := 0; i < 10; i++ {
		circle := shapeFactory.GetCircle(colors[r.Intn(len(colors))])
		circle.(*Circle).SetX(r.Intn(100))
		circle.(*Circle).SetY(r.Intn(100))
		circle.(*Circle).SetRadius(50)
		circle.Draw()
	}
}
