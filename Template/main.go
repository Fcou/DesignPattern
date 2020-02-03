/*
模板方法模式：定义一个操作中的算法骨架，而将一些步骤延迟到子类中。
模板方法使得子类可以不改变一个算法的结构，即可重定义该算法的某些特定步骤。

例如：每份考卷，题目，选项都是一样的，但填写的每个人的个人姓名，选择的答案是变化的
以下代码来源：pibigstar
*/

package main

import "fmt"

type Cooker interface {
	open()
	fire()
	cooke()
	outfire()
	close()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) open() {
	fmt.Println("打开开关")
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

func (CookMenu) close() {
	fmt.Println("关闭开关")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outfire()
	cook.close()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}

func main() {
	xhs := new(XiHongShi)
	cjd := new(ChaoJiDan)

	doCook(xhs)
	doCook(cjd)
}
