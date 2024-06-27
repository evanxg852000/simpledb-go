package query

import "fmt"

type RID struct {
	BlockNum int64
	Slot     int64
}

func NewRID(block_num int64, slot int64) RID {
	return RID{BlockNum: block_num, Slot: slot}
}

func (rId *RID) Equals(other RID) bool {
	return rId.BlockNum == other.BlockNum && rId.Slot == other.Slot
}

func (rId *RID) String() string {
	return fmt.Sprintf("{block: %d, slot: %d}", rId.BlockNum, rId.Slot)
}
