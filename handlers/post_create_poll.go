package handlers

import (
	"net/http"

	"github.com/hemansutanty/planningpoker/data"
	"github.com/hemansutanty/planningpoker/service"
)

// swagger:route POST /planningpoker/v1/create-poll createPoll createPoll
// Returns the location for a set of coordinates
// responses:
//  200: createPollMetasResponse

// CreatePoll handler function to pass the request onto service layer
func (p *PlanningPokerMeta) CreatePoll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle CreatePoll")
	prod := r.Context().Value(KeyplanningPokerData{}).(data.DNSRequest)
	response, apiErr := service.CreatePoll(&prod)
	if apiErr != nil {
		http.Error(rw, "Unable to create poll", http.StatusInternalServerError)
		return
	}
	jsonUnmarshalErr := response.ToJSON(rw)
	if jsonUnmarshalErr != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
