package main

import "./foo"

func main() {
	foo.MAX
	foo.internal_const

	foo.FooFunc(5)
	foo.internalFunc(5)
}
