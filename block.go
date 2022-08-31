package main

import (
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	MerKelRoot    []byte
	TimeStamp     int64
	Bits          int64
	Nonce         int64
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          targetBits,
		Nonce:         0,
		Data:          []byte(data)}

	// block.SetHash()
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})
}
