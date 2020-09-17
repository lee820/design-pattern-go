package composite

import (
	"container/list"
	"reflect"
	"strconv"
)

//Employee 职员类
type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates *list.List
}

//NewEmployee 实例化职员类
func NewEmployee(name, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}

//Add 添加职员的下属
func (e *Employee) Add(emp Employee) {
	e.Subordinates.PushBack(emp)
}

//Remove 删除职员的下属
func (e *Employee) Remove(emp Employee) {
	for i := e.Subordinates.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
		}
	}
}

//GetSubordinates 获取职员下属列表
func (e *Employee) GetSubordinates() *list.List {
	return e.Subordinates
}

//ToString 获取职员的string信息
func (e *Employee) ToString() string {
	return "[ Name: " + e.Name + ", dept: " + e.Dept + ", Salary: " + strconv.Itoa(e.Salary) + " ]"
}
