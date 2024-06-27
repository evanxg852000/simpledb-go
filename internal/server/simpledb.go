package server

import (
	"log"

	"github.com/evanxg852000/simpledb/internal/buffer"
	"github.com/evanxg852000/simpledb/internal/file"
	walog "github.com/evanxg852000/simpledb/internal/log"
	"github.com/evanxg852000/simpledb/internal/metadata"
	"github.com/evanxg852000/simpledb/internal/plan"
	"github.com/evanxg852000/simpledb/internal/tx/recovery"
)

const (
	BLOCK_SIZE  = 400
	BUFFER_SIZE = 8
	LOG_FILE    = "simpledb.log"
)

type SimpleDB struct {
	fileManager     *file.FileManager
	bufferManager   *buffer.BufferManager
	logManager      *walog.LogManager
	metadataManager *metadata.MetadataManager
	planner         *plan.Planner
}

// A constructor useful for debugging.
func NewSimpleDB(directory string, blockSize int, bufferSize int) *SimpleDB {
	fileManager, err := file.NewFileManager(directory, int64(blockSize))
	if err != nil {
		log.Fatalf("could not open the database")
	}

	logManager, err := walog.NewLogManager(fileManager, LOG_FILE)
	if err != nil {
		log.Fatalf("could not open the database")
	}

	bufferManager := buffer.NewBufferManager(fileManager, logManager, bufferSize)

	tx := recovery.NewTransaction(fileManager, logManager, bufferManager)
	isNew := fileManager.IsNew()
	if isNew {
		log.Println("creating new database")
	} else {
		log.Println("recovering existing database")
		tx.Recover()
	}

	metadataManager := metadata.NewMetadataManager(isNew, tx)
	planner := plan.NewPlanner(
		plan.NewBasicQueryPlanner(metadataManager),
		plan.NewBasicUpdatePlanner(metadataManager),
	)
	tx.Commit()

	return &SimpleDB{
		fileManager:     fileManager,
		logManager:      logManager,
		bufferManager:   bufferManager,
		metadataManager: metadataManager,
		planner:         planner,
	}
}

func (sdb *SimpleDB) FileManager() *file.FileManager {
	return sdb.fileManager
}

func (sdb *SimpleDB) LogManager() *walog.LogManager {
	return sdb.logManager
}

func (sdb *SimpleDB) BufferManager() *buffer.BufferManager {
	return sdb.bufferManager
}

func (sdb *SimpleDB) NewTx() *recovery.Transaction {
	return recovery.NewTransaction(sdb.fileManager, sdb.logManager, sdb.bufferManager)
}

func (sdb *SimpleDB) MetadataManager() *metadata.MetadataManager {
	return sdb.metadataManager
}

func (sdb *SimpleDB) Planner() *plan.Planner {
	return sdb.planner
}
