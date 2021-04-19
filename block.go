type Block struct {
	Timestamp 			int64
	Data 				[]byte
	PrevBlockHash 		[]byte
	Hash 				[]byte
}

func (i *block) SetHash(){
	timestamp := []byte(strconv.FormatInt(i.Timestamp, 10))
	headers := bytes.Join([][]byte{i.PrevBlockHash, i.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	i.Hash = hash[:]
}

func newblock(data string, prevblockHash []byte) *block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block 
}