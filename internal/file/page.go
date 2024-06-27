package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Page struct {
	data []byte
}

// For creating data buffers
func NewPage(blockSize int64) Page {
	return Page{data: make([]byte, blockSize)}
}

// For creating log pages
func NewPageWithData(data []byte) Page {
	return Page{data: data}
}

func (page *Page) ReadInt(offset int64) (int64, error) {
	var value int64
	buffer := bytes.NewBuffer(page.data[offset:])
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (page *Page) WriteInt(offset int64, value int64) error {
	buffer := bytes.NewBuffer(nil)
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	if int(offset+8) > len(page.data) {
		return fmt.Errorf("not enough space to encode data")
	}

	copy(page.data[offset:], buffer.Bytes())
	return nil
}

func (page *Page) ReadFloat(offset int64) (float64, error) {
	var value float64
	buffer := bytes.NewBuffer(page.data[offset:])
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (page *Page) WriteFloat(offset int64, value float64) error {
	buffer := bytes.NewBuffer(nil)
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	if int(offset+8) > len(page.data) {
		return fmt.Errorf("not enough space to encode data")
	}

	copy(page.data[offset:], buffer.Bytes())
	return nil
}

func (page *Page) ReadString(offset int64) (string, error) {
	data, err := page.ReadBytes(offset)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (page *Page) WriteString(offset int64, value string) (int64, error) {
	n, err := page.WriteBytes(offset, []byte(value))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (page *Page) ReadBytes(offset int64) ([]byte, error) {
	var length uint64
	buffer := bytes.NewBuffer(page.data[offset:])
	err := binary.Read(buffer, binary.LittleEndian, &length)
	if err != nil {
		return []byte{}, err
	}

	data := make([]byte, length)
	n, err := buffer.Read(data)
	if err != nil {
		return []byte{}, err
	}

	if n != int(length) {
		return []byte{}, fmt.Errorf("early EOF")
	}

	return data, nil
}

func (page *Page) WriteBytes(offset int64, value []byte) (int64, error) {
	length := uint64(len(value))

	buffer := bytes.NewBuffer(nil)
	err := binary.Write(buffer, binary.LittleEndian, length)
	if err != nil {
		return 0, err
	}

	_, err = buffer.Write(value)
	if err != nil {
		return 8, err
	}

	encodedData := buffer.Bytes()
	if int(offset)+len(encodedData) > len(page.data) {
		return 8, fmt.Errorf("not enough space to encode data")
	}

	copy(page.data[offset:], encodedData)
	return 8 + int64(len(encodedData)), nil
}

func (page *Page) Data() []byte {
	return page.data
}

func GetEncodingLength(len int64) int64 {
	return int64(8 + len)
}
