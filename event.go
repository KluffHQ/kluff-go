package kluff

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

type Event struct {
	Type EventType `json:"type"`
	Data []byte    `json:"data"`
}
