package dto

import (
	"context"
	"encoding/json"
	"time"
)

// MySetting ...
type MySetting struct {
	Product    string    `json:"product"`
	Module     string    `json:"module"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	AssignedAt time.Time `json:"assignedAt"`
}

// JsonString ...
func (a *MySetting) JsonString(ctx context.Context) string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}
