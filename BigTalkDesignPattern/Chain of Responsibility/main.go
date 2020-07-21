/*
责任链模式：使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。
将这个对象连成一个链，并沿着这条链传递该请求，直到有一个对象处理它为止。

以下代码来自：https://github.com/pibigstar/go-demo/tree/master/design/chain
感谢
*/

package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}

// 广告过滤
type AdHandler struct {
	handler Handler
}

func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
	ad.next(ad.handler, newContent)
}

func (ad *AdHandler) next(handler Handler, content string) {
	if ad.handler != nil {
		ad.handler.Handle(content)
	}
}

func (ad *AdHandler) setHandler(h Handler) {
	ad.handler = h
}

// 涉黄过滤
type YellowHandler struct {
	handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	fmt.Println(newContent)
	yellow.next(yellow.handler, newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if yellow.handler != nil {
		yellow.handler.Handle(content)
	}
}

func (yellow *YellowHandler) setHandler(h Handler) {
	yellow.handler = h
}

// 敏感词过滤
type SensitiveHandler struct {
	handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤。。。")
	newContent := strings.Replace(content, "敏感词", "***", 1)
	fmt.Println(newContent)
	sensitive.next(sensitive.handler, newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if sensitive.handler != nil {
		sensitive.handler.Handle(content)
	}
}

func (sensitive *SensitiveHandler) setHandler(h Handler) {
	sensitive.handler = h
}

func main() {
	adHandler := new(AdHandler)
	yellowHandler := new(YellowHandler)
	sensitiveHandler := new(SensitiveHandler)
	// 将责任链串起来
	adHandler.setHandler(yellowHandler)
	yellowHandler.setHandler(sensitiveHandler)

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}
