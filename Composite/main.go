/*
组合模式：将对象组合成树形结构以表示“部分-整体”的层次结构。
组合模式使得用户对单个对象和组合对象的使用具有一致性。

坑1：interface，不能用指针方式传递。
坑2：interface不能拥有属性，但可以通过定义get\set方法，统一实现其接口的结构体的操作属性的方式。
坑3：记得初始化map。
*/

package main

import "fmt"

type Company interface {
	setName(name string)
	getName() string
	add(c Company)
	remove(c Company)
	display(depth int)
}

type ConcreateCompany struct {
	name     string
	children map[string]Company
}

func (cc *ConcreateCompany) setName(name string) {
	cc.children = make(map[string]Company)
	cc.name = name
}

func (cc *ConcreateCompany) getName() string {
	return cc.name
}

func (cc *ConcreateCompany) add(c Company) {
	cc.children[c.getName()] = c
}

func (cc *ConcreateCompany) remove(c Company) {
	delete(cc.children, c.getName())
}

func (cc *ConcreateCompany) display(depth int) {
	fmt.Println(manyStr(depth) + cc.name)
	for _, v := range cc.children {
		v.display(depth + 4)
	}
}

type HRDepartment struct {
	name string
}

func (hr *HRDepartment) setName(name string) {
	hr.name = name
}
func (hr *HRDepartment) getName() string {
	return hr.name
}
func (hr *HRDepartment) add(c Company) {
}
func (hr *HRDepartment) remove(c Company) {
}
func (hr *HRDepartment) display(depth int) {
	fmt.Println(manyStr(depth) + hr.name)
}

type FinanceDepartment struct {
	name string
}

func (f *FinanceDepartment) setName(name string) {
	f.name = name
}
func (f *FinanceDepartment) getName() string {
	return f.name
}
func (f *FinanceDepartment) add(c Company) {
}
func (f *FinanceDepartment) remove(c Company) {
}
func (f *FinanceDepartment) display(depth int) {
	fmt.Println(manyStr(depth) + f.name)
}

func manyStr(depth int) (s string) {
	for i := 0; i < depth; i++ {
		s = s + "-"
	}
	return
}

func main() {
	root := new(ConcreateCompany)
	root.setName("总公司")
	root.add(&HRDepartment{name: "总公司人力部门"})
	root.add(&FinanceDepartment{name: "总公司财务部门"})

	comp1 := new(ConcreateCompany)
	comp1.setName("北美县分公司")
	comp1.add(&HRDepartment{name: "北美县分公司人力部门"})
	comp1.add(&FinanceDepartment{name: "北美县分公司财务部门"})
	root.add(comp1)

	comp2 := new(ConcreateCompany)
	comp2.setName("北美加拿大村办事处")
	comp2.add(&HRDepartment{name: "北美加拿大村办事处人力部门"})
	comp2.add(&FinanceDepartment{name: "北美加拿大村办事处财务部门"})
	comp1.add(comp2)

	comp3 := new(ConcreateCompany)
	comp3.setName("欧县分公司")
	comp3.add(&HRDepartment{name: "欧县分公司人力部门"})
	comp3.add(&FinanceDepartment{name: "欧县分公司财务部门"})
	root.add(comp3)

	comp4 := new(ConcreateCompany)
	comp4.setName("欧县德村办事处")
	comp4.add(&HRDepartment{name: "欧县德村办事处人力部门"})
	comp4.add(&FinanceDepartment{name: "欧县德村办事处财务部门"})
	comp3.add(comp4)

	root.display(0)
}
