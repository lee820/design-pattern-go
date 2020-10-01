package observer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Observer test:", ObserverTest)
}

func ObserverTest(t *testing.T) {
	subject := NewSubject()

	hexOB := NewHexObserver()
	DecOB := NewDecimalistmObserver()
	binayOB := NewBinaryObserver()

	hexOB.Attach(subject)
	DecOB.Attach(subject)
	binayOB.Attach(subject)

	fmt.Println("First state change: 8")
	subject.SetState(8)
	fmt.Println("First state change: 15")
	subject.SetState(15)
}
