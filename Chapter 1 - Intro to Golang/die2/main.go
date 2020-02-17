package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type (
	dieThrow struct {
		die, value int
	}

	Die interface {
		Throw(chn chan dieThrow)
	}

	weightedDie struct {
		index, preferedValue int
	}

	die struct {
		index int
	}
)

func main() {

	intChan := make(chan dieThrow)

	go func() {
		for {
			dieThrow := <-intChan
			fmt.Printf("Die %v threw a %v\r\n", dieThrow.die, dieThrow.value)
		}
	}()

	dice := make([]Die, 2)
	dice[0] = &die{1}
	dice[1] = &weightedDie{2, 6}

	for _, die := range dice {
		go die.Throw(intChan)
	}

	for {
		bufio.NewScanner(os.Stdin).Scan()
		fmt.Println("Console Input")
	}
}

func (wd *weightedDie) Throw(chn chan dieThrow) {
	rand.Seed(time.Now().UnixNano())
	for {

		i := rand.Intn(12) + 1

		if i > 6 {
			i = wd.preferedValue
		}

		chn <- dieThrow{die: wd.index, value: i}
		time.Sleep(time.Second)
	}
}

func (d *die) Throw(chn chan dieThrow) {
	rand.Seed(time.Now().UnixNano())
	for {
		chn <- dieThrow{die: d.index, value: rand.Intn(6) + 1}
		time.Sleep(time.Second)
	}
}
