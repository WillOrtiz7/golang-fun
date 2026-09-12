package main

import (
	"awesomeProject/utils"
)

func main() {
	utils.InitEnv()
	utils.InitDatabase()
	utils.InitHttpServer()
}
