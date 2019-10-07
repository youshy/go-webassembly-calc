package main

import (
	"math"
	"strconv"
	"syscall/js"
)

// iv stands for integer value
// wanted to make it nice and simple
func iv(i []js.Value, index int) float64 {
	val := js.Global().Get("document").Call("getElementById", i[index].String()).Get("value").String()
	vti, _ := strconv.ParseFloat(val, 64)
	return math.Round(vti*100) / 100
}

func setResult(i float64) {
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", i)
}

func add(this js.Value, i []js.Value) interface{} {
	setResult(iv(i, 0) + iv(i, 1))
	return js.ValueOf(iv(i, 0) + iv(i, 1))
}

func subtract(this js.Value, i []js.Value) interface{} {
	setResult((iv(i, 0) - iv(i, 1)))
	return js.ValueOf(iv(i, 0) - iv(i, 1))
}

func multiply(this js.Value, i []js.Value) interface{} {
	setResult((iv(i, 0) * iv(i, 1)))
	return js.ValueOf(iv(i, 0) * iv(i, 1))
}

func divide(this js.Value, i []js.Value) interface{} {
	setResult((iv(i, 0) / iv(i, 1)))
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
