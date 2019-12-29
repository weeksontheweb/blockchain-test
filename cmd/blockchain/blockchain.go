package blockchain

//BlockChain struct
type BlockChain []Block

//Add block to Blockchain
func (bc BlockChain) Add(blockchain []Block, block Block) ([]Block, error) {

	blockchain = append(blockchain, block)
	return blockchain, nil
}

//AddGenesisBlock to Blockchain
func (bc BlockChain) AddGenesisBlock(blockchain []Block) ([]Block, error) {

	return blockchain, nil
}
