package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Hello World.")

	hostname, err := os.Hostname()
	if err!=nil {
		fmt.Printf("an error occured when trying to get the hostname %v", err.Error())
	} else {
		fmt.Printf("I'm running on host %v, which has %v cores!",  hostname, runtime.NumCPU())
	}
}
