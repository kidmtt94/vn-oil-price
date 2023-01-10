package main

import (
	"vn_oil_price/backgroundProcess"
	"vn_oil_price/server"
)

func main() {
	// Run backround task
	r := server.SetupRouter()
	backgroundProcess.RunBackgroundTask()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")

}
