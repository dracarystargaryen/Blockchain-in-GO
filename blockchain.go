type Blockchain struct {
	blocks []*Block

}

func (bc *Blockchain) AddBlock(data string){
	prevBlock := bl.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	bc.blocks = append(bl.blocks, NewBlock)
}

func NewGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}

func newBlockchain() *Blockchain{
	return &Blockchain{[]*block{NewGenesisBlock()}}
}

func main() {
	bl := newBlockchain()
	bl.addBlock("Send 1 Bitcoin to Nick")
	bl.addBlock("Send 2 Bitcoin to Nick")

	for _, block := range bl.blocks {
		fmt.Printf("Prev. hash: %x\n", block.prevBlockHash)
		fmt.Printf("Data: %s", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}