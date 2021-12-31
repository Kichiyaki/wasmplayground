package main

import (
	"log"
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)

	log.Println("Hello from Go!")
	js.Global().Set("add", add())

	<-ch
}

func add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var sum float64

		for _, arg := range args {
			if arg.Type() != js.TypeNumber || arg.IsNaN() {
				return 0
			}

			sum += arg.Float()
		}

		return sum
	})
}
