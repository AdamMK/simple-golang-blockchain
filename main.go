package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      Data
	Hash      string
	PreHash   string
}

type Data struct {
	Operation string
	Price     float32
}

var Blockchain []Block


func blockZero() Block {

	t := time.Now()
	initData := Data{"", 0}
	b := Block{0, t.String(), initData, "", ""}
	return b
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	serv := NewServer()
	httpAddr := ":" + os.Getenv("PORT")

	firstBlock := blockZero()
	Blockchain = append(Blockchain, firstBlock)

	quitSig := make(chan os.Signal, 1)
	signal.Notify(quitSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(httpAddr, serv); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()
	log.Println("Listening on port", os.Getenv("PORT"))

	<-quitSig
	log.Print("Server Stopped")
}

