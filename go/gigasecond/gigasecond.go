// Package gigasecond contains a solution to the gigasecond exercism exercise
package gigasecond

import "time"

// Const gigasecond provides an abstraction to the gigasecond function, rendering it more readable
const gigasecond = 1000000000

// AddGigasecond simply adds a giga second (1,000,000,000) to the input time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * gigasecond)
	return t
}
