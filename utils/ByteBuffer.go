package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// ByteBuffer is a Go equivalent of the C# ByteBuffer class
type ByteBuffer struct {
	bitPosition byte
	bitValue    byte
	writeStream *bytes.Buffer
	readStream  *bytes.Reader
}

// NewByteBuffer creates a new ByteBuffer instance
func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		bitPosition: 8,
		bitValue:    0,
		writeStream: new(bytes.Buffer),
		readStream:  nil,
	}
}

// NewByteBufferWithData creates a new ByteBuffer instance with initial data
func NewByteBufferWithData(data []byte) *ByteBuffer {
	return &ByteBuffer{
		bitPosition: 8,
		bitValue:    0,
		writeStream: nil,
		readStream:  bytes.NewReader(data),
	}
}

// ReadInt8 reads a signed byte from the buffer
func (b *ByteBuffer) ReadInt8() int8 {
	b.ResetBitPos()
	var result int8
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadInt16 reads a signed 16-bit integer from the buffer
func (b *ByteBuffer) ReadInt16() int16 {
	b.ResetBitPos()
	var result int16
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadInt32 reads a signed 32-bit integer from the buffer
func (b *ByteBuffer) ReadInt32() int32 {
	b.ResetBitPos()
	var result int32
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadInt64 reads a signed 64-bit integer from the buffer
func (b *ByteBuffer) ReadInt64() int64 {
	b.ResetBitPos()
	var result int64
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadUInt8 reads an unsigned byte from the buffer
func (b *ByteBuffer) ReadUInt8() uint8 {
	b.ResetBitPos()
	var result uint8
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadUInt16 reads an unsigned 16-bit integer from the buffer
func (b *ByteBuffer) ReadUInt16() uint16 {
	b.ResetBitPos()
	var result uint16
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadUInt32 reads an unsigned 32-bit integer from the buffer
func (b *ByteBuffer) ReadUInt32() uint32 {
	b.ResetBitPos()
	var result uint32
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadUInt64 reads an unsigned 64-bit integer from the buffer
func (b *ByteBuffer) ReadUInt64() uint64 {
	b.ResetBitPos()
	var result uint64
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadFloat reads a 32-bit floating-point number from the buffer
func (b *ByteBuffer) ReadFloat() float32 {
	b.ResetBitPos()
	var result float32
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadDouble reads a 64-bit floating-point number from the buffer
func (b *ByteBuffer) ReadDouble() float64 {
	b.ResetBitPos()
	var result float64
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadCString reads a null-terminated string from the buffer
func (b *ByteBuffer) ReadCString() string {
	b.ResetBitPos()
	var tmpString bytes.Buffer

	for {
		tmpChar, err := b.readStream.ReadByte()
		if err != nil || tmpChar == 0 {
			break
		}
		tmpString.WriteByte(tmpChar)
	}

	return tmpString.String()
}

// ReadString reads a string of specified length from the buffer
func (b *ByteBuffer) ReadString(length uint) string {
	if length == 0 {
		return ""
	}

	b.ResetBitPos()
	data := make([]byte, length)
	b.readStream.Read(data)
	return string(data)
}

// ReadBool reads a boolean value from the buffer
func (b *ByteBuffer) ReadBool() bool {
	b.ResetBitPos()
	var result bool
	binary.Read(b.readStream, binary.LittleEndian, &result)
	return result
}

// ReadBytes reads a specified number of bytes from the buffer
func (b *ByteBuffer) ReadBytes(count uint) []byte {
	b.ResetBitPos()
	data := make([]byte, count)
	b.readStream.Read(data)
	return data
}

// WriteInt8 writes a signed byte to the buffer
func (b *ByteBuffer) WriteInt8(data int8) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// HasUnfinishedBitPack checks if there are unfinished bit packs
func (b *ByteBuffer) HasUnfinishedBitPack() bool {
	return b.bitPosition != 8
}

// FlushBits flushes any remaining bits to the buffer
func (b *ByteBuffer) FlushBits() {
	if b.bitPosition == 8 {
		return
	}

	b.writeStream.WriteByte(b.bitValue)
	b.bitValue = 0
	b.bitPosition = 8
}

// ResetBitPos resets the bit position to 8
func (b *ByteBuffer) ResetBitPos() {
	if b.bitPosition > 7 {
		return
	}

	b.bitPosition = 8
	b.bitValue = 0
}

// ResetReadPos resets the read position to the beginning
func (b *ByteBuffer) ResetReadPos() {
	b.readStream.Seek(0, io.SeekStart)
	b.ResetBitPos()
}

// ReadToEnd reads the remaining bytes in the buffer
func (b *ByteBuffer) ReadToEnd() []byte {
	length := b.readStream.Len()
	data := make([]byte, length)
	b.readStream.Read(data)
	return data
}

// GetData gets the data from the buffer (either from readStream or writeStream)
func (b *ByteBuffer) GetData() []byte {
	// Check if both readStream and writeStream are nil
	if b.readStream == nil && b.writeStream == nil {
		fmt.Println("Both readStream and writeStream are nil")
		return nil
	}

	// If readStream is available, read data from it
	if b.readStream != nil {
		data, err := io.ReadAll(b.readStream)
		if err != nil {
			fmt.Println("Error reading from readStream:", err)
			return nil
		}
		return data
	}

	// If writeStream is available and readStream is nil, read data from writeStream
	if b.writeStream != nil {
		data, err := io.ReadAll(b.writeStream)
		if err != nil {
			fmt.Println("Error reading from writeStream:", err)
			return nil
		}
		return data
	}

	return nil
}

// GetSize gets the size of the buffer
func (b *ByteBuffer) GetSize() uint {
	return uint(b.writeStream.Len())
}

// GetCurrentStream gets the current stream based on read or write mode
func (b *ByteBuffer) GetCurrentStream() io.Reader {
	if b.writeStream != nil {
		return b.writeStream
	}
	return b.readStream
}

// Clear resets the buffer to its initial state
func (b *ByteBuffer) Clear() {
	b.bitPosition = 8
	b.bitValue = 0
	b.writeStream = new(bytes.Buffer)
}

// WriteInt16 writes a signed 16-bit integer to the buffer
func (b *ByteBuffer) WriteInt16(data int16) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteInt32 writes a signed 32-bit integer to the buffer
func (b *ByteBuffer) WriteInt32(data int32) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteInt64 writes a signed 64-bit integer to the buffer
func (b *ByteBuffer) WriteInt64(data int64) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteUInt8 writes an unsigned byte to the buffer
func (b *ByteBuffer) WriteUInt8(data uint8) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteUInt16 writes an unsigned 16-bit integer to the buffer
func (b *ByteBuffer) WriteUInt16(data uint16) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteUInt32 writes an unsigned 32-bit integer to the buffer
func (b *ByteBuffer) WriteUInt32(data uint32) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteUInt64 writes an unsigned 64-bit integer to the buffer
func (b *ByteBuffer) WriteUInt64(data uint64) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteFloat writes a 32-bit floating-point number to the buffer
func (b *ByteBuffer) WriteFloat(data float32) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteDouble writes a 64-bit floating-point number to the buffer
func (b *ByteBuffer) WriteDouble(data float64) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteCString writes a null-terminated string to the buffer
func (b *ByteBuffer) WriteCString(str string) {
	b.WriteString(str)
	b.writeStream.WriteByte(0)
}

// WriteString writes a string to the buffer
func (b *ByteBuffer) WriteString(str string) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, []byte(str))
}

// WriteBool writes a boolean value to the buffer
func (b *ByteBuffer) WriteBool(data bool) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteBytes writes a byte array to the buffer
func (b *ByteBuffer) WriteBytes(data []byte) {
	b.FlushBits()
	binary.Write(b.writeStream, binary.LittleEndian, data)
}

// WriteBytesWithLength writes a byte array to the buffer with a fixed length
func (b *ByteBuffer) WriteBytesWithLength(data []byte, length int) {
	b.FlushBits()
	if len(data) > length {
		data = data[:length] // truncate if necessary
	} else if len(data) < length {
		// Pad with zero bytes if necessary
		padding := make([]byte, length-len(data))
		data = append(data, padding...)
	}
	binary.Write(b.writeStream, binary.LittleEndian, data)
}
