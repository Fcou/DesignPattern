/*
中介者模式：用一个中介对象来封装一系列对象交互。
中介者使各对象不需要显式地引用，从而使其耦合松散，而且可以独立的改变他们之间的交互。

聊天室中，服务器作为中介者，转发各个客户端的消息给对应的接受者
*/

package main

import "fmt"

type UnitedNations interface {
	Declare(message string, colleague Country)
}

type UnitedNationsSecurityCouncil struct {
	usa  *USA
	iraq *Iraq
}

func (unsc *UnitedNationsSecurityCouncil) Declare(message string, colleague Country) {
	if colleague == unsc.usa {
		unsc.iraq.GetMessage(message)
	} else if colleague == unsc.iraq {
		unsc.usa.GetMessage(message)
	}
}

type Country interface {
	Declare(message string)
	GetMessage(message string)
}

type USA struct {
	UN UnitedNations
}

func (u *USA) Declare(message string) {
	u.UN.Declare(message, u)
}

func (u *USA) GetMessage(message string) {
	fmt.Println("USA 获得对方消息：", message)
}

type Iraq struct {
	UN UnitedNations
}

func (i *Iraq) Declare(message string) {
	i.UN.Declare(message, i)
}

func (i *Iraq) GetMessage(message string) {
	fmt.Println("Iraq 获得对方消息：", message)
}

func main() {
	UNSC := new(UnitedNationsSecurityCouncil)
	usa := new(USA)
	usa.UN = UNSC
	iraq := new(Iraq)
	iraq.UN = UNSC

	UNSC.usa = usa
	UNSC.iraq = iraq

	usa.Declare("你家有大规模杀伤性洗衣粉！")
	iraq.Declare("你家没洗衣粉吗?")
}
