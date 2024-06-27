package file

import "fmt"

type BlockId struct {
	FileName string
	BlockNum int64
}

func NewBlockId(file_name string, block_num int64) BlockId {
	return BlockId{FileName: file_name, BlockNum: block_num}
}

func (blockId *BlockId) Equals(other BlockId) bool {
	return blockId.BlockNum == other.BlockNum && blockId.FileName == other.FileName
}

func (blockId *BlockId) String() string {
	return fmt.Sprintf("{file: %s, block: %d}", blockId.FileName, blockId.BlockNum)
}
