package main

const (
	Dev = bool(true)
)

func main() {
	if Dev {
		app := http()

		app.Listen("127.0.0.1:3000")
	} else {
		StartConsole()
	}
}
