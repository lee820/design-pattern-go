package filter

import (
	"container/list"
	"reflect"
)

//Person 人的模板
type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

//NewPerson 实例化一个人
func NewPerson(name, gender, marital string) *Person {
	return &Person{
		Name:          name,
		Gender:        gender,
		MaritalStatus: marital,
	}
}

//PersonFilter 人的过滤器接口
type PersonFilter interface {
	PersonFilter(persons *list.List) list.List
}

//MaleFilter 男性过滤器，筛选性别为male的人
type MaleFilter struct{}

//NewMaleFilter 实例化男性过滤器类
func NewMaleFilter() *MaleFilter {
	return &MaleFilter{}
}

//PersonFilter 实现PersonFilter接口
func (mf *MaleFilter) PersonFilter(persons *list.List) list.List {
	var malePersons list.List
	for i := persons.Front(); i != nil; i = i.Next() {
		if i.Value.(Person).Gender == "male" {
			malePersons.PushBack(i.Value)
		}
	}
	return malePersons
}

//FemaleFilter 女性过滤器类
type FemaleFilter struct{}

//NewFemaleFilter 实例化女性过滤器类
func NewFemaleFilter() *FemaleFilter {
	return &FemaleFilter{}
}

//PersonFilter 实现personFilter接口
func (ff *FemaleFilter) PersonFilter(persons *list.List) list.List {
	var femalePersons list.List
	for i := persons.Front(); i != nil; i = i.Next() {
		if i.Value.(Person).Gender == "female" {
			femalePersons.PushBack(i.Value)
		}
	}
	return femalePersons
}

//SingleFilter 单身过滤器类
type SingleFilter struct{}

//NewSingleFilter 实例化单身过滤器类
func NewSingleFilter() *SingleFilter {
	return &SingleFilter{}
}

//PersonFilter 实现PersonFilter过滤器接口
func (sf *SingleFilter) PersonFilter(persons *list.List) list.List {
	var singlePersons list.List
	for i := persons.Front(); i != nil; i = i.Next() {
		if i.Value.(Person).MaritalStatus == "single" {
			singlePersons.PushBack(i.Value)
		}
	}
	return singlePersons
}

//AndFilter 与过滤器类
type AndFilter struct {
	ThisFilter  PersonFilter
	OtherFilter PersonFilter
}

//NewAndFilter 实例化与过滤器
func NewAndFilter(thisFilter, otherFilter PersonFilter) *AndFilter {
	return &AndFilter{
		ThisFilter:  thisFilter,
		OtherFilter: otherFilter,
	}
}

//PersonFilter 实现PersonFilter接口
func (af *AndFilter) PersonFilter(persons *list.List) list.List {
	firstPersons := af.ThisFilter.PersonFilter(persons)
	return af.OtherFilter.PersonFilter(&firstPersons)
}

//OrFilter 或过滤器类
type OrFilter struct {
	ThisFilter  PersonFilter
	OtherFilter PersonFilter
}

//NewOrFilter 实例化或过滤器类
func NewOrFilter(thisFilter, otherFilter PersonFilter) *OrFilter {
	return &OrFilter{
		ThisFilter:  thisFilter,
		OtherFilter: otherFilter,
	}
}

//PersonFilter 实现PersonFilter接口
func (of *OrFilter) PersonFilter(persons *list.List) list.List {
	firstPersons := of.ThisFilter.PersonFilter(persons)
	otherPersons := of.OtherFilter.PersonFilter(persons)
	var allpersons list.List

	//找出既是firstPersons和otherPersons中所有人，去除重复的人。
	for i := otherPersons.Front(); i != nil; i = i.Next() {
		allpersons.PushBack(i.Value)
		for j := firstPersons.Front(); j != nil; j = j.Next() {
			if reflect.DeepEqual(i.Value, j.Value) {
				//do nothing
			} else {
				allpersons.PushBack(j.Value)
			}
		}
	}
	return firstPersons
}
