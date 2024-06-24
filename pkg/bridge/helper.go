package bridge

import "syscall/js"

func Falsey(input js.Value) bool {
	return !input.Truthy()
}

func IsString(input js.Value) bool {
	return input.Type() == js.TypeString
}

func IsNullish(input js.Value) bool {
	return input.Type() == js.TypeNull || input.Type() == js.TypeUndefined
}
