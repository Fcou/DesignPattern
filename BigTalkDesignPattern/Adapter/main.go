/*
适配器模式：将一个类的接口转换为客户希望的另外一个接口。
Adapter模式，使得原本由于原本接口不兼容而不能一起工作的那些类可以一起工作。

已存在两个不同的接口，在不修改两个接口的前提下，新增一个适配器类，在两个类中统一接口
这在调用第三方库的过程中，非常常见，通常自己按照接口创建一个新struct，具体变量，具体方法，自己按照需求写好
最后创建接口，执行接口方法
*/

package main

import "fmt"

type Charge220V interface {
	on220vCharge()
}

type Charge110V interface {
	on110vCharge()
}

type iphoneCharger struct {
}

func (i *iphoneCharger) on110vCharge() {
	fmt.Println("iphone 充电中")
}

type transformer struct {
	charge Charge110V
}

func (t *transformer) on220vCharge() {
	t.charge.on110vCharge()
}

func main() {
	var c Charge220V
	c = &transformer{
		charge: &iphoneCharger{}, //赋值为我们需要的类
	}

	c.on220vCharge() //通过接口.方法 调用我们需要类的方法
}
