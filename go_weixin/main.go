package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type Context struct {
}

type Commentpent interface {
	Mount(c Commentpent, components ...Commentpent) error
	Remove(c Commentpent) error
	Do(ctx *Context) error
}

type BaseComponent struct {
	ChildComponents []Commentpent
}

func (bc *BaseComponent) Mount(c Commentpent, components ...Commentpent) (err error) {
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)
	return
}

func (bc *BaseComponent) Remove(c Commentpent) (err error) {
	panic("implement me")
	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(runFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

func (bc *BaseComponent) Do(ctx *Context) error {
	//panic("implement me")
	return nil
}

func (bc *BaseComponent) ChildsDo(ctx *Context) (err error) {
	for _, childCommponent := range bc.ChildComponents {
		if err := childCommponent.Do(ctx); err != nil {
			return err
		}
	}
	return
}

type CheckoutPageComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *CheckoutPageComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "订单结算页面组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

// AddressComponent 地址组件
type AddressComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *AddressComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "地址组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	checkoutPageComponent := &CheckoutPageComponent{}
	checkoutPageComponent.Mount(&AddressComponent{})
	checkoutPageComponent.Do(&Context{})

}

//https://mp.weixin.qq.com/s/QWfQGlKfwDlW2NkMaUCvJA
// 组合模式的抽象过程的核心是
// 按模块划分：业务逻辑归类,收敛的过程
// 父子关系（树）：把收敛之后的业务对象按父子关系绑定，依次被执行
// 与责任模式的 区别
// 责任链模式：链表
// 组合模式： 树
