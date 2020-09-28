package mediator

import "testing"

func Test(t *testing.T) {
	t.Run("Mediator test:", MediatorTest)
}

func MediatorTest(t *testing.T) {
	lee := NewUser("lee")
	wang := NewUser("wang")

	lee.SendMessage("hello")
	wang.SendMessage("world")
}
