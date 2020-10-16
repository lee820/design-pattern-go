package temp

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Template test:", TemplateTest)
}

func TemplateTest(t *testing.T) {
	football := NewFootBall()
	football.Play()
	fmt.Println("-------------------")
	basketball := NewBasketball()
	basketball.Play()
}
