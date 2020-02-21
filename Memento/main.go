/*
备忘录模式：在不破坏封闭性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
这样以后就可将对象恢复到之前保存的状态。
关键将要保存的状态抽取出来形成一个新结构体保存，操作
*/

package main

import "fmt"

type role struct {
	vit int
	atk int
	def int
}

func (r *role) initState(vit, atk, def int) {
	r.vit = vit
	r.atk = atk
	r.def = def
}

func (r *role) showState() {
	fmt.Println(r.vit)
	fmt.Println(r.atk)
	fmt.Println(r.def)
}

func (r *role) saveState() *roleStateMemento {
	return &roleStateMemento{
		vit: r.vit,
		atk: r.atk,
		def: r.def,
	}
}

func (r *role) fight() {
	r.vit = 5
	r.atk = 2000
	r.def = 1
}
func (r *role) recoveryState(rsm roleStateMemento) {
	r.vit = rsm.getVit()
	r.atk = rsm.getAtk()
	r.def = rsm.getDef()
}

// roleStateMemento 角色状态备份，不暴露给外界
type roleStateMemento struct {
	vit int
	atk int
	def int
}

func (r *roleStateMemento) savState(vit, atk, def int) {
	r.vit = vit
	r.atk = atk
	r.def = def
}

func (r *roleStateMemento) getVit() int {
	return r.vit
}

func (r *roleStateMemento) getAtk() int {
	return r.atk
}

func (r *roleStateMemento) getDef() int {
	return r.def
}

type roleStateCaretaker struct {
	memento roleStateMemento
}

func (r *roleStateCaretaker) getMemento() roleStateMemento {
	return r.memento
}

func (r *roleStateCaretaker) setMemento(rsm roleStateMemento) {
	r.memento = rsm
}

func main() {
	//游戏刚开始
	hero := new(role)
	hero.initState(100, 3000, 50)
	hero.showState()
	//存档
	stateAdminer := new(roleStateCaretaker)
	stateAdminer.setMemento(*hero.saveState())
	//BOSS战
	hero.fight()
	hero.showState()
	//恢复之前状态
	hero.recoveryState(stateAdminer.getMemento())
	hero.showState()
}
