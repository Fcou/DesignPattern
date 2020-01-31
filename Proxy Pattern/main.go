/*
代理模式：为其他对象提供一种代理，以控制对这个对象的访问
不暴露代理对象，对代理对象进行了一层封装，操作时直接操作代理对象即可
以下示例来源：https://github.com/pibigstar/go-demo/blob/master/design/proxy/proxy.go
*/

package main

import "fmt"

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}

}

// 火车代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
	} else {
		fmt.Println("票已售空")
	}
}
