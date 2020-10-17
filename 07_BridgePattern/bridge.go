package bridge

import "fmt"

//DrawAPI 画图抽象接口，桥接模式的抽象接口
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}

//RedCircle 红色圆的类，桥接模式接口
type RedCircle struct{}

//GreenCircle 绿色圆的实体类，桥接模式接口
type GreenCircle struct{}

//NewRedCircle 实例化红色圆
func NewRedCircle() *RedCircle {
	return &RedCircle{}
}

//DrawCircle 红色圆实现DrawAPI方法
func (rc *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

//NewGreenCircle 实例化绿色圆
func NewGreenCircle() *GreenCircle {
	return &GreenCircle{}
}

//DrawCircle 绿色圆实现DrawAPI方法
func (gc *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

//ShapeCircle 桥接模式的实体类
type ShapeCircle struct {
	Radius  int
	X       int
	Y       int
	drawAPI DrawAPI
}

//NewShapeCircle 实例化桥接模式实体类
func NewShapeCircle(radius, x, y int, drawAPI DrawAPI) *ShapeCircle {
	return &ShapeCircle{
		Radius:  radius,
		X:       x,
		Y:       y,
		drawAPI: drawAPI,
	}
}

//Draw 实体类的Draw方法
func (sc *ShapeCircle) Draw() {
	sc.drawAPI.DrawCircle(sc.Radius, sc.X, sc.Y)
}
