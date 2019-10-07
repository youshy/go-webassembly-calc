package main

import (
	"strconv"
	"syscall/js"
)

// iv stands for integer value
// wanted to make it nice and simple
func iv(i []js.Value, index int) int {
	val := js.Global().Get("document").Call("getElementById", i[index].String()).Get("value").String()
	vti, _ := strconv.Atoi(val)
	return vti
}

func add(this js.Value, i []js.Value) interface{} {
	println((iv(i, 0) + iv(i, 1)))
	return js.ValueOf(iv(i, 0) + iv(i, 1))
}

func subtract(this js.Value, i []js.Value) interface{} {
	println((iv(i, 0) - iv(i, 1)))
	return js.ValueOf(iv(i, 0) - iv(i, 1))
}

func multiply(this js.Value, i []js.Value) interface{} {
	println((iv(i, 0) * iv(i, 1)))
	return js.ValueOf(iv(i, 0) * iv(i, 1))
}

func divide(this js.Value, i []js.Value) interface{} {
	println((iv(i, 0) / iv(i, 1)))
	return js.ValueOf(iv(i, 0) / iv(i, 1))
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("multiply", js.FuncOf(multiply))
	js.Global().Set("divide", js.FuncOf(divide))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM IS ON")
	registerCallbacks()
	<-c
}
