package abstractfactory

import "fmt"

//Shape 模型接口
type Shape interface {
	Draw()
}

//Color 色彩接口
type Color interface {
	Fill()
}

//Circle 实现模型接口的圆形类
type Circle struct{}

//Square 实现模型接口的正方形类
type Square struct{}

//Draw Circle类的Draw方法
func (c Circle) Draw() {
	fmt.Println("Circle Draw() method")
}

//Draw Square类的Draw方法
func (s Square) Draw() {
	fmt.Println("Square Draw() method")
}

//Red 实现色彩接口的红色类
type Red struct{}

//Green 实现色彩接口的绿色类
type Green struct{}

//Fill 红色类的Fill方法
func (r Red) Fill() {
	fmt.Println("Red Fill() method")
}

//Fill 绿色类的Fill方法
func (g Green) Fill() {
	fmt.Println("Green Fill() method")
}

//AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(colorName string) Color
}

//ShapeFactory 模型工厂的类
type ShapeFactory struct{}

//ColorFactory 色彩工厂的类
type ColorFactory struct{}

//GetShape 模型工厂实例获取模型子类的方法
func (sh ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

//GetColor 模型工厂实例不需要色彩方法
func (sh ShapeFactory) GetColor(colorName string) Color {
	return nil
}

//GetShape 色彩工厂实例不需要获取模型方法
func (cf ColorFactory) GetShape(shapeName string) Shape {
	return nil
}

//GetColor 色彩工厂实例，获取具体色彩子类
func (cf ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

//FactoryProducer 超级工厂类，用于获取工厂实例
type FactoryProducer struct{}

//GetFactory 获取工厂方法
func (fp FactoryProducer) GetFactory(factoryname string) AbstractFactory {
	switch factoryname {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

//NewFactoryProducer 创建FactoryProducer类
func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}
