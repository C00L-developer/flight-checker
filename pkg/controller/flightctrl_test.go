package controller_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/C00L-developer/flight-checker/pkg/controller"
	"github.com/C00L-developer/flight-checker/pkg/pb"
	"github.com/stretchr/testify/require"
)

func TestFlightCtrl(t *testing.T) {
	testcases := []struct {
		desc    string
		flights []*pb.Flight
		result  *pb.Flight
		isVaild bool
		err     error
	}{
		{
			"happy path",
			[]*pb.Flight{
				{
					Source: "SFO",
					Target: "EWR",
				},
			},
			&pb.Flight{
				Source: "SFO",
				Target: "EWR",
			},
			true,
			nil,
		},
		{
			"happy path",
			[]*pb.Flight{
				{
					Source: "ATL",
					Target: "EWR",
				},
				{
					Source: "SFO",
					Target: "ATL",
				},
			},
			&pb.Flight{
				Source: "SFO",
				Target: "EWR",
			},
			true,
			nil,
		},
		{
			"happy path",
			[]*pb.Flight{
				{
					Source: "IND",
					Target: "EWR",
				},
				{
					Source: "SFO",
					Target: "ATL",
				},
				{
					Source: "GSO",
					Target: "IND",
				},
				{
					Source: "ATL",
					Target: "GSO",
				},
			},
			&pb.Flight{
				Source: "SFO",
				Target: "EWR",
			},
			true,
			nil,
		},
		{
			"zero path",
			[]*pb.Flight{},
			nil,
			false,
			controller.ErrZeroPath,
		},
		{
			"empty airport name",
			[]*pb.Flight{
				{
					Target: "SFO",
					Source: "",
				},
			},
			nil,
			false,
			controller.ErrInvalidAirport,
		},
		{
			"same flight",
			[]*pb.Flight{
				{
					Target: "SFO",
					Source: "SFO",
				},
			},
			nil,
			false,
			controller.ErrSameFlight,
		},
		{
			"cycle path",
			// 1->2->1
			[]*pb.Flight{
				{
					Target: "1",
					Source: "2",
				},
				{
					Target: "2",
					Source: "1",
				},
			},
			nil,
			false,
			controller.ErrInvalidPath,
		},
		{
			"mutli targets",
			// 1->2, 1->3
			[]*pb.Flight{
				{
					Target: "1",
					Source: "2",
				},
				{
					Target: "1",
					Source: "3",
				},
			},
			nil,
			false,
			controller.ErrInvalidPath,
		},
		{
			"mutli sources",
			// 1->3, 2->3
			[]*pb.Flight{
				{
					Target: "1",
					Source: "3",
				},
				{
					Target: "2",
					Source: "3",
				},
			},
			nil,
			false,
			controller.ErrInvalidPath,
		},
		{
			"cycle and line path",
			// 1->2, 3->4->5->3
			[]*pb.Flight{
				{
					Target: "1",
					Source: "2",
				},
				{
					Target: "5",
					Source: "3",
				},
				{
					Target: "4",
					Source: "5",
				},
				{
					Target: "3",
					Source: "4",
				},
			},
			nil,
			false,
			controller.ErrInvalidPath,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fc := &controller.FlightCtrl{}
			res, err := fc.GetSortedFlight(context.TODO(), &pb.GetSortedFlightRequest{Flights: tc.flights})
			if tc.isVaild {
				require.NoError(t, err)
				require.Equal(t, tc.result, res.Result)
			} else {
				require.EqualError(t, tc.err, err.Error())
			}
		})
	}
}

func TestRandomFlight(t *testing.T) {
	for i := 0; i < 10000; i++ {
		n := rand.Intn(100) + 1
		flights := getRandomPath(n)
		expected := &pb.Flight{
			Source: strconv.Itoa(0),
			Target: strconv.Itoa(n),
		}
		fc := &controller.FlightCtrl{}
		res, err := fc.GetSortedFlight(context.TODO(), &pb.GetSortedFlightRequest{Flights: flights})
		require.NoError(t, err)
		require.Equal(t, expected, res.Result)
	}
}

func getRandomPath(size int) []*pb.Flight {
	a := make([]int, size)

	for i := 0; i < size; i++ {
		a[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	res := make([]*pb.Flight, size)
	for i := 0; i < size; i++ {
		res[i] = &pb.Flight{
			Source: strconv.Itoa(a[i]),
			Target: strconv.Itoa(a[i] + 1),
		}
	}

	return res
}
