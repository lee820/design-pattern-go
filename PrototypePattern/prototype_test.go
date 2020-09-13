package prototype

import (
	"fmt"
	"testing"
)

var shapManager *ShapeManager

func Test(t *testing.T) {
	t.Run("添加模型类的原型:", AddShape)
	t.Run("获取模型类的原型:", GetShape)
}

func AddShape(t *testing.T) {
	shapManager = NewShapeManager()
	circleProto := NewCircle("circle")
	squareProto := NewSquare("square")

	shapManager.Set("circle", circleProto)
	shapManager.Set("square", squareProto)
}

func GetShape(t *testing.T) {
	cProto := shapManager.Get("circle")
	if cProto == nil {
		fmt.Println("获取circle原型失败.")
	} else {
		cl := cProto.Clone().(*Circle) //cl获取Circle的原型
		fmt.Println("Cirle name: ", cl.GetName())
		cl.Draw()
	}

	sProto := shapManager.Get("square")
	if sProto == nil {
		fmt.Println("获取square原型失败.")
	} else {
		sl := sProto.Clone().(*Square)
		fmt.Println("Square name: ", sl.GetName())
		sl.Draw()
	}
}
