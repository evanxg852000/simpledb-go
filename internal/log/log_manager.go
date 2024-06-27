package log

import (
	"sync"

	"github.com/evanxg852000/simpledb/internal/file"
)

type LogManager struct {
	logFile      string
	fileManager  *file.FileManager
	logPage      file.Page
	currentBlock file.BlockId
	latestLSN    int64
	lastSavedLSN int64
	mu           *sync.Mutex
}

func NewLogManager(fm *file.FileManager, logFile string) (*LogManager, error) {
	logManager := &LogManager{
		logFile:      logFile,
		fileManager:  fm,
		logPage:      file.NewPageWithData(make([]byte, fm.BlockSize())),
		currentBlock: file.BlockId{},
		latestLSN:    0,
		lastSavedLSN: 0,
		mu:           new(sync.Mutex),
	}

	numBlock, err := fm.BlockCount(logFile)
	if err != nil {
		return logManager, err
	}

	if numBlock == 0 {
		logManager.currentBlock, err = logManager.appendNewBlock()
		if err != nil {
			return logManager, err
		}
		return logManager, nil
	}

	logManager.currentBlock = file.NewBlockId(logFile, numBlock-1)
	fm.Read(logManager.currentBlock, &logManager.logPage)
	return logManager, nil
}

// Ensures that the log record corresponding to the
// specified LSN has been written to disk.
// All earlier log records will also be written to disk.
func (lm *LogManager) Flush(lsn int64) error {
	if lsn >= lm.lastSavedLSN {
		return lm.flushFile()
	}
	return nil
}

func (lm *LogManager) Iterator() (LogIterator, error) {
	err := lm.flushFile()
	if err != nil {
		return LogIterator{}, err
	}
	return NewLogIterator(lm.fileManager, lm.currentBlock)
}

// Appends a log record to the log buffer.
// The record consists of an arbitrary array of bytes.
// Log records are written right to left in the buffer.
// The size of the record is written before the bytes.
// The beginning of the buffer contains the location
// of the last-written record (the "boundary").
// Storing the records backwards makes it easy to read them in reverse order.
func (lm *LogManager) Append(data []byte) (int64, error) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	dataSpaceStart, err := lm.logPage.ReadInt(0)
	if err != nil {
		return 0, err
	}
	bytesNeeded := int64(8 + len(data))

	// does the log record fit?
	if dataSpaceStart-bytesNeeded < 8 {
		lm.flushFile()
		lm.currentBlock, err = lm.appendNewBlock()
		if err != nil {
			return 0, err
		}

		dataSpaceStart, err = lm.logPage.ReadInt(0)
		if err != nil {
			return 0, err
		}
	}

	offset := dataSpaceStart - bytesNeeded
	_, err = lm.logPage.WriteBytes(offset, data)
	if err != nil {
		return 0, err
	}
	err = lm.logPage.WriteInt(0, offset) // update the new data boundary
	if err != nil {
		return 0, err
	}

	lm.latestLSN += 1
	return lm.latestLSN, nil
}

// Appends a new page to the log file
func (lm *LogManager) appendNewBlock() (file.BlockId, error) {
	blockId, err := lm.fileManager.Append(lm.logFile)
	if err != nil {
		return file.BlockId{}, err
	}

	// Store the valid data start offset at the start of the page.
	// Note that record are insert in log page from the end of the page.
	err = lm.logPage.WriteInt(0, lm.fileManager.BlockSize())
	if err != nil {
		return file.BlockId{}, err
	}

	err = lm.fileManager.Write(blockId, &lm.logPage)
	return blockId, err
}

// flushFile flushes syncs the current page to the file
func (lm *LogManager) flushFile() error {
	return lm.fileManager.Write(lm.currentBlock, &lm.logPage)
}
