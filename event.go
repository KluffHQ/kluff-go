package kluff

import (
	"encoding/json"
	"errors"
	"net/http"
)

type EventType string

const (
	TRIGGER EventType = "TRIGGER"
	ACTION  EventType = "ACTION"
)

type TriggerType string

const (
	BEFORE_CREATE TriggerType = "BEFORE_CREATE"
	AFTER_CREATE  TriggerType = "AFTER_CREATE"
	BEFORE_UPDATE TriggerType = "BEFORE_UPDATE"
	AFTER_UPDATE  TriggerType = "AFTER_UPDATE"
	BEFORE_DELETE TriggerType = "BEFORE_DELETE"
	AFTER_DELETE  TriggerType = "AFTER_DELETE"
)

type EventRequest struct {
	*http.Request
	Type EventType
	Data json.RawMessage
}

type TriggerData struct {
	Type    TriggerType    `json:"type"`
	OldData map[string]any `json:"old"`
	Data    map[string]any `json:"data"`
}

type Event struct {
	Request *EventRequest
	w       http.ResponseWriter
}

func ParseEvent(w http.ResponseWriter, r *http.Request) (*Event, error) {
	req := struct {
		Type EventType       `json:"type"`
		Data json.RawMessage `json:"data"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return &Event{
		Request: &EventRequest{
			Request: r,
			Type:    req.Type,
			Data:    req.Data,
		},
		w: w,
	}, nil
}

func (e *Event) Log(value any) {
	//TODO: impl loging
}

func (e *Event) GetInteractor() (*Interactor, error) {
	token := e.Request.Header.Get("authorization")
	if token == "" {
		return nil, errors.New("invalid token")
	}
	return Get(token)
}

func (e *Event) ParseTriggerData() (*TriggerData, error) {
	if e.Request.Type != TRIGGER {
		return nil, errors.New("invalid request type")
	}
	var data TriggerData
	if err := json.Unmarshal(e.Request.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
