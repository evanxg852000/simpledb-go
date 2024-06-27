package file

import (
	"bytes"
	"encoding/binary"
)

type ByteBuffer struct {
	buffer *bytes.Buffer
}

func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		buffer: bytes.NewBuffer(nil),
	}
}

func (bb *ByteBuffer) WriteInt(value int64) error {
	return binary.Write(bb.buffer, binary.LittleEndian, value)
}

func (bb *ByteBuffer) WriteFloat(value float64) error {
	return binary.Write(bb.buffer, binary.LittleEndian, value)
}

func (bb *ByteBuffer) WriteString(value string) error {
	return bb.WriteBytes([]byte(value))
}

func (bb *ByteBuffer) WriteBytes(value []byte) error {
	length := uint64(len(value))
	err := binary.Write(bb.buffer, binary.LittleEndian, length)
	if err != nil {
		return err
	}

	_, err = bb.buffer.Write(value)
	if err != nil {
		return err
	}

	return nil
}

func (bb *ByteBuffer) Data() []byte {
	return bb.buffer.Bytes()
}
