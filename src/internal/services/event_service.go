package services

import (
	"errors"
	"scti/internal/models"
	"scti/internal/repos"
	"strings"
	"time"

	"github.com/google/uuid"
)

type EventService struct {
	EventRepo *repos.EventRepo
}

func NewEventService(repo *repos.EventRepo) *EventService {
	return &EventService{
		EventRepo: repo,
	}
}

func (s *EventService) CreateEvent(event *models.Event) error {
	event.ID = uuid.New().String()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	event.Slug = strings.ToLower(event.Slug)

	if err := s.EventRepo.CreateEvent(event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetEventBySlug(Slug string) (models.Event, error) {
	slug := strings.ToLower(Slug)
	event, err := s.EventRepo.GetEventBySlug(slug)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	events, err := s.EventRepo.GetAllEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) UpdateEvent(NewEvent *models.Event) (*models.Event, error) {
	StoredEvent, err := s.EventRepo.GetEventByID(NewEvent.ID)
	if err != nil {
		return nil, err
	}

	if NewEvent.Name != "" {
		StoredEvent.Name = NewEvent.Name
	}

	if NewEvent.Slug != "" {
		slug := strings.ToLower(NewEvent.Slug)
		StoredEvent.Slug = slug
	}

	if NewEvent.Description != "" {
		StoredEvent.Description = NewEvent.Description
	}

	if NewEvent.Location != "" {
		StoredEvent.Location = NewEvent.Location
	}

	if !NewEvent.StartDate.IsZero() {
		StoredEvent.StartDate = NewEvent.StartDate
	}

	if !NewEvent.EndDate.IsZero() {
		StoredEvent.EndDate = NewEvent.EndDate
	}

	if NewEvent.Redes != "" {
		StoredEvent.Redes = NewEvent.Redes
	}

	StoredEvent.UpdatedAt = time.Now()

	if err := s.EventRepo.UpdateEvent(&StoredEvent); err != nil {
		return nil, err
	}
	return &StoredEvent, nil
}

func (s *EventService) UpdateEventBySlug(Slug string, NewEvent *models.Event) (*models.Event, error) {
	slug := strings.ToLower(Slug)
	StoredEvent, err := s.EventRepo.GetEventBySlug(slug)
	if err != nil {
		return nil, err
	}

	if NewEvent.Name != "" {
		StoredEvent.Name = NewEvent.Name
	}

	if NewEvent.Description != "" {
		StoredEvent.Description = NewEvent.Description
	}

	if NewEvent.Slug != "" {
		slug := strings.ToLower(NewEvent.Slug)
		StoredEvent.Slug = slug
	}

	if NewEvent.Location != "" {
		StoredEvent.Location = NewEvent.Location
	}

	if !NewEvent.StartDate.IsZero() {
		StoredEvent.StartDate = NewEvent.StartDate
	}

	if !NewEvent.EndDate.IsZero() {
		StoredEvent.EndDate = NewEvent.EndDate
	}

	if NewEvent.Redes != "" {
		StoredEvent.Redes = NewEvent.Redes
	}

	StoredEvent.UpdatedAt = time.Now()

	if err := s.EventRepo.UpdateEvent(&StoredEvent); err != nil {
		return nil, err
	}
	return &StoredEvent, nil
}

func (s *EventService) DeleteEventBySlug(Slug string) error {
	slug := strings.ToLower(Slug)
	exists, err := s.EventRepo.ExistsEventBySlug(slug)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("event not found")
	}

	if err := s.EventRepo.DeleteEventBySlug(slug); err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetUserByID(userID string) (models.User, error) {
	user, err := s.EventRepo.GetUserByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *EventService) RegisterToEvent(userID string, Slug string) error {
	slug := strings.ToLower(Slug)
	if err := s.EventRepo.RegisterToEvent(userID, slug); err != nil {
		return err
	}
	return nil
}

func (s *EventService) UnregisterToEvent(userID string, Slug string) error {
	slug := strings.ToLower(Slug)
	if err := s.EventRepo.UnregisterToEvent(userID, slug); err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetEventAtendeesBySlug(slug string) (*[]models.EventUser, error) {
	slug = strings.ToLower(slug)
	event, err := s.EventRepo.GetEventAtendeesBySlug(slug)
	if err != nil {
		return nil, err
	}
	return event, nil
}
