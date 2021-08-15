package data

import (
	"encoding/json"
	"io"
)

// CreatePollResponse struct
type WelcomeResponse struct {
	Message string `json:"message"`
}

//ToJSON encodes response to Json
func (p *WelcomeResponse) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
