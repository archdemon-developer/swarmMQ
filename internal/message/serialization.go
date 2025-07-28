package message

import (
	"encoding/binary"
)

func writeStringWithLength(buffer []byte, offset int, value string, lengthBytes int) int {
	switch lengthBytes {
	case 2:
		binary.BigEndian.PutUint16(buffer[offset:], uint16(len(value)))
	case 4:
		binary.BigEndian.PutUint32(buffer[offset:], uint32(len(value)))
	}

	offset += lengthBytes
	copy(buffer[offset:], []byte(value))
	offset += len(value)
	return offset
}

func Serialize(msg *Message) ([]byte, error) {
	totalSize := 18 + len(msg.Destination) + len(msg.ProducerID) + len(msg.Payload) + len(msg.ID)
	buffer := make([]byte, totalSize)
	offset := 0

	offset = writeStringWithLength(buffer, offset, msg.Destination, 2)
	offset = writeStringWithLength(buffer, offset, msg.ProducerID, 2)

	binary.BigEndian.PutUint32(buffer[offset:], uint32(len(msg.Payload)))
	offset += 4
	copy(buffer[offset:], msg.Payload)
	offset += len(msg.Payload)

	binary.BigEndian.PutUint64(buffer[offset:], uint64(msg.Timestamp))
	offset += 8

	writeStringWithLength(buffer, offset, msg.ID, 2)

	return buffer, nil
}
