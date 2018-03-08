// Package space contains a solution to the space age Exercism exercise
package space

// Planet is a type to represent a planets name
type Planet string

const earthYearInseconds = 31557600

// Age calculates
func Age(ageInSeconds float64, planet Planet) float64 {

	orbitalPeriods := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return ageInSeconds / (earthYearInseconds * orbitalPeriods[planet])
}
