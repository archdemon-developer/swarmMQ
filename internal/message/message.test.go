package message

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNewMessage_Valid(t *testing.T) {
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

func TestNewMessage_InvalidDestination(t *testing.T) {
	//Arrange
	testPayload := []byte("test message")
	testDestination := ""
	testProducerID := "test-producer"

	//Act
	msg, err := NewMessage(testPayload, testDestination, testProducerID)

	//Assert
	if err == nil {
		t.Errorf("Expected error to exist due to empty destination")
	} else if !strings.Contains(err.Error(), "destination") {
		t.Errorf("Expected error to be about destination")
	}

	if msg != nil {
		t.Errorf("Expected message to be empty due to empty destination")
	}
}

func TestNewMessage_InvalidPayload(t *testing.T) {
	//Arrange
	testPayload := []byte{}
	testDestination := "test-destination"
	testProducerID := "test-producer"

	//Act
	msg, err := NewMessage(testPayload, testDestination, testProducerID)

	//Assert
	if err == nil {
		t.Errorf("Expected error due to empty payload")
	} else if !strings.Contains(err.Error(), "payload") {
		t.Errorf("Expected error to be about payload")
	}

	if msg != nil {
		t.Errorf("Expected message to be empty due to empty payload")
	}
}

func TestNewMessage_InvalidProducerID(t *testing.T) {
	//Arrange
	testPayload := []byte("hello")
	testDestination := "test-destination"
	testProducerID := ""

	//Act
	msg, err := NewMessage(testPayload, testDestination, testProducerID)

	//Assert
	if err == nil {
		t.Errorf("Expected error due to empty producerID")
	} else if !strings.Contains(err.Error(), "producer id") {
		t.Errorf("Expected error to be about producerID")
	}

	if msg != nil {
		t.Errorf("Expected message to be empty due to invalid producerID")
	}
}

func TestSerialize_ValidMesageInput(t *testing.T) {

	msg := &Message{
		Destination: "test",
		ProducerID:  "prod1",
		Payload:     []byte("hi"),
		ID:          "abc",
	}

	data, _ := Serialize(msg)
	fmt.Println("Buffer size: ", len(data))
}
