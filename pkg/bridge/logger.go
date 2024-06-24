package bridge

import "syscall/js"

type IConsole interface {
	Log(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Dir(data interface{})
	DirXML(data interface{})
}

type JSConsole struct {
	console js.Value
}

func NewGlobalConsole() *JSConsole {
	console := js.Global().Get("console")
	return &JSConsole{console: console}
}

func (c *JSConsole) Log(args ...interface{}) {
	c.console.Call("log", args...)
}

func (c *JSConsole) Warn(args ...interface{}) {
	c.console.Call("warn", args...)
}

func (c *JSConsole) Error(args ...interface{}) {
	c.console.Call("error", args...)
}

func (c *JSConsole) Debug(args ...interface{}) {
	c.console.Call("debug", args...)
}

func (c *JSConsole) Info(args ...interface{}) {
	c.console.Call("info", args...)
}

func (c *JSConsole) Trace(args ...interface{}) {
	c.console.Call("trace", args...)
}

func (c *JSConsole) Dir(v interface{}) {
	c.console.Call("dir", v)
}

func (c *JSConsole) DirXML(v interface{}) {
	c.console.Call("dirxml", v)
}

var globalConsole IConsole

func Console() IConsole {
	if globalConsole != nil {
		return globalConsole
	}

	if js.Global().IsUndefined() {
		return mockLogger
	}

	globalConsole = NewGlobalConsole()
	return globalConsole
}
