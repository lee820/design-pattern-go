package strategy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Strategy test:", StrategyTest)
}

func StrategyTest(t *testing.T) {
	context := NewContext(NewOperationAdd())
	fmt.Println("11 + 6 = ", context.ExecuteStrategy(11, 6))

	context = NewContext(NewOperationSubtract())
	fmt.Println("12 - 7 =", context.ExecuteStrategy(12, 7))

	context = NewContext(NewOperationMultiply())
	fmt.Println("12 * 11 =", context.ExecuteStrategy(12, 11))
}
