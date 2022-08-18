package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

//定义区块结构
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

//定义区块链结构
type BlockChain struct {
	blocks []*Block
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          1,
		Nonce:         1,
		Data:          []byte(data)}
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}

	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})
}

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr("IntToByte", err)
	return buffer.Bytes()
}

func CheckErr(pos string, err error) {
	if err != nil {
		fmt.Println("error, pos:", pos, err)
		os.Exit(1)
	}
}

func NewBlockChain() *BlockChain {
	block := NewGenesisBlock()
	return &BlockChain{blocks: []*Block{block}}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	fmt.Println("hello world!8886")
	bc := NewBlockChain()
	bc.AddBlock("a send b 1 btc")
	bc.AddBlock("c send d 1 btc")
	for _, block := range bc.blocks {
		fmt.Printf("versin: %d\n", block.Version)
		fmt.Printf("timestamp: %d\n", block.TimeStamp)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("data: %s\n", block.Data)
	}
}
