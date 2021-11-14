package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func main() {
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	iLog := log.New(file, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
