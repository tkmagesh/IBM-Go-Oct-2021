package main

func main() {
	helloWorld := fn()
	helloWorld()
}

func fn() func() {
	return func() {
		println("Hello World")
	}
}
