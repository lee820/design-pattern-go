package builder

import (
	"container/list"
	"fmt"
)

//Item 所有商品的接口
type Item interface {
	Name() string
	Price() float32
}

//Food 实物的接口
type Food interface {
	Kind() string
}

//Staple 主食，实物接口的实例
type Staple struct{}

//Drink 饮料，实物接口的实例
type Drink struct{}

//Snack 小吃，实物接口的实例
type Snack struct{}

//Kind 获取主食类型名称
func (staple Staple) Kind() string {
	return "staple"
}

//Kind 获取饮料类型名称
func (drink Drink) Kind() string {
	return "drink"
}

//Kind 获取小吃类型名称
func (snack Snack) Kind() string {
	return "snack"
}

//CocaCola 可口可乐商品类，实现Item接口，并且组合了Drink基类
type CocaCola struct {
	Drink
}

//NewCocaCola 实例化可口可乐
func NewCocaCola() *CocaCola {
	return &CocaCola{}
}

//Name 获取可口可乐商品名
func (co *CocaCola) Name() string {
	return "coca-cola"
}

//Price 获取可口可乐价格
func (co *CocaCola) Price() float32 {
	return 2.5
}

//PepsiCola 百事可乐商品类，实现Item接口，并且组合了Drink基类
type PepsiCola struct {
	Drink
}

//NewPepsiCOla 实例化百事可乐
func NewPepsiCOla() *PepsiCola {
	return &PepsiCola{}
}

//Name 获取百事可乐商品名
func (pe *PepsiCola) Name() string {
	return "coca-cola"
}

//Price 获取百事可乐价格
func (pe *PepsiCola) Price() float32 {
	return 2.76
}

//Chips 薯条商品类，实现Item接口，组合Snack基类
type Chips struct {
	Snack
}

//NewChips 实例化薯条类
func NewChips() *Chips {
	return &Chips{}
}

//Name 获取薯条的名称
func (ch *Chips) Name() string {
	return "chips"
}

//Price 获取薯条价格
func (ch *Chips) Price() float32 {
	return 9.9
}

//Taco Taco Tuesday!!!
type Taco struct {
	Staple
}

//NewTaco 实例化taco
func NewTaco() *Taco {
	return &Taco{}
}

//Name 获取Taco名称
func (t *Taco) Name() string {
	return "taco"
}

//Price 获取Taco价格
func (t *Taco) Price() float32 {
	return 16.8
}

//BeefBurger 牛肉汉堡，实现Item接口，组合Staple基类
type BeefBurger struct {
	Staple
}

//NewBeefBurger 实例化牛肉汉堡
func NewBeefBurger() *BeefBurger {
	return &BeefBurger{}
}

//Name 获取牛肉汉堡名称
func (bf *BeefBurger) Name() string {
	return "beefBurger"
}

//Price 获取牛肉汉堡价格
func (bf *BeefBurger) Price() float32 {
	return 19.8
}

//Meal 点单类
type Meal struct {
	Items *list.List
}

//NewMeal 实例化点单类
func NewMeal() *Meal {
	l := list.New()
	return &Meal{l}
}

//AddItem 添加新的商品
func (m *Meal) AddItem(item Item) {
	m.Items.PushBack(item)
}

//GetCost 获取当前订单总额
func (m *Meal) GetCost() float32 {
	var cost float32 = 0
	for i := m.Items.Front(); i != nil; i = i.Next() {
		item, ok := (i.Value).(Item)
		if ok {
			cost += item.Price()
		}
	}
	return cost
}

//ShowItems 显示当前所有餐品
func (m *Meal) ShowItems() {
	for i := m.Items.Front(); i != nil; i = i.Next() {
		item, ok := (i.Value).(Item)
		if ok {
			fmt.Println(item.Name(), ": ", item.Price(), "元")
		}
	}
}
