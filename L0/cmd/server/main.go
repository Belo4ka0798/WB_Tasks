package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"l0/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serv, err := server.NewServer("internal/config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	HandleInterrupt(serv)
	err = serv.Up()
	if err != nil {
		logrus.Fatal("SERVER DOWN!")
	}

}

func HandleInterrupt(s *server.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		s.Down()
		os.Exit(0)
	}()
}
