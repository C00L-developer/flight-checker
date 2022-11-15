package controller

import (
	"context"
	"errors"

	"github.com/C00L-developer/flight-checker/pkg/pb"
)

type FlightCtrl struct {
	pb.UnsafeFlightServiceServer
}

func validFlight(f *pb.Flight) error {
	if f.Source == "" || f.Target == "" {
		return ErrInvalidAirport
	}
	if f.Source == f.Target {
		return ErrSameFlight
	}
	return nil
}

// GetSortedFlight gets the sorted flight.
func (c FlightCtrl) GetSortedFlight(ctx context.Context, req *pb.GetSortedFlightRequest) (*pb.GetSortedFlightResponse, error) {
	if len(req.Flights) == 0 {
		return nil, ErrZeroPath
	}

	prev, next := make(map[string]string), make(map[string]string)
	var start, end string
	for _, flight := range req.Flights {
		if err := validFlight(flight); err != nil {
			return nil, err
		}
		// u -> v
		u, v := flight.Source, flight.Target
		// check if have multiple sources or targets
		if _, found := next[u]; found {
			return nil, ErrInvalidPath
		}
		if _, found := prev[v]; found {
			return nil, ErrInvalidPath
		}
		// check if have a cycle
		if next[v] == u {
			return nil, ErrInvalidPath
		}
		// update prev and next
		start, end = u, v
		if prevU, found := prev[u]; found {
			start = prevU
		}
		if nextV, found := next[v]; found {
			end = nextV
		}
		next[start] = end
		prev[end] = start
	}

	return &pb.GetSortedFlightResponse{
		Result: &pb.Flight{
			Source: start,
			Target: end,
		},
	}, nil
}

var (
	ErrSameFlight     = errors.New("target should not be same with source")
	ErrInvalidAirport = errors.New("airport name should not be empty")
	ErrInvalidPath    = errors.New("invalid path")
	ErrZeroPath       = errors.New("path not be empty")
)
