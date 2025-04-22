package main

import "fmt"

func main() {
	evento := []string{"event1", "event2", "event3", "event4"}
	evento = evento[:1]
	fmt.Println(evento)
}
