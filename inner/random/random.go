package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func HeadsOrTails() bool {
	return rand.Intn(2) != 0
}
