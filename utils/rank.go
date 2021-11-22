package utils

import (
	"math"
)

// Ref: https://wiki.mozilla.org/User:Jesse/NewFrecency

// half life of 1week(7days===604800sec)
const halflife int = 604800

// half-life decay
// => lambda = ln2/decay_time
func Lambda() float64 {
	return math.Log(2) / float64(halflife)
}

// age= current unix timestamp(in sec) - last accessed
// frecency_score= 1000*hits*e^(-lambda*age)
// each hits worth 1000 points
func FrecencyScore(hits, age int64) int64 {
	return int64(float64(1000) * float64(hits) * math.Exp(-Lambda()*float64(age)))
}
