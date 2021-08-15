package service

import (
	"github.com/hemansutanty/planningpoker/data"
	"github.com/hemansutanty/planningpoker/errors"
)

//CreatePoll is used to create a poll
func CreatePoll(p *data.CreatePollRequest) (*data.CreatePollResponse, *errors.APIError) {

	dnsResponse := &data.CreatePollResponse{
		PollID:          10,
		PollDescription: "New Poll created",
	}
	return dnsResponse, nil
}
