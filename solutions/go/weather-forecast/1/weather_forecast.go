// Package weather reports conditions
// at a location.
package weather

// CurrentCondition must be string.
var CurrentCondition string
// CurrentLocation must be string.
var CurrentLocation string

// Forecast() returns the conditions at a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
