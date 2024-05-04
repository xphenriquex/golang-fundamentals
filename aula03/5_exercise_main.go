package main

import "fmt"

type Key interface {
	Press() interface{}
}

type Keyboard struct {
	LastKeyPressed interface{}
	keys           []Key
}

func (keyboard *Keyboard) PressAllKeys() {
	for _, v := range keyboard.keys {
		keyboard.LastKeyPressed = v.Press()
	}
}

type Letter struct {
}

func (l Letter) Press() interface{} {
	return "a"
}

type Number struct {
}

func (n Number) Press() interface{} {
	return 8
}

type KeyHelloWorld struct {
	key string
}

func (khw KeyHelloWorld) Press() interface{} {
	return khw.key
}

//constructor
func NewKeyHelloWorld() KeyHelloWorld {
	return KeyHelloWorld{"Hello World"}
	// return KeyHelloWorld{key: "Hello World"}
}

func exerciseMain() {
	// var letterALowercase = Letter{}
	// var number1 = Number{}

	var keys []Key
	// keys = append(keys, Letter{}, Number{}, NewKeyHelloWorld())
	keys = append(keys, Letter{})
	keys = append(keys, Number{})
	keys = append(keys, NewKeyHelloWorld())

	keyboard := Keyboard{keys: keys}
	keyboard.PressAllKeys()

	fmt.Println(keyboard.LastKeyPressed)

	switch v := keyboard.LastKeyPressed.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("default", v)
	}
}
