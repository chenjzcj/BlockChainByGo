package main

import (
	"time"
)

type Block struct {
	//版本
	Version int64
	//前区块的哈希值
	PrevBlockHash []byte
	//当前区块的哈希值,为了简化代码
	Hash []byte
	//梅克尔根
	MerKelRoot []byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64

	//交易信息
	Data []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash TODO
		MerKelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       targetBits,
		Nonce:      0,
		Data:       []byte(data)}

	//block.SetHash()
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

/*func (block *Block) SetHash() {
	//func Sum256(data []byte) [Size]byte {

	//func Join(s [][]byte, sep []byte) []byte {

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
}*/

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
