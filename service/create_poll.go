package service

import (
	"net/http"
	"strconv"

	"github.com/hemansutanty/planningpoker/data"
	"github.com/hemansutanty/planningpoker/errors"
)

//CreatePoll is used to create a poll
func CreatePoll(p *data.CreatePollRequest) (*data.CreatePollResponse, *errors.APIError) {

	apiError := &errors.APIError{}
	// Parsing the string values to float values
	x, parseXError := strconv.ParseFloat(p.XCoordinate, 64)
	if parseXError != nil {
		apiError.Status = http.StatusBadRequest
		apiError.Message = "Unable to parse X coordinate"
		return nil, apiError
	}
	y, parseYError := strconv.ParseFloat(p.YCoordinate, 64)
	if parseYError != nil {
		apiError.Status = http.StatusBadRequest
		apiError.Message = "Unable to parse Y coordinate"
		return nil, apiError
	}
	z, parseZError := strconv.ParseFloat(p.ZCoordinate, 64)
	if parseZError != nil {
		apiError.Status = http.StatusBadRequest
		apiError.Message = "Unable to parse Z coordinate"
		return nil, apiError
	}
	vel, parseVelError := strconv.ParseFloat(p.Velocity, 64)
	if parseVelError != nil {
		apiError.Status = http.StatusBadRequest
		apiError.Message = "Unable to parse velocity"
		return nil, apiError
	}

	// Finding location using mathematical equation
	location := x*data.SectorID + y*data.SectorID + z*data.SectorID + vel

	dnsResponse := &data.DNSResponse{
		Location: location,
	}
	return dnsResponse, nil
}
