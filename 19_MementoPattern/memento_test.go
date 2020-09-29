package memento

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Memento test:", MementoTest)
}

func MementoTest(t *testing.T) {
	careTaker := NewCareTaker()
	originator := NewOriginator("State #1")
	originator.SetState("State #2")
	careTaker.Add(1, originator.SaveStateToMemento())
	originator.SetState("State #3")
	careTaker.Add(2, originator.SaveStateToMemento())
	originator.SetState("State #4")

	fmt.Println("Current state:", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("First save state:", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(2))
	fmt.Println("Second save state:", originator.GetState())
}
