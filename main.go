package main

import (
	"azrb/cmd"
	"azrb/config"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
