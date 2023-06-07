// Package weather
// a new line.
package weather

// CurrentCondition a string.
var CurrentCondition string

// CurrentLocation a string.
var CurrentLocation string

/* Forecast something
 a lot of strings. */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
