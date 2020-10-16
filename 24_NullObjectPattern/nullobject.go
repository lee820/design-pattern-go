package nullobject

//AbstractCustomer 客户对象接口
type AbstractCustomer interface {
	Isnil() bool
	GetName() string
}

//RealCustomer 真实客户对象类
type RealCustomer struct {
	Name string
}

//NewRealCustomer 实例化真实客户对象
func NewRealCustomer(name string) *RealCustomer {
	return &RealCustomer{
		Name: name,
	}
}

//Isnil 真实客户对象nil判断
func (r *RealCustomer) Isnil() bool {
	return false
}

//GetName 获取真实客户对象名称
func (r *RealCustomer) GetName() (name string) {
	return r.Name
}

//NullCustomer 空客户对象类
type NullCustomer struct {
	Name string
}

//NewNullCustomer 实例化空客户对象
func NewNullCustomer() *NullCustomer {
	return &NullCustomer{}
}

//Isnil 空客户对象nil判断
func (n *NullCustomer) Isnil() bool {
	return true
}

//GetName 空客户对象获取名称
func (n *NullCustomer) GetName() (name string) {
	return "Not Available in Customer Database"
}

//CustomerFactory 客户工厂类
type CustomerFactory struct {
	Names []string
}

//NewCustomerFactory 实例化客户工厂类
func NewCustomerFactory() *CustomerFactory {
	return &CustomerFactory{
		Names: []string{"Bob", "Lily", "James"},
	}
}

//GetCustomer 从客户工厂类获取客户对象
func (c *CustomerFactory) GetCustomer(name string) AbstractCustomer {
	for _, v := range c.Names {
		if v == name {
			return NewRealCustomer(name)
		}
	}
	return NewNullCustomer()
}
