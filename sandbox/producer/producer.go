package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/fellah/pprof/sandbox/chend"
)

func main() {
	chend := chend.NewChend()

	go func() {
		log.Fatalln(http.ListenAndServe(":8080", nil))
	}()

	go Tick()

	<-chend
	log.Println("end")
}

func Tick() {
	for {
		log.Print("TICK")
		time.Sleep(time.Second)
	}
}
