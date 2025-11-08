package services

import (
	"context"
	"encoding/json"

	"github.com/rugi123/email-service/internal/domain/models"
)

type EventService struct {
	publisher EventPublisher
}

type EventPublisher interface {
	Publish(ctx context.Context, email []byte) error
}

func NewEventService(publisher EventPublisher) *EventService {
	return &EventService{
		publisher: publisher,
	}
}

func (s EventService) CreateEvent(ctx context.Context, email *models.Email) error {

	data, err := json.Marshal(email)
	if err != nil {
		return err
	}
	err = s.publisher.Publish(ctx, data)
	return err
}
