package messages

import (
	"encoding/json"
	"github.com/thankala/gregor_chair_common/enums"
)

// BaseEvent struct to hold common fields
type BaseEvent struct {
	Event enums.Event     `json:"event"`
	Data  json.RawMessage `json:"data"`
}
