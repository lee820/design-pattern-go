package iterator

//NameRepository 姓名仓库
type NameRepository struct {
	Names []string
}

//NewNameRepository 实例化姓名仓库
func NewNameRepository() *NameRepository {
	return &NameRepository{
		Names: make([]string, 0),
	}
}

//GetIterator 获取姓名仓库的迭代器
func (nr *NameRepository) GetIterator() func() (string, bool) {
	index := 0
	return func() (name string, ok bool) {
		if index >= len(nr.Names) {
			return
		}
		name, ok = nr.Names[index], true
		index++
		return
	}
}

//SetName 向姓名仓库添加名字
func (nr *NameRepository) SetName(name string) {
	nr.Names = append(nr.Names, name)
}
