package db

import "encoding/json"

func (d *Data) ParseMap() (map[string]any, error) {
	v := map[string]any{}
	err := json.Unmarshal(d.Result, &v)
	return v, err
}

func (d *Data) ParseMapSlice() ([]map[string]any, error) {
	v := []map[string]any{}
	err := json.Unmarshal(d.Result, &v)
	return v, err
}
