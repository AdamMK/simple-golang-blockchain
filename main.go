package goBlockchain

import (
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Block struct {
	Index 		int
	Timestamp 	string
	Price		int
	Operation	string
	Hash		string
	PreHash 	string
}

var Blockchain []Block

