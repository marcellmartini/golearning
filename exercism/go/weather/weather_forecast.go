// Package weather forecast the current weather
// condition of various cities in Goblinocus.
package weather

// CurrentCondition is a current condition of a cities in Goblinocus.
var CurrentCondition string

// CurrentLocation is a current location of a cities in Goblinocus.
var CurrentLocation string

// Forecast return the current condition of a Goblinocus' citie at a current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
