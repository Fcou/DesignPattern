/*
抽象工厂模式：提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类。

main2.0版本，用简单工厂改进抽象工厂
*/

package main

import (
	"fmt"
)

//TokenName 名称接口
type TokenName interface {
	name()
}

// ETHTokenName 基于ETH发行的token，显示名称
type ETHTokenName struct {
}

func (e *ETHTokenName) name() {
	fmt.Println("ETH token name: EEE")
}

// RVNTokenName 基于RVN发行的token，显示名称
type RVNTokenName struct {
}

func (r *RVNTokenName) name() {
	fmt.Println("RVN token name: RRR")
}

//TokenNums 数量接口
type TokenNums interface {
	nums()
}

// ETHTokenNums 基于ETH发行的token，显示数量
type ETHTokenNums struct {
}

func (e *ETHTokenNums) nums() {
	fmt.Println("ETH token nums: 9999")
}

// RVNTokenNums 基于RVN发行的token，显示数量
type RVNTokenNums struct {
}

func (r *RVNTokenNums) nums() {
	fmt.Println("RVN token nums: 100000000")
}

//factory 简单工厂
type factory struct {
	factoryname string
}

//initFactory 选择设置工厂名称
func (e *factory) initFactory(name string) {
	e.factoryname = name
	//e.Factoryname = "ETH"
}

// CreateTokenA 根据工厂名称，选择创建名称
func (e *factory) CreateTokenName() TokenName {
	if e.factoryname == "ETH" {
		return &ETHTokenName{}
	} else if e.factoryname == "RVN" {
		return &RVNTokenName{}
	}
	return nil
}

// CreateTokenB 根据工厂名称，选择创建数量
func (e *factory) CreateTokenNums() TokenNums {
	if e.factoryname == "ETH" { //如果要增加另外的工厂名称，则要修改这里，增加类别
		return &ETHTokenNums{}
	} else if e.factoryname == "RVN" {
		return &RVNTokenNums{}
	}
	return nil
}

func main() {

	store := new(factory)
	store.initFactory("ETH") //选择设置工厂名称 “ETH”可以从配置中读取
	//store.initFactory("RVN")

	tName := store.CreateTokenName()
	tNums := store.CreateTokenNums()

	tName.name()
	tNums.nums()

	//call("ETH", "name", []string{"A"}) //

}
