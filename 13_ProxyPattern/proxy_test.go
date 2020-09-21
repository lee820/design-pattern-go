package proxy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Proxy test:", ProxyTest)
}

func ProxyTest(t *testing.T) {
	pr := NewMyProxyImage("test.img")
	pr.Display() //需要从磁盘加载
	fmt.Println("--------------------")
	pr.Display() //不需要从磁盘加载，前面已经加载过了
}
