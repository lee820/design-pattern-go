package composite

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("addEmployee: ", AllEmployee)
}

func AllEmployee(t *testing.T) {
	ceo := NewEmployee("lee", "ceo", 9999)

	pm := NewEmployee("wang", "technology", 1000)
	programmer1 := NewEmployee("jolin", "technology", 8000)
	programmer2 := NewEmployee("jay", "technology", 8000)

	minister := NewEmployee("zhou", "accounting", 5000)
	finance1 := NewEmployee("zzz", "accounting", 3000)
	finance2 := NewEmployee("Mary", "accounting", 2900)

	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	//打印所有职员
	fmt.Println(ceo.ToString())
	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
	}
}
