// Package space contains a solution to the space age Exercism exercise
package space

type Planet string

const (
	earth   Planet = "Earth"
	mercury Planet = "Mercury" //orbital period 0.2408467 Earth years
	venus   Planet = "Venus"   //orbital period 0.61519726 Earth years
	mars    Planet = "Mars"    //orbital period 1.8808158 Earth years
	jupiter Planet = "Jupitor" //orbital period 11.862615 Earth years
	saturn  Planet = "Saturn"  //orbital period 29.447498 Earth years
	uranus  Planet = "Uranus"  //orbital period 84.016846 Earth years
	neptune Planet = "Neptune" //orbital period 164.79132 Earth years
)

const earthYearInseconds = 31557600

// Age calculates
func Age(ageInSeconds float64, planet Planet) float64 {
	var ageInYears float64

	switch planet {
	case "Earth":
		ageInYears = ageInSeconds / earthYearInseconds
	case "Mercury":
		ageInYears = ageInSeconds / (earthYearInseconds * 0.2408467)
	case "Venus":
		ageInYears = ageInSeconds / (earthYearInseconds * 0.61519726)
	case "Mars":
		ageInYears = ageInSeconds / (earthYearInseconds * 1.8808158)
	case "Jupiter":
		ageInYears = ageInSeconds / (earthYearInseconds * 11.862615)
	case "Saturn":
		ageInYears = ageInSeconds / (earthYearInseconds * 29.447498)
	case "Uranus":
		ageInYears = ageInSeconds / (earthYearInseconds * 84.016846)
	case "Neptune":
		ageInYears = ageInSeconds / (earthYearInseconds * 164.79132)
	}
	return ageInYears
}
