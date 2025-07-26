package message

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Message struct {
	ID          string `json:"id"`
	Payload     []byte `json:"payload"`
	Destination string `json:"destination"`
	Priority    int    `json:"priority"`
	Timestamp   int64  `json:"timestamp"`
	ProducerID  string `json:"producer"`
}

func GenerateID() (string, error) {
	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)
	return hex.EncodeToString(bytes), err
}

func CurrentTimestamp() int64 {
	return time.Now().UnixNano()
}

func NewMessage(payload []byte, destination, producerID string) (*Message, error) {

	messageId, err := GenerateID()

	if err != nil {
		return nil, err
	}

	return &Message{
		ID:          messageId,
		Payload:     payload,
		Destination: destination,
		ProducerID:  producerID,
		Timestamp:   CurrentTimestamp(),
		Priority:    5,
	}, nil
}
