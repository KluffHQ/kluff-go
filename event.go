package kluff

import (
	"encoding/json"
	"errors"
)

type EventType string

const (
	TRIGGER EventType = "TRIGGER"
	ACTION  EventType = "ACTION"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type Event struct {
	Data  json.RawMessage `json:"data"`
	Type  string          `json:"type"`
	Token string          `json:"token"` // used to authenticate the DB interactor
}

// used parse an event from bytes
func ParseEvent(data []byte) (*Event, error) {
	var evt Event
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	return &evt, nil
}

// the the kluff Interactor
func (e *Event) GetInteractor() (*Interactor, error) {
	return Get(e.Token)
}

// this is only valid when the request is coming from the a trigger
// returns an error is the request is coming from other thing else
func (e *Event) ParseTrigger() (*Trigger, error) {
	if e.Type != string(TRIGGER) {
		return nil, ErrInvalidRequest
	}
	return ParseTrigger(e.Data)
}

// parse the data from the action info `map[string]any`
// this return an error if the data does not conform to the `map[string]any`
func (e *Event) ParseMapData() (map[string]any, error) {
	var data map[string]any
	if err := json.Unmarshal(e.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
