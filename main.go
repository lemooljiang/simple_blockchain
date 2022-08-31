package main

import "fmt"

func main() {
	fmt.Println("hello world!6833")
	bc := NewBlockChain()
	bc.AddBlock("a send b 1.666 btc")
	bc.AddBlock("c send d 1 btc")
	for _, block := range bc.blocks {
		fmt.Printf("versin: %d\n", block.Version)
		fmt.Printf("timestamp: %d\n", block.TimeStamp)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
