package main

import (
	"flag"
	"fmt"
	"github.com/dayan-be/access-client/client"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//显示版本号信息　
	version := flag.Bool("v", false, "version")
	flag.Parse()

	if *version {
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	//2.log
	logrus.SetLevel(logrus.DebugLevel)

	//1.load configer
	cfg := client.Config()
	cfg.Load("config.yaml")

	// TODO: add one client to test

	cli := client.NewClient()
	go cli.Run()
	cli.Login(cfg.Accounts[0].PhoneNum, cfg.Accounts[0].Password)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	<-c
	os.Exit(0)
}