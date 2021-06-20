package main

const (
	GUI = false
)

func main() {
	if GUI {
		StartGui()
	} else {
		StartConsole()
	}
}
