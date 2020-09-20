package facade

import "fmt"

//Shape 模型接口
type Shape interface {
	Draw()
}

//Circle 圆形类
type Circle struct{}

//Rectangle 矩形类
type Rectangle struct{}

//Square 正方形类
type Square struct{}

//NewCircle 实例化圆形类
func NewCircle() *Circle {
	return &Circle{}
}

//Draw 实现Shape接口
func (c *Circle) Draw() {
	fmt.Println("Circle Draw method.")
}

//NewRectangle 实例化矩形类
func NewRectangle() *Rectangle {
	return &Rectangle{}
}

//Draw 矩形类实现Shape接口
func (r *Rectangle) Draw() {
	fmt.Println("Rectangle Draw method.")
}

//NewSquare 实例化正方形类
func NewSquare() *Square {
	return &Square{}
}

//Draw 正方形类实现Shape接口
func (s *Square) Draw() {
	fmt.Println("Square Draw method.")
}

//ShapeMaker 外观类
type ShapeMaker struct {
	circle    Circle
	square    Square
	rectangle Rectangle
}

//NewShapeMaker 实例化外观类
func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle:    Circle{},
		rectangle: Rectangle{},
		square:    Square{},
	}
}

//DrawCircle 调用circle的Draw方法
func (shapeMk *ShapeMaker) DrawCircle() {
	shapeMk.circle.Draw()
}

//DrawRectangle 调用rectangle的Draw方法
func (shapeMk *ShapeMaker) DrawRectangle() {
	shapeMk.rectangle.Draw()
}

//DrawSquare 条用square的Draw方法
func (shapeMk *ShapeMaker) DrawSquare() {
	shapeMk.square.Draw()
}
