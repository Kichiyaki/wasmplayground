package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)

	doc := js.Global().Get("document")

	a := newApp(doc)
	a.start()

	<-ch
}

type app struct {
	doc js.Value
}

func newApp(doc js.Value) *app {
	return &app{
		doc,
	}
}

func (a *app) start() {
	a.addEventListeners()
}

func (a *app) addEventListeners() {
	form := a.doc.Call("querySelector", "form")

	form.Call("addEventListener", "submit", a.add())
}

func (a *app) add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 || args[0].Type() != js.TypeObject {
			return nil
		}

		ev := args[0]
		ev.Call("preventDefault")

		aStr := ev.Get("target").Get("a").Get("value").String()
		bStr := ev.Get("target").Get("b").Get("value").String()

		aVal, err := strconv.ParseFloat(aStr, 10)
		if err != nil {
			a.showError(err.Error())
			return nil
		}

		bVal, err := strconv.ParseFloat(bStr, 10)
		if err != nil {
			a.showError(err.Error())
			return nil
		}

		a.showError("")
		a.showResult(aVal + bVal)

		return nil
	})
}

func (a *app) showResult(result float64) {
	p := a.doc.Call("querySelector", "#result")

	p.Set("innerHTML", result)
}

func (a *app) showError(msg string) {
	p := a.doc.Call("querySelector", "#error")

	p.Set("innerHTML", msg)
}
