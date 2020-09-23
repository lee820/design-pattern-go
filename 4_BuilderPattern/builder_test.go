package builder

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("测试Taco套餐", OrderTacoMeal)
	t.Run("测试牛肉汉堡套餐", OrderBeefBurger)
}

func OrderTacoMeal(t *testing.T) {
	meal := NewMeal()
	meal.AddItem(NewCocaCola())
	meal.AddItem(NewTaco())
	meal.AddItem(NewChips())
	meal.ShowItems()
	fmt.Printf("您共花费: %f元\n", meal.GetCost())
}

func OrderBeefBurger(t *testing.T) {
	meal := NewMeal()
	meal.AddItem(NewCocaCola())
	meal.AddItem(NewBeefBurger())
	meal.AddItem(NewBeefBurger())
	meal.ShowItems()
	fmt.Printf("您共花费: %f元\n", meal.GetCost())
}
