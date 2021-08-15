package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// CreatePollRequest struct
//swagger:model
type CreatePollRequest struct {
	// required: true
	CreatedBy string `json:"created_by" validate:"required"`
	// required: true
	PollDescription string `json:"poll_description" validate:"required"`
}

//FromJSONCreatePollRequest decodes request from json to struct
func (p *CreatePollRequest) FromJSONCreatePollRequest(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//ValidateCreatePollRequest validated the Request struct
func (p *CreatePollRequest) ValidateCreatePollRequest() error {
	validate := validator.New()
	return validate.Struct(p)
}

// CreatePollResponse struct
type CreatePollResponse struct {
	PollID          int64  `json:"poll_id"`
	PollDescription string `json:"poll_description"`
}

//ToJSON encodes response to Json
func (p *CreatePollResponse) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// a location response
//swagger:response createPollsMetasResponse
type createPollMetasResponseWrapper struct {
	//All Stock Metas in the system
	//in:body
	Body CreatePollResponse
}
