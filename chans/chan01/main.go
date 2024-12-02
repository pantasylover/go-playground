package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gamemann/go-playground/chans/chan01/reader"
	"github.com/gamemann/go-playground/chans/chan01/writer"
)

func main() {
	myChan := make(chan string)

	go reader.StartReader(myChan)
	go writer.StartWriter(myChan)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	fmt.Println("Cleaning up...")
}
