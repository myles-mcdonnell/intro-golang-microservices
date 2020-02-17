package main

import (
	"bufio"
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("An error occurred %v", err.(error).Error())
		}
	}()

	var i *bufio.Reader

	i.Buffered()
}
