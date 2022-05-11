package main

import (
	"counter/cron"
	"counter/router"
	"fmt"
	"github.com/fvbock/endless"
	"io/ioutil"
	"log"
	"syscall"
)

func main() {
	router := router.SetupRouter()
	cronTask := cron.InitCron()
	go cronTask.Start()
	defer cronTask.Stop()
	server := endless.NewServer(":8000", router)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Actual pid is %d", pid)
		ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
