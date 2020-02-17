//https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
package main

import "fmt"

func main() {
	defaultThing := NewThing()
	fmt.Println(defaultThing.TellMe())

	customColourThing := NewThing(SetColour(Orange))
	fmt.Println(customColourThing.TellMe())

	totallyCustomThing := NewThing(SetColour(Pink), SetSize(100))
	fmt.Println(totallyCustomThing.TellMe())
}

type (
	Colour string

	config struct {
		colour Colour
		size int
	}

	thing struct {
		config *config
	}
)

const (
	Orange Colour = "Orange"
	Pink Colour = "Pink"
	White Colour = "White"
)

func SetColour(colour Colour) func(*config) {
	return func(c *config) {
		c.colour = colour
	}
}

func SetSize(size int) func(*config) {
	return func(c *config) {
		c.size = size
	}
}

func NewThing(opts ...func(*config)) *thing {
	conf := &config{
		colour: White,
		size:   10,
	}

	for _, opt := range opts {
		opt(conf)
	}

	return &thing{config:conf}
}

func (t *thing) TellMe() string {
	return fmt.Sprintf("I am %v and size %v", t.config.colour, t.config.size)
}

