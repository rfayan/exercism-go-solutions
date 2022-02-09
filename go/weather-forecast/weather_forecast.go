// Package weather provides the ability to print the Forecast
// for a given location and condition and retrieve CurrentLocation
// and CurrentCondition based on the last method call.
package weather

// CurrentCondition stores the current weather condition as a string and
// is set with the second argument passed to the last call of Forecast.
var CurrentCondition string

// CurrentLocation stores the current weather condition as a string and
// is set with the first argument passed to the last call of Forecast.
var CurrentLocation string

// Forecast takes two strings, the name of a city and a weather condition,
// sets the variables CurrentLocation and CurrentCondition with them
// and returns a formatted string with weather forecast for this location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
