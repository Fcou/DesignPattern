/*
命令模式：将一个请求封装成一个对象，从而使你可用不同的请求对客户进行参数化。
对请求排队或者记录请求日志，以及支持可撤销操作。
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
