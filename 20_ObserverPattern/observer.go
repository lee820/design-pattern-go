package observer

import (
	"container/list"
	"fmt"
)

//Subject 观察目标类
type Subject struct {
	state    int
	observer *list.List
}

//NewSubject 实例化观察目标
func NewSubject() *Subject {
	return &Subject{
		state:    0,
		observer: list.New(),
	}
}

//GetState 获取观察目标状态值
func (s *Subject) GetState() int {
	return s.state
}

//SetState 设置观察目标状态值
func (s *Subject) SetState(st int) {
	s.state = st
	//通知观察者
	s.NotifyAllObserver()
}

//Attach 将观察者绑定到自己
func (s *Subject) Attach(ob Observer) {
	s.observer.PushBack(ob)
}

//NotifyAllObserver 通知所有观察值
func (s *Subject) NotifyAllObserver() {
	for i := s.observer.Front(); i != nil; i = i.Next() {
		i.Value.(Observer).Update()
	}
}

//Observer 观察者抽象接口
type Observer interface {
	Update()
}

//BinaryObserver 二进制观察者类
type BinaryObserver struct {
	subject *Subject
}

//NewBinaryObserver 实例化二进制观察者类
func NewBinaryObserver() *BinaryObserver {
	return &BinaryObserver{}
}

//Attach 将二进制观察者绑定到观察目标
func (bo *BinaryObserver) Attach(sub *Subject) {
	bo.subject = sub
	bo.subject.Attach(bo) //观察目标将本观察者加入更新链
}

//Update 二进制观察者接受被观察目标的更新信息
func (bo *BinaryObserver) Update() {
	fmt.Printf("Binary String: %b \n", bo.subject.GetState())
}

//DecimalismObserver 十进制观察者类
type DecimalismObserver struct {
	subject *Subject
}

//NewDecimalistmObserver 实例化十进制观察者
func NewDecimalistmObserver() *DecimalismObserver {
	return &DecimalismObserver{}
}

//Attach 将十进制观察者绑定到观察目标
func (do *DecimalismObserver) Attach(sub *Subject) {
	do.subject = sub
	do.subject.Attach(do)
}

//Update 十进制观察者接收观察目标的更新
func (do *DecimalismObserver) Update() {
	fmt.Printf("Decimalism string: %d\n", do.subject.GetState())
}

//HexObserver 十六进制观察者类
type HexObserver struct {
	subject *Subject
}

//NewHexObserver 实例化十六进制观察者类
func NewHexObserver() *HexObserver {
	return &HexObserver{}
}

//Attach 将十六进制观察者绑定到观察目标
func (ho *HexObserver) Attach(sub *Subject) {
	ho.subject = sub
	ho.subject.Attach(ho)
}

//Update 十六进制观察者接收观察目标的更新信息。
func (ho *HexObserver) Update() {
	fmt.Printf("Hex string: %X\n", ho.subject.GetState())
}
