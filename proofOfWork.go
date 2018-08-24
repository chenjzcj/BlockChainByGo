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

	//目标值
	target *big.Int
}

const targetBits = 24

func NewProofOfWork(block *Block) *ProofOfWork {

	//00000000000000000.....01
	target := big.NewInt(1)

	//0x000000100000000000000
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

	//1.拼装数据
	//2.哈希值转成big.Int类型

	var hash [32]byte
	var nonce int64 = 0
	var hashInt big.Int

	fmt.Println("Begin Mining...")
	fmt.Printf("target hash :    %x\n", pow.target.Bytes())

	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if hashInt.Cmp(pow.target) == -1 {
			//fmt.Printf("found nonce,nonce :%d,hash : %x\n",nonce,hash)
			fmt.Printf("found hash : %x, nonce : %d\n,", hash, nonce)
			break
		} else {
			//fmt.Printf("not found nonce,current nonce :%d,hash : %x\n", nonce, hash)
			nonce++
		}
	}
	return nonce, hash[:]

	/*for nonce{
		hash := sha256(block数据+nonce)
		if 转换(hash) < pow.target{
			找到了
		}else {
			nonce++
		}
	}
	return nonce,hash[:]*/
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
