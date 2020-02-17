/*
抽象工厂模式：提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类。

读取文件配置+反射+抽象工厂模式
*/

package main

import (
	"fmt"
)

//Token 代币
type Token interface {
	name()
}

// ETHTokenA 基于ETH发行的token，代币名称为A
type ETHTokenA struct {
}

func (e *ETHTokenA) name() {
	fmt.Println("ETH token name: A")
}

// ETHTokenB 基于ETH发行的token,代币名称为B
type ETHTokenB struct {
}

func (e *ETHTokenB) name() {
	fmt.Println("ETH token name: B")
}

// RVNTokenA 基于ERVN发行的token，代币名称为A
type RVNTokenA struct {
}

func (r *RVNTokenA) name() {
	fmt.Println("RVN token name: A")
}

// RVNTokenB 基于ERVN发行的token，代币名称为B
type RVNTokenB struct {
}

func (r *RVNTokenB) name() {
	fmt.Println("RVN token name: B")
}

//Factory 工厂接口
type Factory interface {
	CreateToken(name string) Token
}

// ETHFactory ETH工厂
type ETHFactory struct {
}

// CreateToken ETH工厂根据需求，创建代币
func (e *ETHFactory) CreateToken(name string) Token {
	if name == "A" {
		return &ETHTokenA{}
	} else if name == "B" {
		return &ETHTokenB{}
	}
	return nil
}

// RVNFactory RVN工厂
type RVNFactory struct {
}

// CreateToken RVN工厂根据需求，创建代币
func (r *RVNFactory) CreateToken(name string) Token {
	if name == "A" {
		return &RVNTokenA{}
	} else if name == "B" {
		return &RVNTokenB{}
	}
	return nil
}

// TokenFactoryStore 代币工厂提供者
type TokenFactoryStore struct {
	factory Factory
}

func (t *TokenFactoryStore) createToken(name string) Token {
	return t.factory.CreateToken(name)
}

func main() {

	store := new(TokenFactoryStore)
	// 提供ETH工厂
	store.factory = new(ETHFactory)
	EthToken := store.createToken("A")
	EthToken.name()

	// 提供RVN工厂
	store.factory = new(RVNFactory)
	RvnToken := store.createToken("B")
	RvnToken.name()

}
