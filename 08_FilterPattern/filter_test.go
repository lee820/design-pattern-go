package filter

import (
	"container/list"
	"fmt"
	"testing"
)

var allPersons *list.List

func Test(t *testing.T) {
	t.Run("AddPersons", AddPersons)
	t.Run("GetAllMales: ", GetAllMales)
	t.Run("GetAllFemales: ", GetAllFemales)
	t.Run("GetAllSingles: ", GetAllSingles)
	t.Run("GetSingleMales: ", AndFilterTest)
	t.Run("GetSingleOrFemales: ", OrFilterTest)
}

func PrintPersons(persons *list.List) {
	for i := persons.Front(); i != nil; i = i.Next() {
		fmt.Printf("Person: [ Name: %s, Gender: %s, Marital Status: %s ] \n",
			i.Value.(Person).Name,
			i.Value.(Person).Gender,
			i.Value.(Person).MaritalStatus)
	}
}

func AddPersons(t *testing.T) {
	allPersons = list.New()
	lee := NewPerson("lee", "male", "married")
	lu := NewPerson("lv", "male", "single")
	yang := NewPerson("yang", "female", "married")
	zhu := NewPerson("zhu", "female", "single")

	allPersons.PushBack(*lee)
	allPersons.PushBack(*lu)
	allPersons.PushBack(*yang)
	allPersons.PushBack(*zhu)
	PrintPersons(allPersons)
}

func GetAllMales(t *testing.T) {
	maleFilter := NewMaleFilter()
	allmales := maleFilter.PersonFilter(allPersons)
	PrintPersons(&allmales)
}

func GetAllFemales(t *testing.T) {
	femaleFilter := NewFemaleFilter()
	allFemales := femaleFilter.PersonFilter(allPersons)
	PrintPersons(&allFemales)
}

func GetAllSingles(t *testing.T) {
	singleFilter := NewSingleFilter()
	allSingles := singleFilter.PersonFilter(allPersons)
	PrintPersons(&allSingles)
}

func AndFilterTest(t *testing.T) {
	maleFilter := NewMaleFilter()
	singleFilter := NewSingleFilter()
	andFilter := NewAndFilter(maleFilter, singleFilter)
	persons := andFilter.PersonFilter(allPersons)
	PrintPersons(&persons)
}

func OrFilterTest(t *testing.T) {
	femaleFilter := NewFemaleFilter()
	singleFilter := NewSingleFilter()
	orFilter := NewOrFilter(femaleFilter, singleFilter)
	persons := orFilter.PersonFilter(allPersons)
	PrintPersons(&persons)
}
