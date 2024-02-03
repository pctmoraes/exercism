// Package weather provides weather forecasts for the cities of Goblinocus.
package weather

// CurrentCondition represents the weather condition of a given city.
var CurrentCondition string

// CurrentLocation represents the location of a given city.
var CurrentLocation string

// Forecast returns the weather condition of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
