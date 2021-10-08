package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
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
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		firstBlock := blockZero()
		spew.Dump(firstBlock)
		Blockchain = append(Blockchain, firstBlock)
	}()
	log.Fatal(run())
}

