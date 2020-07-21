package Singleton

import "sync"

//传统方式实现单例模式
var (
	instance *Instance
	lock     sync.Mutex
)

type Instance struct {
	Name string
}

// 双重检查
func GetInstance(name string) *Instance {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil { //防止第一个go程创建实例解锁后，有之前排队阻塞go程，醒来再次创建实例
			instance = &Instance{Name: name}
		}
	}
	return instance
}
