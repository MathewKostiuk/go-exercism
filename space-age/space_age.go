package space

type Planet string

const mercury = 0.2408467
const venus = 0.61519726
const mars = 1.8808158
const jupiter = 11.862615
const saturn = 29.447498
const uranus = 84.016846
const neptune = 164.79132
const year = 31557600

func Age(seconds float64, p Planet) float64 {
	var a float64
	switch p {
	case "Mercury":
		a = (seconds / year) / mercury
	case "Venus":
		a = (seconds / year) / venus
	case "Earth":
		a = seconds / year
	case "Mars":
		a = (seconds / year) / mars
	case "Jupiter":
		a = (seconds / year) / jupiter
	case "Saturn":
		a = (seconds / year) / saturn
	case "Uranus":
		a = (seconds / year) / uranus
	case "Neptune":
		a = (seconds / year) / neptune
	}

	return a
}
