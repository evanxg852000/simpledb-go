package file

import (
	"os"
	"sync"
)

type RandomAccessFile struct {
	file *os.File
	mu   *sync.Mutex
}

func NewRandomAccessFile(name string) (RandomAccessFile, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return RandomAccessFile{}, err
	}

	return RandomAccessFile{file: file, mu: new(sync.Mutex)}, nil
}

func (raf *RandomAccessFile) WriteAt(data []byte, offset int64) (int, error) {
	raf.mu.Lock()
	defer raf.mu.Unlock()

	n, err := raf.file.WriteAt(data, offset)
	if err != nil {
		return 0, err
	}

	err = raf.file.Sync()
	return n, err
}

func (raf *RandomAccessFile) ReadAt(data []byte, offset int64) (int, error) {
	raf.mu.Lock()
	defer raf.mu.Unlock()
	return raf.file.ReadAt(data, offset)
}

func (raf *RandomAccessFile) Size() (int64, error) {
	raf.mu.Lock()
	defer raf.mu.Unlock()
	info, err := raf.file.Stat()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func (raf *RandomAccessFile) Close() {
	raf.mu.Lock()
	defer raf.mu.Unlock()
	raf.file.Close()
}

// func (raf *RandomAccessFile) Seek(offset int64) (int64, error) {
// 	return raf.file.Seek(offset, io.SeekStart)
// }
