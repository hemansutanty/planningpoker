package service

import (
	"github.com/hemansutanty/planningpoker/data"
	"github.com/hemansutanty/planningpoker/errors"
)

//Welcome is used to create a poll
func Welcome() (*data.CreatePollResponse, *errors.APIError) {

	dnsResponse := &data.CreatePollResponse{
		PollID:          10,
		PollDescription: "New Poll created",
	}
	return dnsResponse, nil
}
