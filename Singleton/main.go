/*
单例模式：保证一个类只有一个实例，并提供一个访问它的全局访问点。
通常我们可以让一个全局变量使得一个对象被访问，但它不能防止你实例化多个对象。
一个最好的办法就是，让类自身负责保存它的唯一实例。这个类可以保证没有其他实例可以被创建，
并且它可以提供一个访问该实例的方法。
单例模式因为Singleton类封装它的唯一实例，这样它可以严格地控制客户怎样访问它以及何时访问它。
简单地说就是对唯一实例的受控访问。

程杰. 大话设计模式
以下程序来自：pibigstar，第一次用sync.Once
*/

package main

import (
	"fmt"
	"sync"
)

var (
	goInstance *Instance
	once       sync.Once
)

type Instance struct {
	Name string
	Age  int
}

// 使用go 实现单例模式
func (i *Instance) getInstance(name string, age int) {
	if goInstance == nil {
		once.Do(func() { //仅调用一次匿名函数
			goInstance = &Instance{
				Name: name,
				Age:  age,
			}
		})
	}
	return
}

func main() {
	goInstance.getInstance("Fcou", 40)
	goInstance.getInstance("Jack", 50)
	fmt.Println(*goInstance)
}
