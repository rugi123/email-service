package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rugi123/email-service/internal/domain/models"
)

type Worker struct {
	ID     string
	js     nats.JetStreamContext
	stream string
}

func NewWorker(js nats.JetStreamContext, stream string, workerID string) (*Worker, error) {
	consumerName := fmt.Sprintf("worker-%s", workerID)

	_, err := js.AddConsumer(stream, &nats.ConsumerConfig{
		Durable:       consumerName,
		DeliverPolicy: nats.DeliverAllPolicy,
		AckPolicy:     nats.AckExplicitPolicy,
		AckWait:       30 * time.Second,
		MaxDeliver:    3,
	})
	if err != nil {
		return nil, err
	}

	return &Worker{
		ID:     workerID,
		js:     js,
		stream: stream,
	}, nil
}

func (w *Worker) Start() (chan models.Email, error) {
	ch := make(chan models.Email, 10)

	sub, err := w.js.PullSubscribe("email.send", fmt.Sprintf("worker-%s", w.ID))
	if err != nil {
		return nil, err
	}
	go w.processMessages(sub, ch)
	return ch, nil
}

func (w *Worker) processMessages(sub *nats.Subscription, ch chan models.Email) {
	for {
		msgs, err := sub.Fetch(10, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				continue
			}
			log.Printf("Fetch error: %v", err)
			continue
		}

		for _, msg := range msgs {
			var email models.Email
			err = json.Unmarshal(msg.Data, &email)
			fmt.Println(err)
			ch <- email
		}
	}
}
