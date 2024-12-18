package gadget

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValues() string
}

type MyType int

func (t MyType) MethodWithoutParameters() {
	fmt.Println("Method without parameters")
}
func (t MyType) MethodWithParameters(param float64) {
	fmt.Println("Method with parameters called with", param)
}
func (t MyType) MethodWithReturnValues() string {
	return "Hi, from MethodWithReturnValues"
}
func (t MyType) MethodNotInterface() {
	fmt.Println("Method not interface")
}
