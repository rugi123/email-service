package nats

import "github.com/nats-io/nats.go"

func NewJetStream() (*nats.JetStreamContext, error) {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	streamConfig := &nats.StreamConfig{
		Name:     "EMAILS",
		Subjects: []string{"email.*"},
	}

	_, err = js.AddStream(streamConfig)
	if err != nil {
		return nil, err
	}

	return &js, nil
}
