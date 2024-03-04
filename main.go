package main

import (
	"collection.com/config"
	"collection.com/router"
)

var logger *config.Logger

//TODO: checar se os métodos de básicos estão funcionando para item e linguagem
//TODO: criar o restante dos métodos para item e linguagem
//TODO: criar os outros tipos
//TODO: incorporar os outros tipos no item

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
