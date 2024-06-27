package log

import "github.com/evanxg852000/simpledb/internal/file"

type LogIterator struct {
	fileManager    *file.FileManager
	blockId        file.BlockId
	currentPage    file.Page
	currentOffset  int64
	dataSpaceStart int64
}

func NewLogIterator(fm *file.FileManager, blockId file.BlockId) (LogIterator, error) {
	logIterator := LogIterator{
		fileManager:    fm,
		blockId:        blockId,
		currentPage:    file.NewPageWithData(make([]byte, fm.BlockSize())),
		currentOffset:  0,
		dataSpaceStart: 0,
	}

	err := logIterator.moveToBlock(blockId)
	if err != nil {
		return LogIterator{}, err
	}
	return logIterator, nil
}

func (logIter *LogIterator) HasNext() bool {
	return logIter.currentOffset < logIter.fileManager.BlockSize() || logIter.blockId.BlockNum > 0
}

func (logIter *LogIterator) Next() ([]byte, error) {
	if logIter.currentOffset == logIter.fileManager.BlockSize() {
		logIter.blockId = file.NewBlockId(logIter.blockId.FileName, logIter.blockId.BlockNum-1)
		logIter.moveToBlock(logIter.blockId)
	}

	data, err := logIter.currentPage.ReadBytes(logIter.currentOffset)
	if err != nil {
		return []byte{}, err
	}

	logIter.currentOffset = logIter.currentOffset + 8 + int64(len(data))
	return data, nil
}

// Moves a page of the file specified by blockId
// and positions the cursor at the first record in that block
func (logIter *LogIterator) moveToBlock(blockId file.BlockId) error {
	logIter.fileManager.Read(blockId, &logIter.currentPage)
	dataSpaceStart, err := logIter.currentPage.ReadInt(0)
	if err != nil {
		return err
	}
	logIter.dataSpaceStart = dataSpaceStart
	logIter.currentOffset = dataSpaceStart
	return nil
}
