package kluff

import "encoding/json"

type ActionEvent struct {
	Data json.RawMessage `json:"data"`
	// add more fields below
}
