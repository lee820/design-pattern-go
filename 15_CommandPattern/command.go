package command

import (
	"container/list"
	"fmt"
)

//Order 订单命令接口
type Order interface {
	Execute()
}

//Stock 股票类
type Stock struct {
	Name     string
	Quantity int
}

//NewStock 实例化一只股票
func NewStock() *Stock {
	return &Stock{
		Name:     "maotai",
		Quantity: 100,
	}
}

//Buy 股票的买方法
func (s *Stock) Buy() {
	fmt.Printf("Stock [ Name: %s, Quantity: %d ] bought.\n",
		s.Name,
		s.Quantity)
}

//Sell 股票的卖方法
func (s *Stock) Sell() {
	fmt.Printf("Stock [ Name: %s, Quantity: %d ] sold.\n",
		s.Name,
		s.Quantity)
}

//BuyStock 买股票类，实现了Order接口的订单类
type BuyStock struct {
	StockToBy Stock
}

//NewBuyStock 实例化买股票订单类
func NewBuyStock(st Stock) *BuyStock {
	return &BuyStock{
		StockToBy: st,
	}
}

//Execute 处理买股票订单
func (bs *BuyStock) Execute() {
	bs.StockToBy.Buy()
}

//SellStock 卖股票类，实现了Order接口的订单类
type SellStock struct {
	StockToSell Stock
}

//NewSellStock 实例化卖股票订单类
func NewSellStock(st Stock) *SellStock {
	return &SellStock{
		StockToSell: st,
	}
}

//Execute 处理卖股票订单
func (ss *SellStock) Execute() {
	ss.StockToSell.Sell()
}

//Broker 创建命令调用类
type Broker struct {
	OrderList *list.List
}

//NewBroker 实例化命令调用类
func NewBroker() *Broker {
	return &Broker{
		OrderList: list.New(),
	}
}

//TakeOrder 添加命令
func (b *Broker) TakeOrder(order Order) {
	b.OrderList.PushBack(order)
}

//PlaceOrders 处理命令
func (b *Broker) PlaceOrders() {
	for i := b.OrderList.Front(); i != nil; {
		//处理完一个指令，就将该指令删除
		nextOrder := i.Next()
		i.Value.(Order).Execute()
		b.OrderList.Remove(i)
		i = nextOrder
	}
}
