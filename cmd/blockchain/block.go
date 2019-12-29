package blockchain

//package main

//Block struct
type Block struct {
	Version             int
	TimeStamp           int64
	Payload             []byte
	HashOfThisBlock     []byte
	HashOfPreviousBlock []byte
}

//New block
func (b *Block) New(payload []byte) (Block, error) {

	var block Block

	block.Version = 1
	block.Payload = payload

	return block, nil
}
