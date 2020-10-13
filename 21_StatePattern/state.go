package state

import "fmt"

//Context 状态保存 类
type Context struct {
	state State
}

//NewContext 实例化状态保存类
func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

//SetState 设置状态保存类当前的状态
func (c *Context) SetState(s State) {
	c.state = s
}

//GetState 获取状态保存类当前的状态
func (c *Context) GetState() State {
	return c.state
}

//State 状态接口
type State interface {
	DoAction(context *Context)
	ToString() string
}

//StartState 开始状态类
type StartState struct{}

//NewStatartState 实例化开始状态类
func NewStatartState() *StartState {
	return &StartState{}
}

//DoAction 开始状态类的DoAction，实现State接口
func (start *StartState) DoAction(context *Context) {
	fmt.Println("Now is start state")
	context.state = start
}

//ToString 返回开始状态类名称
func (start *StartState) ToString() string {
	return "Start state"
}

//StopState 停止状态类
type StopState struct{}

//NewStopState 实例化停止状态类
func NewStopState() *StopState {
	return &StopState{}
}

//DoAction 停止状态类方法，实现State接口
func (stop *StopState) DoAction(context *Context) {
	fmt.Println("Now is stop state")
	context.state = stop
}

//ToString 返回停止状态类名称
func (stop *StopState) ToString() string {
	return "Stop state"
}
