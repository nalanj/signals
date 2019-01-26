package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	end := make(chan bool, 1)

	signal.Notify(sig)

	go func() {
		for {
			s := <-sig
			fmt.Println("got ", s)

			if s != syscall.SIGUSR1 && s != syscall.SIGUSR2 {
				end <- true
			}
		}
	}()

	// just wait forever until we get something over the channel
	<-end
}
