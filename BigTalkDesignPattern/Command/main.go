/*
命令模式：将一个请求封装成一个对象，从而使你可用不同的请求对客户进行参数化。
对请求排队或者记录请求日志，以及支持可撤销操作。

第一，它能较容易地设计一个命令队列；
第二，在需要的情况下，可以较容易地将命令记入日志；
第三，允许接收请求的一方决定是否要否决请求。
第四，可以容易地实现对请求的撤销和重做；
第五，由于加进新的具体命令类不影响其他的类，因此增加新的具体命令类很容易。其实还有最关键的优点就是命令模式把
请求一个操作的对象与知道怎么执行一个操作的对象分割开。

敏捷开发原则告诉我们，不要为代码添加基于猜测的、实际不需要的功能。如果不清楚一个系统是否需要命令模式，
一般就不要着急去实现它，事实上，在需要的时候通过重构实现这个模式并不困难，只有在真正需要如撤销/恢复操作等功能时，
把原来的代码重构为命令模式才有意义。

程杰. 大话设计模式
*/

package main

import (
	"fmt"
	"time"
)

//Barbecuer 烧烤师傅结构体
type Barbecuer struct {
}

// BakeBeef 烤牛肉方法
func (b *Barbecuer) BakeBeef() {
	fmt.Println("一份烤牛肉串")
}

// BakeMutton 烤羊肉方法
func (b *Barbecuer) BakeMutton() {
	fmt.Println("一份烤羊肉串")
}

//ICommand 抽象命令接口
type ICommand interface {
	setReceiver(b *Barbecuer)
	executeCommand()
}

// BakeBeefCommand 烤牛肉串命令
type BakeBeefCommand struct {
	receiver *Barbecuer
}

func (b *BakeBeefCommand) setReceiver(bar *Barbecuer) {
	b.receiver = bar
}
func (b *BakeBeefCommand) executeCommand() {
	b.receiver.BakeBeef()
}

// BakeMuttonCommand 烤羊肉串命令
type BakeMuttonCommand struct {
	receiver *Barbecuer
}

func (b *BakeMuttonCommand) setReceiver(bar *Barbecuer) {
	b.receiver = bar
}
func (b *BakeMuttonCommand) executeCommand() {
	b.receiver.BakeMutton()
}

//Waiter 记录用户需求，创建修改命令队列，最后通知命令执行者
type Waiter struct {
	orders map[string]ICommand
}

func (w *Waiter) initWaiter() {
	w.orders = make(map[string]ICommand)
}

func (w *Waiter) AddOrders(c ICommand) {
	w.orders[time.Now().String()] = c

}

func (w *Waiter) CancleOrders(time string) {
	delete(w.orders, time)
}

func (w *Waiter) Notify() {
	for _, v := range w.orders {
		v.executeCommand()
	}
}

func main() {
	barbercuer := new(Barbecuer)
	waiter := new(Waiter)
	waiter.initWaiter()

	mutton := new(BakeMuttonCommand)
	mutton.setReceiver(barbercuer)
	mutton2 := new(BakeMuttonCommand)
	mutton2.setReceiver(barbercuer)
	beef := new(BakeBeefCommand)
	beef.setReceiver(barbercuer)

	waiter.AddOrders(beef)
	waiter.AddOrders(mutton)
	waiter.AddOrders(mutton2)
	waiter.Notify()

}
