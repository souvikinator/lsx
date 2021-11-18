package utils

import (
	"math"
)

// half life of 1week(7days===604800)
// need it in seconds or days?
// we'll find out which one works
const halflife int = 7

// half-life decay
// decay_time=ln2 / lambda
// => lambda = ln2/decay_time
func Lambda() float64 {
	return math.Log(2) / float64(halflife)
}

// age= last_visited-time.Now() [-ve value]
// freq_score= hits*e^(-lambda*age)
// [-lambda to balance negative value]
func FrecencyScore(hits, age int64) float64 {
	return math.Round(float64(1000) * float64(hits) * math.Exp(-Lambda()*float64(age)))
}
