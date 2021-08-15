// Package handlers is to handle
//
// The purpose of this application is to estimate the complexity of a task based on a poll
//
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: https
//     Host: localhost
//     BasePath: /
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Hemansu Tanty<hemansu.tanty@gmail.com>
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
// swagger:meta
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/hemansutanty/planningpoker/data"
)

//PlanningPokerMeta struct
type PlanningPokerMeta struct {
	l *log.Logger
}

//NewPlanningPokerMeta passes a PlanningPokerMeta object
func NewPlanningPokerMeta(l *log.Logger) *PlanningPokerMeta {
	return &PlanningPokerMeta{l}
}

//KeyplanningPokerData struct
type KeyplanningPokerData struct{}

//MiddlewarePlanningPokerRequestValidation is responsible for Deserializing request and validating the request before apssing to the service layer
func (p PlanningPokerMeta) MiddlewarePlanningPokerRequestValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.PlanningPokerRequest{}
		err := prod.FromJSONDNSRequest(r.Body)
		if err != nil {
			p.l.Println("Error Deserializing Create Poll Request", err)
			http.Error(rw, "Unable to unmarshal Json", http.StatusBadRequest)
			return
		}

		err = prod.ValidateDNSRequest()
		if err != nil {
			p.l.Println("Error Validating product", err)
			http.Error(rw, "Unable to validate request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyplanningPokerData{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
