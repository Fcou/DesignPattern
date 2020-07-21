/*
建造者模式:将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
建造者模式可以将一个产品的内部表象与产品的生成过程分割开来，从而可以使一个建造过程生成具有
不同的内部表象的产品对象。如果我们用了建造者模式，那么用户就只需指定需要建造的类型就可以得到它们，
而具体建造的过程和细节就不需知道了。

*/

package main

import "fmt"

//Product 产品类，由多个零件组成
type Product struct {
	parts []string
}

//Add 添加新零件方法
func (p *Product) Add(part string) {
	p.parts = append(p.parts, part) //防止扩容，一定要返回新切片
}

//Show 展示该产品全部零件方法
func (p *Product) Show() {
	fmt.Println("该产品全部零件如下：")
	for _, v := range p.parts {
		fmt.Println(v)
	}
}

//Builder 抽象建设者类，确定产品由两个零件PartA和PartB组成，并有一个得到最终产品的方法GetResult()
type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() Product
}

//BuilderOne 具体建设者One
type BuilderOne struct {
	product Product
}

//BuildPartA 把零件A添加到产品中
func (b *BuilderOne) BuildPartA() {
	b.product.Add("零件A")
}

//BuildPartB 把零件B添加到产品中
func (b *BuilderOne) BuildPartB() {
	b.product.Add("零件B")
}

//GetResult 返回具体建设者One最终完成的产品
func (b *BuilderOne) GetResult() Product {
	return b.product
}

//BuilderTwo 具体建设者Two
type BuilderTwo struct {
	product Product
}

//BuildPartA 把零件X添加到产品中
func (b *BuilderTwo) BuildPartA() {
	b.product.Add("零件X")
}

//BuildPartB 把零件Y添加到产品中
func (b *BuilderTwo) BuildPartB() {
	b.product.Add("零件Y")
}

//GetResult 返回具体建设者Two最终完成的产品
func (b *BuilderTwo) GetResult() Product {
	return b.product
}

//Director 指挥者类
type Director struct {
}

//Construct 指挥建造流程,抽象出流程细节
func (d *Director) Construct(b Builder) {
	b.BuildPartA()
	b.BuildPartB()
}

func main() {
	director := new(Director)

	builderOne := new(BuilderOne)
	director.Construct(builderOne) //指挥者用builderOne的方法来构建产品
	p1 := builderOne.GetResult()
	p1.Show()

	builderTwo := new(BuilderTwo)
	director.Construct(builderTwo) //指挥者用builderTwo的方法来构建产品
	p2 := builderTwo.GetResult()
	p2.Show()
}
