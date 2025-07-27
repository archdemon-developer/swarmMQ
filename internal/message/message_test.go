package message

import (
	"bytes"
	"testing"
	"time"
)

func TestValidNewMessage(t *testing.T) {
	// Arrange
	testPayload := []byte("test message")
	testDestination := "test-destination"
	testProducerID := "test-producer"

	// Act
	msg, err := NewMessage(testPayload, testDestination, testProducerID)

	// Assert
	if msg == nil {
		t.Errorf("Expected message to be created, got nil")
	}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if msg.Destination != testDestination {
		t.Errorf("Expected destination %s, got %s", testDestination, msg.Destination)
	}

	if !bytes.Equal(msg.Payload, testPayload) {
		t.Errorf("Expected payload %v, got %v", testPayload, msg.Payload)
	}

	if msg.ProducerID != testProducerID {
		t.Errorf("Expected producer ID %s, got %s", testProducerID, msg.ProducerID)
	}

	if msg.ID == "" {
		t.Errorf("Expected message ID to be generated, got empty string")
	}

	if msg.Priority != 5 {
		t.Errorf("Expected priority 5, got %d", msg.Priority)
	}
	if msg.Timestamp == 0 {
		t.Errorf("Expected timestamp to be set, got 0")
	}

	now := time.Now().UnixNano()
	timeDiff := now - msg.Timestamp
	if timeDiff < 0 || timeDiff > 5*int64(time.Second) {
		t.Errorf("Expected timestamp to be recent, got %d nanoseconds ago", timeDiff)
	}
}
