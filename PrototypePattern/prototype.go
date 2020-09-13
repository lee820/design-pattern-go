package prototype

import "fmt"

//ShapeClone 有克隆方法的接口
type ShapeClone interface {
	Clone() ShapeClone
}

//ShapeManager 模型管理类，用于保存所有Clone接口的原型
type ShapeManager struct {
	ShapeList map[string]ShapeClone
}

//NewShapeManager 新建模型管理类
func NewShapeManager() *ShapeManager {
	return &ShapeManager{
		ShapeList: make(map[string]ShapeClone),
	}
}

//Set 添加原型
func (s *ShapeManager) Set(name string, proto ShapeClone) {
	s.ShapeList[name] = proto
}

//Get 获取原型
func (s *ShapeManager) Get(name string) ShapeClone {
	return s.ShapeList[name]
}

//Circle 圆形类，实现shapeclone接口
type Circle struct {
	Name string
}

//NewCircle 新建一个CIrcle类
func NewCircle(name string) *Circle {
	return &Circle{Name: name}
}

//GetName 获取circle的名字
func (c *Circle) GetName() string {
	return c.Name
}

//Clone 返回circle类的复制
func (c *Circle) Clone() ShapeClone {
	circleClone := c
	return circleClone
}

//Draw circle的方法
func (c *Circle) Draw() {
	fmt.Println("Circle Draw().")
}

//Square 正方形模型类
type Square struct {
	Name string
}

//NewSquare 新建一个Square类
func NewSquare(name string) *Square {
	return &Square{Name: name}
}

//GetName 获取circle的名字
func (s *Square) GetName() string {
	return s.Name
}

//Draw square类的Draw方法
func (s *Square) Draw() {
	fmt.Println("Square Draw().")
}

//Clone 返回circle类的复制
func (s *Square) Clone() ShapeClone {
	squareClone := s
	return squareClone
}
