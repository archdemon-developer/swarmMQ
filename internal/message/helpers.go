package message

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func GenerateID() (string, error) {
	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)
	return hex.EncodeToString(bytes), err
}

func CurrentTimestamp() int64 {
	return time.Now().UnixNano()
}
