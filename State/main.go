/*
状态模式：当一个对象的内在状态发生改变时，允许改变其行为，这个对象看起来像是改变了其类。
状态模式主要解决的是当控制一个对象状态转换的条件表达式过于复杂时的情况。
把状态的判断逻辑转移到表示不同状态的一系列类当中，可以把复杂的判断逻辑简化。
当然，如果这个状态判断很简单，那就没必要用‘状态模式’了。


*/
package main

import "fmt"

type state interface {
	runSpeed(m marathon)
}

type firstStageState struct {
}

func (f *firstStageState) runSpeed(m marathon) {
	if m.distance <= 10 {
		fmt.Println("第一阶段：跑步配速：8’00‘/公里")
	} else {
		m.setState(new(secondStageState))
		m.setRunSpeed()
	}
}

type secondStageState struct {
}

func (f *secondStageState) runSpeed(m marathon) {
	if m.distance <= 35 {
		fmt.Println("第二阶段：跑步配速：7’30‘/公里")
	} else {
		m.setState(new(SprintStageState))
		m.setRunSpeed()
	}
}

type SprintStageState struct {
}

func (f *SprintStageState) runSpeed(m marathon) {
	if m.distance <= 42.185 {
		fmt.Println("冲刺阶段：跑步配速：7’00‘/公里")
	}
}

type marathon struct {
	distance     float32
	currentState state
}

func (m *marathon) initState() {
	m.distance = 0
	m.currentState = new(firstStageState)
}

func (m *marathon) setState(s state) {
	m.currentState = s
}

func (m *marathon) setRunSpeed() {
	m.currentState.runSpeed(*m)
}

func main() {
	mar := new(marathon)
	mar.initState()

	mar.distance = 8
	mar.setRunSpeed()
	mar.distance = 22
	mar.setRunSpeed()
	mar.distance = 41
	mar.setRunSpeed()
}