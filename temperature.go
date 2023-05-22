package temperature

func CelsiusToFahrenheit(t float64) float64 {
	return (t * 9 / 5) + 32
}

func FahrenheitToCelsius(t float64) float64 {
	return (t - 32) * 5 / 9
}
