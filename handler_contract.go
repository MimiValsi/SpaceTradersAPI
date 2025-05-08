package main

import "time"

type contract struct {
	Data struct {
		id            string
		FactionSymbol string
		Type          string
		terms         struct {
			Deadline time.Time
			Payment  struct {
				OnAccepted  int32
				OnFulfulled int32
			}
			Deliver []struct {
				TradeSymbol       string
				DestinationSymbol string
				UnitsRequired     int32
				UnitsFulfilled    int32
			}
		}
		Accepted         bool
		Fulfilled        bool
		DeadlineToAccept time.Time
	}
}
