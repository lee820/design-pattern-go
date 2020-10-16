package strategy

//Strategy 策略类接口
type Strategy interface {
	DoOperation(num1, num2 int) int
}

//OperationAdd 加法策略类
type OperationAdd struct{}

//NewOperationAdd 实例化加法策略类
func NewOperationAdd() *OperationAdd {
	return &OperationAdd{}
}

//DoOperation 执行策略操作
func (add *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

//OperationSubtract 减法策略类
type OperationSubtract struct{}

//NewOperationSubtract 实例化减法策略类
func NewOperationSubtract() *OperationSubtract {
	return &OperationSubtract{}
}

//DoOperation 减法策略类执行的操作
func (sub *OperationSubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

//OperationMultiply 乘法策略类
type OperationMultiply struct{}

//NewOperationMultiply 实例化乘法策略类
func NewOperationMultiply() *OperationMultiply {
	return &OperationMultiply{}
}

//DoOperation 乘法策略类执行操作
func (multi *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1 * num2
}

//Context 策略的使用类
type Context struct {
	strategy Strategy
}

//NewContext 实例化策略使用类
func NewContext(str Strategy) *Context {
	return &Context{
		strategy: str,
	}
}

//ExecuteStrategy 执行当前策略
func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}
