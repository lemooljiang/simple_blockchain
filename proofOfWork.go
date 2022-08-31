package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWork struct {
	block *Block

	target *big.Int
}

const targetBits = 6

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)

	target.Lsh(target, uint(256-targetBits))

	pow := ProofOfWork{block: block, target: target}
	return &pow
}

func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	block := pow.block
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(targetBits),
		IntToByte(nonce),
		block.Data}

	data := bytes.Join(tmp, []byte{})
	return data
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hash [32]byte
	var nonce int64 = 0
	var hashInt big.Int

	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("fond nonce,%d,%x", nonce, hash)
			break
		} else {
			nonce++

		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
