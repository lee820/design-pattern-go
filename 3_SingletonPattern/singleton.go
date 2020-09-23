package singleton

import "sync"

//Singleton 单例模式类
type Singleton struct{}

//Instance1 单例的实例1 懒汉模式获取
var Instance1 *Singleton

//GetInstance1 获取实例（懒汉模式）
func GetInstance1() *Singleton {
	if Instance1 == nil {
		Instance1 = &Singleton{}
	}
	return Instance1
}

//Instance2 实例2 饿汉模式获取
var Instance2 *Singleton = &Singleton{}

//GetInstance2 获取实例2（饿汉模式）
func GetInstance2() *Singleton {
	return Instance2
}

//Instance3 带锁的懒汉模式
var Instance3 *Singleton
var mut sync.Mutex

//GetInstance3 获取实例(带锁的懒汉模式)
func GetInstance3() *Singleton {
	mut.Lock()
	defer mut.Unlock()
	if Instance3 == nil {
		Instance3 = &Singleton{}
	}
	return Instance3
}

//Instance4 双重锁懒汉模式
var Instance4 *Singleton
var mut2 sync.Mutex

//GetInstance4 懒汉模式双重锁获取实例
func GetInstance4() *Singleton {
	if Instance4 == nil {
		mut2.Lock()
		defer mut2.Unlock()
		if Instance4 == nil {
			Instance4 = &Singleton{}
		}
	}
	return Instance4
}

//Instance5 sync.Once实现
var Instance5 *Singleton
var once sync.Once

//GetInstance5 sync.Once 获取实例，sync.Once只会执行一次
func GetInstance5() *Singleton {
	once.Do(func() {
		Instance5 = &Singleton{}
	})
	return Instance5
}
