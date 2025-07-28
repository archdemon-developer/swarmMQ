package message

type Message struct {
	ID          string `json:"id"`
	Payload     []byte `json:"payload"`
	Destination string `json:"destination"`
	Priority    int    `json:"priority"`
	Timestamp   int64  `json:"timestamp"`
	ProducerID  string `json:"producer"`
}

func NewMessage(payload []byte, destination, producerID string) (*Message, error) {

	messageId, err := GenerateID()

	if err != nil {
		return nil, err
	}

	msg := &Message{
		ID:          messageId,
		Payload:     payload,
		Destination: destination,
		ProducerID:  producerID,
		Timestamp:   CurrentTimestamp(),
		Priority:    5,
	}

	if err = Validate(msg); err != nil {
		return nil, err
	}

	return msg, nil
}
