package factory

import "fmt"

//Shape 模型接口
type Shape interface {
	Draw()
}

//Circle 圆形类
type Circle struct{}

//Square 正方形类
type Square struct{}

//Rectangle 矩形类
type Rectangle struct{}

//NewCircle 实例化Circle类
func NewCircle() *Circle {
	return &Circle{}
}

//Draw Circle实现Shape接口的Draw方法
func (c *Circle) Draw() {
	fmt.Println("Circle Draw() method.")
}

//NewSquare 实例化Square类
func NewSquare() *Square {
	return &Square{}
}

//Draw Square实现Shape接口的Draw方法
func (s *Square) Draw() {
	fmt.Println("Square Draw() method.")
}

//NewRectangle 实例化Rectangle类
func NewRectangle() *Rectangle {
	return &Rectangle{}
}

//Draw Rectangle实现Shape接口的Draw方法
func (r *Rectangle) Draw() {
	fmt.Println("Rectangle Draw() method.")
}

//ShapeFactory 模型工厂类
type ShapeFactory struct{}

//NewShapeFactory 实例化模型工厂
func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{}
}

//GetShape 从模型工厂获取指定的模型类
func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return NewCircle()
	case "square":
		return NewSquare()
	case "rectangle":
		return NewRectangle()
	default:
		return nil
	}
}
