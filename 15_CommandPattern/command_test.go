package command

import "testing"

func Test(t *testing.T) {
	t.Run("Order test:", OrderTest)
}

func OrderTest(t *testing.T) {
	maotaiStock := NewStock()

	buyStockOrder := NewBuyStock(*maotaiStock)
	sellStockOrder := NewSellStock(*maotaiStock)

	broker := NewBroker()
	broker.TakeOrder(buyStockOrder)
	broker.TakeOrder(sellStockOrder)

	broker.PlaceOrders()
}
