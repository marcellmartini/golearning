package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.New(rand.NewSource(time.Now().UnixMilli()))
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	SeedWithTime()
	return (rand.Int() % 20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
//
// You can think rand.Float64() as a percent
// of a desired range of numbers and use the
// following formula as a return:
//
//	min + rand.Float64() * (max - min)
//
// In our example:
//
//	0 + rand.Float64() * (12 - 0) or
//	rand.Float64() * 12
func GenerateWandEnergy() float64 {
	SeedWithTime()
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	var animals = []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}

	SeedWithTime()

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
