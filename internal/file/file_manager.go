package file

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type FileManager struct {
	directory   string
	blockSize   int64
	openedFiles map[string]RandomAccessFile
	mu          sync.Mutex
	isNew       bool
}

func NewFileManager(directory string, block_size int64) (*FileManager, error) {
	_, err := os.Stat(directory)
	isNew := false
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(directory, 0755)
			isNew = true
		}

		if err != nil {
			return nil, err
		}
	}

	dirFile, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	files, err := dirFile.Readdir(0)
	if err != nil {
		return nil, err
	}

	openedFiles := map[string]RandomAccessFile{}
	for _, entry := range files {
		if entry.IsDir() {
			continue
		}

		rafName := entry.Name()
		rafPath := filepath.Join(directory, rafName)
		raf, err := NewRandomAccessFile(rafPath)
		if err != nil {
			return nil, err
		}
		openedFiles[rafName] = raf
	}

	return &FileManager{
		directory:   directory,
		blockSize:   block_size,
		openedFiles: openedFiles,
		isNew:       isNew,
	}, nil
}

func (fm *FileManager) Read(block BlockId, page *Page) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	raf, err := fm.getFile(block.FileName)
	if err != nil {
		return err
	}

	raf.ReadAt(page.Data(), block.BlockNum*fm.blockSize)
	return nil
}

func (fm *FileManager) Write(block BlockId, page *Page) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	raf, err := fm.getFile(block.FileName)
	if err != nil {
		return err
	}

	raf.WriteAt(page.Data(), block.BlockNum*fm.blockSize)
	return nil
}

func (fm *FileManager) Append(fileName string) (BlockId, error) {
	block_num, err := fm.BlockCount(fileName)
	if err != nil {
		return BlockId{}, err
	}

	block := NewBlockId(fileName, block_num)
	data := make([]byte, fm.blockSize)

	fm.mu.Lock()
	defer fm.mu.Unlock()
	raFile, exists := fm.openedFiles[block.FileName]
	if !exists {
		return BlockId{}, fmt.Errorf("file not found: `%s` ", block.FileName)
	}

	raFile.WriteAt(data, block.BlockNum*fm.blockSize)
	return block, nil
}

func (fm *FileManager) BlockCount(fileName string) (int64, error) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	raf, err := fm.getFile(fileName)
	if err != nil {
		return 0, err
	}

	size, err := raf.Size()
	if err != nil {
		return 0, err
	}

	return (size / fm.blockSize), nil
}

func (fm *FileManager) BlockSize() int64 {
	return fm.blockSize
}

func (fm *FileManager) getFile(fileName string) (RandomAccessFile, error) {
	raf, exists := fm.openedFiles[fileName]
	if exists {
		return raf, nil
	}

	rafPath := filepath.Join(fm.directory, fileName)
	raf, err := NewRandomAccessFile(rafPath)
	if err != nil {
		return RandomAccessFile{}, err
	}

	fm.openedFiles[fileName] = raf
	return raf, nil
}

func (fm *FileManager) IsNew() bool {
	return fm.isNew
}
