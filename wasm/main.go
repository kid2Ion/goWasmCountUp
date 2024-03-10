package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	setUpButton()
	<-c
}

func setUpButton() {
	document := js.Global().Get("document")
	button := document.Call("getElementById", "countUp")

	button.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		countUp()
		return nil
	}))
}

var count int

func countUp() {
	count++
	fmt.Println("current count:", count)
	updateDisplay(count)
}

func updateDisplay(count int) {
	document := js.Global().Get("document")
	countDisplay := document.Call("getElementById", "countDisplay")
	countDisplay.Set("innerText", count)
}
