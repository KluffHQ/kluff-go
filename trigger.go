package kluff

import "encoding/json"

type TriggerType string

const (
	BEFORE_CREATE TriggerType = "BEFORE_CREATE"
	AFTER_CREATE  TriggerType = "AFTER_CREATE"
	BEFORE_UPDATE TriggerType = "BEFORE_UPDATE"
	AFTER_UPDATE  TriggerType = "AFTER_UPDATE"
	BEFORE_DELETE TriggerType = "BEFORE_DELETE"
	AFTER_DELETE  TriggerType = "AFTER_DELETE"
)

type Trigger struct {
	Data   map[string]any `json:"data"`
	Old    map[string]any `json:"old"` // used for update triggers
	Object string         `json:"object"`
	Type   TriggerType    `json:"type"`
}

func ParseTrigger(data []byte) (*Trigger, error) {
	var t Trigger
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
