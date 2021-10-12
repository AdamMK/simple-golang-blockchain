package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"io"
	"net/http"
)

func (s *Server) requests() {
	s.HandleFunc("/", s.showBlockchain()).Methods("GET")
	s.HandleFunc("/", s.addBlock()).Methods("POST")
}

func (s *Server) showBlockchain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(Blockchain, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(bytes))
	}
}

func (s *Server) addBlock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d Data

		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			respondWithJSON(w, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], d)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, d)
			return
		}
		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			newBlockchain := append(Blockchain, newBlock)
			longestChain(newBlockchain)
			spew.Dump(Blockchain)
		}

		respondWithJSON(w, http.StatusCreated, newBlock)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}