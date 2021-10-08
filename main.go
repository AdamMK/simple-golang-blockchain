package goBlockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data		Data
	Hash      string
	PreHash   string
}

type Data struct {
	Operation string
	Price     float32
}

var Blockchain []Block

func createHash(block Block) string {
	//data to be sha256Hash for each block
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data.Operation + fmt.Sprintf("%f", block.Data.Price) + block.PreHash
	//create a new hash
	h := sha256.New()
	h.Write([]byte(record))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}

func generateBlock(prevBlock Block, data Data) (Block, error)  {

	var newBlock Block

	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data.Operation = data.Operation
	newBlock.Data.Price = data.Price
	newBlock.PreHash = prevBlock.Hash
	newBlock.Hash = createHash(newBlock)

	return newBlock, nil
}


