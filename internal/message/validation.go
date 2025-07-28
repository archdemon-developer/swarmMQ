package message

import "errors"

func Validate(message *Message) error {
	if message.Destination == "" {
		return errors.New("destination cannot be empty")
	}

	if len(message.Payload) == 0 {
		return errors.New("payload cannot be empty")
	}

	if message.ProducerID == "" {
		return errors.New("producer id cannot be empty")
	}

	return nil
}
