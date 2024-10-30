package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Gateway")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan

	fmt.Println("Done")
}
