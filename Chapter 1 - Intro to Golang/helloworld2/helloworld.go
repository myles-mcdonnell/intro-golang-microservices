package main

import (
	"fmt"
	"os"
	"runtime"
	"github.com/spf13/viper"
)

const NAME_ENV_KEY  = "NAME"

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetDefault(NAME_ENV_KEY, "Joe Bloggs")

	fmt.Printf("Hello %v\r\n", viper.GetString(NAME_ENV_KEY))

	hostname, err := os.Hostname()
	if err!=nil {
		fmt.Printf("an error occured when trying to get the hostname %v", err.Error())
	} else {
		fmt.Printf("I'm running on host %v, which has %v cores!",  hostname, runtime.NumCPU())
	}
}
