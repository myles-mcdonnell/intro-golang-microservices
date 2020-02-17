package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type dieThrow struct {
	die, value int
}

func main() {

	intChan := make(chan dieThrow)

	go func() {
		for {
			dieThrow := <-intChan
			fmt.Printf("Die %v threw a %v\r\n", dieThrow.die, dieThrow.value)
		}
	}()

	for i := 1;i<4;i++ {
		go throwDie(intChan, i)
	}

	for {
		bufio.NewScanner(os.Stdin).Scan()
		fmt.Println("Console Input")
	}
}

func throwDie( chn chan dieThrow, die int) {
	rand.Seed(time.Now().UnixNano())
	for {
		chn <- dieThrow{die:die,value: rand.Intn(6)+1}
		time.Sleep(time.Second)
	}
}


