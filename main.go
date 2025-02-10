package main

import (
	"github.com/Younes-khadraoui/azrb/cmd"
	"github.com/Younes-khadraoui/azrb/config"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
