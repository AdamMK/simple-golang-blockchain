package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func createHash(block Block) string {
	//data to be sha256Hash for each block
	record := strconv.Itoa(block.Index) +
		block.Timestamp +
		block.Data.Operation +
		fmt.Sprintf("%f", block.Data.Price) +
		block.PreHash
	//create a new hash
	h := sha256.New()
	h.Write([]byte(record))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}

func generateBlock(prevBlock Block, data Data) (Block, error) {

	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PreHash = prevBlock.Hash
	newBlock.Hash = createHash(newBlock)

	return newBlock, nil
}

func isBlockValid(currBlock, prevBlock Block) bool {

	if prevBlock.Index+1 != currBlock.Index ||
		prevBlock.Hash != currBlock.PreHash ||
		createHash(currBlock) != currBlock.Hash {
		return false
	}
	return true
}

func longestChain(newBlocks []Block)  {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
