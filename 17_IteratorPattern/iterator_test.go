package iterator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Iterator test:", GetIteratorTest)
}

func GetIteratorTest(t *testing.T) {
	nameRepository := NewNameRepository()
	nameRepository.SetName("Jay")
	nameRepository.SetName("JJ")
	nameRepository.SetName("Bob")
	nameRepository.SetName("Mary")

	it := nameRepository.GetIterator()
	for {
		name, ok := it()
		if !ok {
			break
		}
		fmt.Println("Get name: ", name)
	}
}
