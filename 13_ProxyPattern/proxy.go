package proxy

import "fmt"

//Image 接口
type Image interface {
	Display()
}

//RealImage 原本的image类
type RealImage struct {
	FileName string
}

//NewRealImage 实例化RealImage
func NewRealImage(filename string) *RealImage {
	return &RealImage{
		FileName: filename,
	}
}

//Display RealImage实现Image接口的Display方法
func (ri *RealImage) Display() {
	fmt.Printf("Displaying %s.\n", ri.FileName)
}

//MyProxyImage 代理Image类
type MyProxyImage struct {
	Realimg  *RealImage
	FileName string
}

//NewMyProxyImage 实例化代理Image类
func NewMyProxyImage(filename string) *MyProxyImage {
	return &MyProxyImage{
		FileName: filename,
	}
}

//Display 实现Image接口函数
func (pi *MyProxyImage) Display() {
	if pi.Realimg == nil {
		pi.Realimg = NewRealImage(pi.FileName)
	}
	pi.Realimg.Display()
}
