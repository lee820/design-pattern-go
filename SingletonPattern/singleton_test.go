package singleton

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("懒汉模式:", Lazy)
	t.Run("饿汉模式:", Hungry)
	t.Run("懒汉加锁模式:", LazyWithMutex)
	t.Run("懒汉双重锁模式:", LazyWithDoubleMutex)
	t.Run("懒汉Once模式:", LazyWithOnce)
}

func Lazy(t *testing.T) {
	ins := GetInstance1()
	if ins == nil {
		fmt.Println("懒汉模式获取失败")
	}
	fmt.Println("懒汉模式获取成功")
}

func Hungry(t *testing.T) {
	ins := GetInstance2()
	if ins == nil {
		fmt.Println("饿汉模式获取失败")
	}
	fmt.Println("饿汉模式获取成功")
}

func LazyWithMutex(t *testing.T) {
	ins := GetInstance3()
	if ins == nil {
		fmt.Println("懒汉加锁模式获取失败")
	}
	fmt.Println("懒汉加锁模式获取成功")
}

func LazyWithDoubleMutex(t *testing.T) {
	ins := GetInstance4()
	if ins == nil {
		fmt.Println("懒汉双重锁模式获取失败")
	}
	fmt.Println("懒汉双重锁模式获取成功")
}

func LazyWithOnce(t *testing.T) {
	ins := GetInstance5()
	if ins == nil {
		fmt.Println("懒汉Once模式获取失败")
	}
	fmt.Println("懒汉Once模式获取成功")
}
