package nullobject

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Nullobject test:", NullObjectPatternTest)
}

func NullObjectPatternTest(t *testing.T) {
	customerFactory := NewCustomerFactory()

	customer1 := customerFactory.GetCustomer("Lily")
	customer2 := customerFactory.GetCustomer("Allen")
	customer3 := customerFactory.GetCustomer("James")
	customer4 := customerFactory.GetCustomer("Bob")
	customer5 := customerFactory.GetCustomer("Jay")

	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
	fmt.Println(customer5.GetName())
}
