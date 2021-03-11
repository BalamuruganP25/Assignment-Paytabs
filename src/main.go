package main

import (
	logservice "Paytabs/src/Logservice"
	config "Paytabs/src/config"
	server "Paytabs/src/server"
	"fmt"
)

func main() {
	fmt.Println("Paytabs application started")
	config.ReadEnvironmentVariable()
	logservice.CreateLogMainFolder()
	logservice.CreateLogSubFolder()
	logservice.Applicationdetails()
	server.ApiServer()

}
