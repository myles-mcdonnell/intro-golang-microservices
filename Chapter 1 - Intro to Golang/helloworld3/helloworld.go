package main

import (
	"fmt"
	"github.com/myles-mcdonnell/helloworld2/math"
	"github.com/spf13/viper"
)

const X_KEY  = "X"
const Y_KEY  = "Y"

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	x := viper.GetInt32(X_KEY)
	y := viper.GetInt32(Y_KEY)

	fmt.Printf("%v + %v = %v\r\n", x, y, math.Add(x, y))
}
