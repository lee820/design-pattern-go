package memento

//Memento 备忘录类
type Memento struct {
	state string
}

//NewMemento 实例化备忘录类
func NewMemento(st string) *Memento {
	return &Memento{
		state: st,
	}
}

//GetState 获取备忘录类的状态
func (m *Memento) GetState() string {
	return m.state
}

//Originator 初始类
type Originator struct {
	state string
}

//NewOriginator 实例化初始类
func NewOriginator(st string) *Originator {
	return &Originator{
		state: st,
	}
}

//SetState 初始化类设置状态
func (o *Originator) SetState(st string) {
	o.state = st
}

//GetState 从初始类中获取状态
func (o *Originator) GetState() string {
	return o.state
}

//SaveStateToMemento 将初始类状态保存到备忘录类
func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.state)
}

//GetStateFromMemento 将备忘录类的状态读取到初始类
func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
}

//CareTaker 保存类，用于存储备忘录实例
type CareTaker struct {
	MementoList map[int]*Memento
}

//NewCareTaker 实例化保存类
func NewCareTaker() *CareTaker {
	return &CareTaker{
		MementoList: make(map[int]*Memento),
	}
}

//Add 保存类添加备忘录实例
func (ct *CareTaker) Add(index int, memento *Memento) {
	ct.MementoList[index] = memento
}

//Get 保存类获取备忘录实例
func (ct *CareTaker) Get(index int) *Memento {
	return ct.MementoList[index]
}
