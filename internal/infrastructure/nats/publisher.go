package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

type NATSPublisher struct {
	js nats.JetStreamContext
}

func NewPublisher(js nats.JetStreamContext) *NATSPublisher {
	return &NATSPublisher{
		js: js,
	}
}

func (p *NATSPublisher) Publish(ctx context.Context, email []byte) error {
	subject := "email.send"

	ack, err := p.js.PublishAsync(subject, email)
	if err != nil {
		return err
	}

	select {
	case <-ack.Ok():
		fmt.Println("Email published succefully")
	case err := <-ack.Err():
		return fmt.Errorf("publish email error: %v", err)
	}
	return nil
}
